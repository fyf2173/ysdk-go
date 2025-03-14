package xreq

import (
	"compress/gzip"
	"context"
	"errors"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"
)

// Mock implementation of IResponse interface
type mockResponse struct {
	Data string
	Err  error
}

func (m *mockResponse) Unmarshal(data []byte, _ interface{}) error {
	if m.Err != nil {
		return m.Err
	}
	m.Data = string(data)
	return nil
}

func TestClientRequest(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	slog.SetDefault(logger)
	t.Run("successful request", func(t *testing.T) {
		// Setup mock server
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"message": "success"}`))
		}))
		defer server.Close()

		client := NewClientDefault()
		resp := &mockResponse{}

		// Execute request
		err := client.Request(context.Background(), "GET", server.URL, nil, resp)

		// Validate results
		if err != nil {
			t.Fatalf("Expected no error, got: %v", err)
		}
		if resp.Data != `{"message": "success"}` {
			t.Errorf("Expected response data to be %q, got %q", `{"message": "success"}`, resp.Data)
		}
	})

	t.Run("error status code", func(t *testing.T) {
		// Setup mock server that returns error status
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
		}))
		defer server.Close()

		client := NewClientDefault()
		resp := &mockResponse{}

		// Execute request
		err := client.Request(context.Background(), "GET", server.URL, nil, resp)

		// Validate results
		if err == nil {
			t.Fatal("Expected error for status 404, got nil")
		}
		if !strings.Contains(err.Error(), "errorstatus:404") {
			t.Errorf("Expected error message to contain 'errorstatus:404', got: %v", err)
		}
	})

	t.Run("with request parameters", func(t *testing.T) {
		// Setup mock server that checks for Content-Type header
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Header.Get("Content-Type") != "application/json" {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"message": "params received"}`))
		}))
		defer server.Close()

		client := NewClientDefault()
		resp := &mockResponse{}
		params := struct {
			Name string `json:"name"`
		}{Name: "test"}

		// Execute request
		err := client.Request(context.Background(), "POST", server.URL, params, resp)

		// Validate results
		if err != nil {
			t.Fatalf("Expected no error, got: %v", err)
		}
		if resp.Data != `{"message": "params received"}` {
			t.Errorf("Expected response data to be %q, got %q", `{"message": "params received"}`, resp.Data)
		}
	})

	t.Run("with custom request options", func(t *testing.T) {
		// Setup mock server that checks for custom header
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Header.Get("X-Custom-Header") != "test-value" {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"message": "custom options applied"}`))
		}))
		defer server.Close()

		client := NewClientDefault()
		resp := &mockResponse{}

		// Custom request option
		customOption := func(req *http.Request) {
			req.Header.Set("X-Custom-Header", "test-value")
		}

		// Execute request with custom option
		err := client.Request(context.Background(), "GET", server.URL, nil, resp, customOption)

		// Validate results
		if err != nil {
			t.Fatalf("Expected no error, got: %v", err)
		}
		if resp.Data != `{"message": "custom options applied"}` {
			t.Errorf("Expected response data to be %q, got %q", `{"message": "custom options applied"}`, resp.Data)
		}
	})

	t.Run("unmarshal error", func(t *testing.T) {
		// Setup mock server
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"message": "success"}`))
		}))
		defer server.Close()

		client := NewClientDefault()
		// Create a response that will return an error during unmarshaling
		resp := &mockResponse{Err: errors.New("unmarshal error")}

		// Execute request
		err := client.Request(context.Background(), "GET", server.URL, nil, resp)

		// Validate results
		if err == nil {
			t.Fatal("Expected unmarshal error, got nil")
		}
		if err.Error() != "unmarshal error" {
			t.Errorf("Expected error message 'unmarshal error', got: %v", err)
		}
	})

	t.Run("http client error", func(t *testing.T) {
		client := NewClientDefault()
		resp := &mockResponse{}

		// Use invalid URL to trigger HTTP client error
		err := client.Request(context.Background(), "GET", "http://invalid-url-that-should-not-exist.example", nil, resp)

		// Validate results
		if err == nil {
			t.Fatal("Expected HTTP client error, got nil")
		}
	})
}

