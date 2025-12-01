package storage

import (
	"github.com/karoljaro/go-uptime-monitor/domain"
	"testing"
	"time"
)

func TestMemoryTargetRepository_Save(t *testing.T) {
	id := "1"
	url := "https://example.com"
	name := "My API"
	interval := 30 * time.Second

	repo := NewMemoryTargetRepository()
	target := domain.NewTarget(id, url, name, interval)

	err := repo.Save(target)

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	found, err := repo.FindByID(id)

	if err != nil {
		t.Errorf("expected to find target, got error: %v", err)
	}

	if found.ID != id {
		t.Errorf("expected ID %s, got %s", id, found.ID)
	}

	if found.URL != url {
		t.Errorf("expected URL %s, got %s", url, found.URL)
	}

	if found.Name != name {
		t.Errorf("expected name %s, got %s", name, found.Name)
	}

	if found.Interval != interval {
		t.Errorf("expected interval %s, got %s", interval, found.Interval)
	}

	if found.IsActive == false {
		t.Errorf("expected IsActive true, got false")
	}

}


