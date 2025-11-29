package domain

import (
	"testing"
	"time"
)

func TestNewTarget(t *testing.T) {
	id := "target-1"
	url := "https://example.com"
	name := "My API"
	interval := 30 * time.Second

	target := NewTarget(id, url, name, interval)

	if target.ID != id {
		t.Errorf("expected ID %s, got %s", id, target.ID)
	}

	if target.URL != url {
		t.Errorf("expected URL %s, got %s", url, target.URL)
	}

	if target.Name != name {
		t.Errorf("expected name %s, got %s", name, target.Name)
	}

	if target.Interval != interval {
		t.Errorf("expected interval %s, got %s", interval, target.Interval)
	}

	if target.IsActive == false {
		t.Errorf("expected IsActive true, got false")
	}
}

func TestIsValid(t *testing.T) {
	id := "target-1"
	name := "My API"

	url1 := "https://example.com"
	url2 := ""

	interval1 := 30 * time.Second
	interval2 := 0 * time.Second

	target1 := NewTarget(id, url1, name, interval1)
	target2 := NewTarget(id, url2, name, interval1)
	target3 := NewTarget(id, url1, name, interval2)
	target4 := NewTarget(id, url2, name, interval2)

	if target1.IsValid() == false {
		t.Errorf("expected interval true got false for target1")
	}

	if target2.IsValid() == true {
		t.Errorf("expected interval false, got true for target2")
	}

	if target3.IsValid() == true {
		t.Errorf("expected interval false, got true for target3")
	}

	if target4.IsValid() == true {
		t.Errorf("expected interval false, got true for target4")
	}
}
