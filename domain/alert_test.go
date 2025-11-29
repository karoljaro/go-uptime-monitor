package domain

import (
	"testing"
	"time"
)

func TestNewAlert(t *testing.T) {
	id := "alert-1"
	targetID := "target-1"
	alertType := "Down"
	message := "Service is down"

	beforeCreatedAt := time.Now()
	alert := NewAlert(id, targetID, alertType, message)
	afterCreatedAt := time.Now()

	if alert.ID != id {
		t.Errorf("expected ID %s, got %s", id, alert.ID)
	}

	if alert.TargetID != targetID {
		t.Errorf("expected targetID %s, got %s", targetID, alert.TargetID)
	}

	if alert.Type != alertType {
		t.Errorf("expected type %s, got %s", alertType, alert.Type)
	}

	if alert.Message != message {
		t.Errorf("expected message %s, got %s", message, alert.Message)
	}

	if alert.CreatedAt.Before(beforeCreatedAt) || alert.CreatedAt.After(afterCreatedAt) {
		t.Errorf("CheckedAt should be between %v and %v, got %v", beforeCreatedAt, afterCreatedAt, alert.CreatedAt)
	}

	if alert.ResolvedAt != nil {
		t.Errorf("expected resolvedAt nil, got %s", alert.ResolvedAt)
	}

	if alert.IsResolved != false {
		t.Errorf("expected IsResolved false, got %t", alert.IsResolved)
	}
}

func TestResolve(t *testing.T) {
	alert := NewAlert("alert-1", "target-1", "Down", "Service is down")

	if alert.ResolvedAt != nil {
		t.Errorf("expected resolvedAt nil, got %s", alert.ResolvedAt)
	}

	if alert.IsResolved != false {
		t.Errorf("expected IsResolved false, got %t", alert.IsResolved)
	}

	beforeResolve := time.Now()
	alert.Resolve()
	afterResolve := time.Now()

	if alert.ResolvedAt == nil {
		t.Errorf("expected resolvedAt to not be nil")
	}

	if alert.ResolvedAt.Before(beforeResolve) || alert.ResolvedAt.After(afterResolve) {
		t.Errorf("ResolvedAt should be between %v and %v, got %v", beforeResolve, afterResolve, alert.ResolvedAt)
	}

	if alert.IsResolved == false {
		t.Errorf("expected IsResolved true, got %t", alert.IsResolved)
	}
}
