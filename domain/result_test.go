package domain

import (
	"testing"
	"time"
)

func TestNewResult(t *testing.T) {
	id := "alert-1"
	targetID := "target-1"
	status := "OK"
	statusCode := 200
	responseTime := 150 * time.Millisecond

	beforeCheckedAt := time.Now()
	result := NewResult(id, targetID, status, statusCode, responseTime)
	afterCheckedAt := time.Now()

	if result.ID != id {
		t.Errorf("expected ID %s, got %s", id, result.ID)
	}

	if result.TargetID != targetID {
		t.Errorf("expected tragetID %s, got %s", targetID, result.TargetID)
	}

	if result.Status != status {
		t.Errorf("expected status %s, got %s", status, result.Status)
	}

	if result.StatusCode != statusCode {
		t.Errorf("expected statusCode %d, got %d", statusCode, result.StatusCode)
	}

	if result.ResponseTime != responseTime {
		t.Errorf("expected responseTime %s, got %s", responseTime, result.ResponseTime)
	}

	if result.Error != nil {
		t.Errorf("expected responseTime nil, got %s", result.Error)
	}

	if result.CheckedAt.Before(beforeCheckedAt) || result.CheckedAt.After(afterCheckedAt) {
		t.Errorf("CheckedAt should be between %v and %v, got %v", beforeCheckedAt, afterCheckedAt, result.CheckedAt)
	}
}