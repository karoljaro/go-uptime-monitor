package domain

import (
	"context"
	"time"
)

type HTTPResponse struct {
	StatusCode   int
	ResponseTime time.Duration
	Error        error
}

type HTTPClient interface {
	Check(ctx context.Context, url string) (*HTTPResponse, error)
}
