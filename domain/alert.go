package domain

import "time"

type Alert struct {
	ID         string
	TargetID   string
	Type       string
	Message    string
	CreatedAt  time.Time
	ResolvedAt *time.Time
	IsResolved bool
}

func NewAlert(id, targetID, alertType, message string) *Alert {
	return &Alert{
		ID:         id,
		TargetID:   targetID,
		Type:       alertType,
		Message:    message,
		CreatedAt:  time.Now(),
		IsResolved: false,
	}
}

func (a *Alert) Resolve() {
	now := time.Now()
	a.ResolvedAt = &now
	a.IsResolved = true
}