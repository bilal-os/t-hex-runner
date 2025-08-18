package utils

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"

	"github.com/spf13/viper"
)

// ProxyServer wraps http.Server so we can stop it later
type ProxyServer struct {
	server *http.Server
}

// StartProxy starts a reverse proxy server and returns a ProxyServer instance
func StartProxy(ready chan<- struct{}) (*ProxyServer, error) {
	endpoint := "http://localhost:4445" // fixed endpoint
	local := ":8888"                     // default listener
	key := viper.GetString("api_key")
	proj := viper.GetString("project_name")

	if key == "" {
		return nil, fmt.Errorf("api_key not set in config")
	}
	if proj == "" {
		proj = "UNNAMED PROJECT"
	}

	targetURL, err := url.Parse(endpoint)
	if err != nil {
		return nil, fmt.Errorf("error parsing endpoint URL: %w", err)
	}

	// request new TestId
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/thex/test", endpoint), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create test request: %w", err)
	}
	req.Header.Add("thex-key", key)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to request TestId: %w", err)
	}
	defer resp.Body.Close()

	testId := resp.Header.Get("thex-test")
	if testId == "" {
		return nil, fmt.Errorf("THex refused to send TestId")
	}

	// setup reverse proxy
	proxy := httputil.NewSingleHostReverseProxy(targetURL)
	proxy.Director = func(req *http.Request) {
		req.Header.Add("thex-key", key)
		req.Header.Add("thex-proj", proj)
		req.Header.Add("thex-test", testId)
		req.URL.Scheme = targetURL.Scheme
		req.URL.Host = targetURL.Host
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.String())
		proxy.ServeHTTP(w, r)
	})

	srv := &http.Server{
		Addr:    local,
		Handler: mux,
	}

	// signal readiness
	close(ready)

	// start serving in goroutine
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Proxy server failed: %v", err)
		}
	}()

	return &ProxyServer{server: srv}, nil
}

// Stop gracefully shuts down the proxy server
func (p *ProxyServer) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return p.server.Shutdown(ctx)
}