// Additional tests for client.Request
func TestClientRequestAdditional(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	slog.SetDefault(logger)

	t.Run("large response body", func(t *testing.T) {
		// Setup mock server that returns a large response body
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			// Generate a response larger than DefaultRespSize
			largeResponse := strings.Repeat(`{"key": "value"},`, DefaultRespSize/16)
			w.Write([]byte(`{"data": [` + largeResponse + `{}]}`))
		}))
		defer server.Close()

		client := NewClientDefault()
		resp := &mockResponse{}

		// Execute request
		err := client.Request(context.Background(), "GET", server.URL, nil, resp)

		// Validate results
		if err != nil {
			t.Fatalf("Expected no error with large response, got: %v", err)
		}
		// We don't check the exact content since it's large, but we verify it's not empty
		if len(resp.Data) == 0 {
			t.Error("Expected non-empty response data")
		}
	})

	t.Run("with query parameters", func(t *testing.T) {
		// Setup mock server that validates query parameters
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			query := r.URL.Query()
			if query.Get("param1") != "value1" || query.Get("param2") != "value2" {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"message": "query params received"}`))
		}))
		defer server.Close()

		client := NewClientDefault()
		resp := &mockResponse{}

		// Add query options as a custom request option
		addQueryParams := func(req *http.Request) {
			q := req.URL.Query()
			q.Add("param1", "value1")
			q.Add("param2", "value2")
			req.URL.RawQuery = q.Encode()
		}

		// Execute request with query parameters
		err := client.Request(context.Background(), "GET", server.URL, nil, resp, addQueryParams)

		// Validate results
		if err != nil {
			t.Fatalf("Expected no error with query parameters, got: %v", err)
		}
		if resp.Data != `{"message": "query params received"}` {
			t.Errorf("Expected response data to be %q, got %q", `{"message": "query params received"}`, resp.Data)
		}
	})

	t.Run("server error (500)", func(t *testing.T) {
		// Setup mock server that returns 500 error
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"error": "internal server error"}`))
		}))
		defer server.Close()

		client := NewClientDefault()
		resp := &mockResponse{}

		// Execute request
		err := client.Request(context.Background(), "GET", server.URL, nil, resp)

		// Validate results
		if err == nil {
			t.Fatal("Expected error for status 500, got nil")
		}
		if !strings.Contains(err.Error(), "errorstatus:500") {
			t.Errorf("Expected error message to contain 'errorstatus:500', got: %v", err)
		}
	})

	t.Run("response with content-encoding gzip", func(t *testing.T) {
		// Setup mock server that returns gzipped response
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Encoding", "gzip")
			w.WriteHeader(http.StatusOK)

			// Create gzipped content
			gzw := gzip.NewWriter(w)
			defer gzw.Close()

			gzw.Write([]byte(`{"message": "gzipped response"}`))
		}))
		defer server.Close()

		client := NewClientDefault()
		resp := &mockResponse{}

		// Execute request - http.Client should handle gzip automatically
		err := client.Request(context.Background(), "GET", server.URL, nil, resp)

		// Validate results
		if err != nil {
			t.Fatalf("Expected no error with gzipped response, got: %v", err)
		}
		if resp.Data != `{"message": "gzipped response"}` {
			t.Errorf("Expected response data to be %q, got %q", `{"message": "gzipped response"}`, resp.Data)
		}
	})

	t.Run("request with timeout", func(t *testing.T) {
		// Setup mock server that delays response
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Sleep longer than our timeout
			time.Sleep(200 * time.Millisecond)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"message": "too late"}`))
		}))
		defer server.Close()

		// Create client with short timeout
		client := &Client{
			Client: http.Client{
				Timeout: 100 * time.Millisecond,
			},
		}
		resp := &mockResponse{}

		// Execute request
		err := client.Request(context.Background(), "GET", server.URL, nil, resp)

		// Validate results
		if err == nil {
			t.Fatal("Expected timeout error, got nil")
		}
		// Check if error is related to timeout
		if !strings.Contains(err.Error(), "timeout") &&
			!strings.Contains(err.Error(), "deadline exceeded") {
			t.Errorf("Expected timeout error, got: %v", err)
		}
	})
}
