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

func TestMemoryTargetRepository_FindByID_NotFound(t *testing.T) {
	id := "1"
	repo := NewMemoryTargetRepository()

	found, err := repo.FindByID(id)

	if found != nil {
		t.Errorf("expected nil, got %v", found)
	}

	if err == nil {
		t.Error("expected error, got nil")
	}
}

func TestMemoryTargetRepository_GetAll(t *testing.T) {
	repo := NewMemoryTargetRepository()

	targets := []struct {
		id   string
		url  string
		name string
	}{
		{"1", "http://example.com/1", "name-1"},
		{"2", "http://example.com/2", "name-2"},
		{"3", "http://example.com/3", "name-3"},
	}

	if arr, _ := repo.GetAll(); len(arr) > 0 {
		t.Error("expected empty array")
	}

	for _, tgt := range targets {
		target := domain.NewTarget(tgt.id, tgt.url, tgt.name, 30*time.Second)
		repo.Save(target)
	}

	foundTargets, err := repo.GetAll()

	if err != nil {
		t.Error("expected nil, got error: %w", err)
	}

	if len(foundTargets) != 3 {
		t.Errorf("expected 3 targets, got %d", len(foundTargets))
	}
}

func TestMemoryTargetRepository_Delete(t *testing.T) {
	id := "1"
	url := "https://example.com"
	name := "My API"
	interval := 30 * time.Second

	repo := NewMemoryTargetRepository()
	target := domain.NewTarget(id, url, name, interval)

	repo.Save(target)

	if found, _ := repo.GetAll(); len(found) == 0 {
		t.Errorf("Target not added, fix it")
	}

	repo.Delete(target.ID)

	_, err := repo.FindByID(target.ID)

	if err == nil {
		t.Error("expected error after delete, but found target")
	}
}

func TestMemoryTargetRepository_Update(t *testing.T) {
	id := "1"
	url := "https://example.com"
	name := "My API"
	interval := 30 * time.Second

	url2 := "https://example.com/posts"
	name2 := "Skibidi"
	interval2 := 10 * time.Second

	repo := NewMemoryTargetRepository()
	target := domain.NewTarget(id, url, name, interval)

	repo.Save(target)

	if found, _ := repo.GetAll(); len(found) == 0 {
		t.Errorf("Target not added, fix it")
	}

	target2 := domain.NewTarget(id, url2, name2, interval2)

	repo.Update(target2)

	found, err := repo.FindByID(id)

	if err != nil {
		t.Error("Unexpected error")
	}

	if found.URL != url2 {
		t.Errorf("expected %s, got %s", url2, found.URL)
	}

	if found.Name != name2 {
		t.Errorf("expected %s, got %s", name2, found.Name)
	}

	if found.Interval != interval2 {
		t.Errorf("expected %d, got %d", interval2, found.Interval)
	}
}

// ======================[RESULT]======================

func TestMemoryResultRepository_Save(t *testing.T) {
	id := "result-1"
	targetID := "targetId-1"
	statusCode := 500
	status := "Internal Server Error"
	responseTime := 15 * time.Millisecond

	repo := NewMemoryResultRepository()
	result := domain.NewResult(id, targetID, status, statusCode, responseTime)

	err := repo.Save(result)

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	found, err := repo.GetLastByTargetID(result.TargetID)

	if err != nil {
		t.Errorf("expected to find target, got error: %v", err)
	}

	if found.ID != id {
		t.Errorf("expected ID %s, got %s", id, found.ID)
	}

	if found.TargetID != targetID {
		t.Errorf("expected targetID %s, got %s", targetID, found.TargetID)
	}

	if found.Status != status {
		t.Errorf("expected status %s, got %s", status, found.Status)
	}

	if found.StatusCode != statusCode {
		t.Errorf("expected ID %d, got %d", statusCode, found.StatusCode)
	}

	if found.ResponseTime != responseTime {
		t.Errorf("expected time %s, got %s", responseTime, found.ResponseTime)
	}
}
