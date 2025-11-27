package domain

import "time"

type Target struct {
	ID        string
	URL       string
	Name      string
	Interval  time.Duration
	IsActive  bool
	CreatedAt time.Time
}

func NewTarget(id, url, name string, interval time.Duration) *Target {
	return &Target{
		ID:        id,
		URL:       url,
		Name:      name,
		Interval:  interval,
		IsActive:  true,
		CreatedAt: time.Now(),
	}
}

func (t *Target) IsValid() bool {
	return len(t.URL) > 0 && t.Interval > 0
}