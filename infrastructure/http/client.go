package http

import (
	"context"
	"net/http"
	"time"
	"github.com/karoljaro/go-uptime-monitor/domain"
)

type DefaultHTTPClient struct {
	client  *http.Client
	timeout time.Duration
}

func NewDefaultHTTPClient(timeout time.Duration) *DefaultHTTPClient {
	return &DefaultHTTPClient{
		client: &http.Client{
			Timeout: timeout,
		},
		timeout: timeout,
	}
}

func (c *DefaultHTTPClient) Check(ctx context.Context, url string) (*domain.HTTPResponse, error) {
	start := time.Now()

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	responseTime := time.Since(start)

	return &domain.HTTPResponse{
		StatusCode: res.StatusCode,
		ResponseTime: responseTime,
	}, nil
}