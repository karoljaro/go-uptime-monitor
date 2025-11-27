package domain

import "time"

type Result struct {
	ID           string
	TargetID     string
	Status       string
	StatusCode   int
	ResponseTime time.Duration
	CheckedAt    time.Time
	Error        error
}

func NewResult(id, targetID, status string, statusCode int, responseTime time.Duration) *Result {
	return &Result{
		ID:           id,
		TargetID:     targetID,
		Status:       status,
		StatusCode:   statusCode,
		ResponseTime: responseTime,
		CheckedAt:    time.Now(),
		Error:        nil,
	}
}
