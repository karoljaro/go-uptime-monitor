package http

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestDefaultHTTPClient_Check_Success(t *testing.T) {
	// Mock Server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
	}))
	defer server.Close()

	client := NewDefaultHTTPClient(5 * time.Second)
	resp, err := client.Check(context.Background(), server.URL)

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if resp.StatusCode != 200 {
		t.Errorf("expected StatusCode 200, got %d", resp.StatusCode)
	}
	if resp.ResponseTime == 0 {
		t.Error("ResponseTime should be measured")
	}
}

func TestDefaultHTTPClient_Check_ServerError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	client := NewDefaultHTTPClient(5 * time.Second)
	resp, err := client.Check(context.Background(), server.URL)

	if err != nil {
		t.Errorf("expected no error (500 is valid response), got %v", err)
	}
	if resp.StatusCode != 500 {
		t.Errorf("expected StatusCode 500, got %d", resp.StatusCode)
	}
}

func TestDefaultHTTPClient_Check_Timeout(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(2 * time.Second) // wait 2 sekunds
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client := NewDefaultHTTPClient(100 * time.Millisecond) // timeout 100ms
	_, err := client.Check(context.Background(), server.URL)

	if err == nil {
		t.Error("expected timeout error, got nil")
	}
}

func TestDefaultHTTPClient_Check_InvalidURL(t *testing.T) {
	client := NewDefaultHTTPClient(5 * time.Second)
	_, err := client.Check(context.Background(), "not-a-valid-url://[invalid]")

	if err == nil {
		t.Error("expected error for invalid URL, got nil")
	}
}

func TestDefaultHTTPClient_Check_ContextCancel(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(1 * time.Second)
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	ctx, cancel := context.WithCancel(context.Background())
	cancel() // cancel context

	client := NewDefaultHTTPClient(5 * time.Second)
	_, err := client.Check(ctx, server.URL)

	if err == nil {
		t.Error("expected error from cancelled context, got nil")
	}
}
