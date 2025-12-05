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
		t.Errorf("expect, erred 3 targets, got %d", len(foundTargets))
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

func TestMemoryResultRepository_FindByTargetID(t *testing.T) {
	repo := NewMemoryResultRepository()

	results := []struct {
		id           string
		targetID     string
		status       string
		statusCode   int
		responseTime time.Duration
	}{
		{"1", "t-1", "OK", 200, 120 * time.Millisecond},
		{"2", "t-2", "TIMEOUT", 504, 3 * time.Second},
		{"3", "t-3", "NOT_FOUND", 404, 45 * time.Millisecond},
		{"4", "t-3", "ERROR", 500, 830 * time.Millisecond},
		{"5", "t-3", "OK", 200, 10 * time.Millisecond},
	}

	if arr, _ := repo.FindByTargetID("t-3"); len(arr) > 0 {
		t.Error("expected empty array")
	}

	for _, rus := range results {
		result := domain.NewResult(rus.id, rus.targetID, rus.status, rus.statusCode, rus.responseTime)
		repo.Save(result)
	}

	foundResults, err := repo.FindByTargetID("t-3")

	if err != nil {
		t.Error("expected nil, got error: %w", err)
	}

	if len(foundResults) != 3 {
		t.Errorf("expected 3 targets, got %d", len(foundResults))
	}

	if _, err := repo.FindByTargetID("nonExistent"); err == nil {
		t.Error("expected error, got nil")
	}
}

func TestMemoryResultRepository_GetLastByTargetID(t *testing.T) {
	repo := NewMemoryResultRepository()

	results := []struct {
		id           string
		targetID     string
		status       string
		statusCode   int
		responseTime time.Duration
	}{
		{"1", "t-1", "OK", 200, 120 * time.Millisecond},
		{"2", "t-2", "TIMEOUT", 504, 3 * time.Second},
		{"3", "t-3", "NOT_FOUND", 404, 45 * time.Millisecond},
	}

	if arr, _ := repo.FindByTargetID("t-3"); len(arr) > 0 {
		t.Error("expected empty array")
	}

	for _, rus := range results {
		result := domain.NewResult(rus.id, rus.targetID, rus.status, rus.statusCode, rus.responseTime)
		repo.Save(result)
	}

	found, err := repo.GetLastByTargetID("t-3")

	if err != nil {
		t.Errorf("expected result, got %v", err)
	}

	if found.ID != "3" {
		t.Errorf("expected ID %s, got %s", "3", found.ID)
	}

	if found.TargetID != "t-3" {
		t.Errorf("expected targetID %s, got %s", "t-3", found.TargetID)
	}

}

// ======================[Alert]======================

func TestMemoryAlertRepository_Save(t *testing.T) {
	id := "alert-1"
	targetID := "target-1"
	alertType := "Error"
	message := "Internal Server Error"

	repo := NewMemoryAlertRepository()
	alert := domain.NewAlert(id, targetID, alertType, message)

	err := repo.Save(alert)

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	foundAlerts, err := repo.FindByTargetID(targetID)

	if err != nil {
		t.Error("expected nil, got error: %w", err)
	}

	if len(foundAlerts) == 0 {
		t.Fatal("expected at least one alert, got 0")
	}

	found := foundAlerts[0]

	if found.ID != alert.ID {
		t.Errorf("expected ID %s, got %s", alert.ID, found.ID)
	}
	if found.TargetID != alert.TargetID {
		t.Errorf("expected TargetID %s, got %s", alert.TargetID, found.TargetID)
	}
	if found.Type != alert.Type {
		t.Errorf("expected Type %s, got %s", alert.Type, found.Type)
	}
	if found.Message != alert.Message {
		t.Errorf("expected Message %s, got %s", alert.Message, found.Message)
	}
}

func TestMemoryAlertRepository_FindByTargetID(t *testing.T) {
	repo := NewMemoryAlertRepository()

	alerts := []struct {
		id        string
		targetID  string
		alertType string
		message   string
	}{
		{"alert-1", "target-1", "Error", "Internal Server Error"},
		{"alert-2", "target-1", "Warning", "High memory usage"},
		{"alert-3", "target-2", "Info", "Service started"},
		{"alert-4", "target-3", "Error", "Database connection failed"},
	}

	for _, ars := range alerts {
		alert := domain.NewAlert(ars.id, ars.targetID, ars.alertType, ars.message)
		repo.Save(alert)
	}

	foundAlerts, err := repo.FindByTargetID("target-3")

	if err != nil {
		t.Errorf("expected alert, got %v", err)
	}

	found := foundAlerts[0]

	if found.ID != "alert-4" {
		t.Errorf("expected ID 'alert-4', got %s", found.ID)
	}

	if found.TargetID != "target-3" {
		t.Errorf("expected TargetID 'target-3', got %s", found.TargetID)
	}

	if _, err := repo.FindByTargetID("nonExistent"); err == nil {
		t.Error("expected error, got nil")
	}
}

func TestMemoryAlertGetUnresolvedByTargetID(t *testing.T) {
	repo := NewMemoryAlertRepository()

	alerts := []struct {
		id        string
		targetID  string
		alertType string
		message   string
	}{
		{"alert-1", "target-1", "Error", "Internal Server Error"},
		{"alert-2", "target-1", "Warning", "High memory usage"},
		{"alert-3", "target-1", "Info", "Service started"},
		{"alert-4", "target-1", "Error", "Database connection failed"},
	}

	for _, ars := range alerts {
		alert := domain.NewAlert(ars.id, ars.targetID, ars.alertType, ars.message)
		repo.Save(alert)
	}

	found, err := repo.GetUnresolvedByTargetID("target-1")

	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}

	if len(found) != 4 {
		t.Errorf("expected array length 4, got %d", len(found))
	}

	nonExist, _ := repo.GetUnresolvedByTargetID("nonExist")

	if len(nonExist) > 0 {
		t.Errorf("expected empty array, got length %d", len(nonExist))
	}

	findByIdAlert, _ := repo.FindByTargetID("target-1")

	findByIdAlert[0].Resolve()
	repo.Update(findByIdAlert[0])
	findByIdAlert[1].Resolve()
	repo.Update(findByIdAlert[1])

	unresolved, err := repo.GetUnresolvedByTargetID("target-1")

	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}

	if len(unresolved) != 2 {
		t.Errorf("expected 2, got %d", len(unresolved))
	}

}

func TestMemoryAlertUpdate(t *testing.T) {
	id := "alert-1"
	targetID := "target-1"
	alertType := "Error"
	message := "Internal Server Error"

	alertType2 := "Warn"
	message2 := "Origin not set"

	repo := NewMemoryAlertRepository()
	alert := domain.NewAlert(id, targetID, alertType, message)

	repo.Save(alert)

	updatedAlert := domain.NewAlert(id, targetID, alertType2, message2)

	err := repo.Update(updatedAlert)

	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}

	foundTargets, err := repo.FindByTargetID(targetID)

	found := foundTargets[0]

	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}

	if found.Type != alertType2 {
		t.Errorf("expected %s, got %s", alertType2, found.Type)
	}

	if found.Message != message2 {
		t.Errorf("expected %s, got %s", message2, found.Message)
	}
}
