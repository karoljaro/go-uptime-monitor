package usecase

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/karoljaro/go-uptime-monitor/domain"
)

// ----- > MOCK DEF < -----

// ========================[Target Repository]========================

type MockTargetRepository struct {
	FindByIDFunc func(id string) (*domain.Target, error)
}

func (m *MockTargetRepository) FindByID(id string) (*domain.Target, error) {
	return m.FindByIDFunc(id)
}

func (m *MockTargetRepository) Delete(id string) error {
	return nil
}

func (m *MockTargetRepository) Update(target *domain.Target) error {
	return nil
}

func (m *MockTargetRepository) GetAll() ([]*domain.Target, error) {
	return []*domain.Target{}, nil
}

func (m *MockTargetRepository) Save(target *domain.Target) error {
	return nil
}

// ========================[Result Repository]========================

type MockResultRepository struct {
	SaveFunc              func(result *domain.Result) error
	SavedResults          []*domain.Result
	GetLastByTargetIDFunc func(targetID string) (*domain.Result, error)
}

func (m *MockResultRepository) Save(result *domain.Result) error {
	m.SavedResults = append(m.SavedResults, result)
	if m.SaveFunc != nil {
		return m.SaveFunc(result)
	}

	return nil
}

func (m *MockResultRepository) GetLastByTargetID(targetID string) (*domain.Result, error) {
	if m.GetLastByTargetIDFunc != nil {
		return m.GetLastByTargetIDFunc(targetID)
	}

	return nil, fmt.Errorf("not found")
}

func (m *MockResultRepository) FindByTargetID(targetID string) ([]*domain.Result, error) {
	return []*domain.Result{}, nil
}

// ========================[Alert Repository]========================

type MockAlertRepository struct {
	SaveFunc                    func(alert *domain.Alert) error
	SavedAlerts                 []*domain.Alert
	GetUnresolvedByTargetIDFunc func(targetID string) ([]*domain.Alert, error)
	UpdateFunc                  func(alert *domain.Alert) error
	UpdatedAlerts               []*domain.Alert
}

func (m *MockAlertRepository) Save(alert *domain.Alert) error {
	m.SavedAlerts = append(m.SavedAlerts, alert)
	if m.SaveFunc != nil {
		return m.SaveFunc(alert)
	}

	return nil
}

func (m *MockAlertRepository) GetUnresolvedByTargetID(targetID string) ([]*domain.Alert, error) {
	if m.GetUnresolvedByTargetIDFunc != nil {
		return m.GetUnresolvedByTargetIDFunc(targetID)
	}

	return make([]*domain.Alert, 0), nil
}

func (m *MockAlertRepository) Update(alert *domain.Alert) error {
	m.UpdatedAlerts = append(m.UpdatedAlerts, alert)
	if m.UpdateFunc != nil {
		return m.UpdateFunc(alert)
	}

	return nil
}

func (m *MockAlertRepository) FindByTargetID(targetID string) ([]*domain.Alert, error) {
	return []*domain.Alert{}, nil
}

// ========================[HTTP Client]========================

type MockHTTPClient struct {
	CheckFunc func(ctx context.Context, url string) (*domain.HTTPResponse, error)
}

func (m *MockHTTPClient) Check(ctx context.Context, url string) (*domain.HTTPResponse, error) {
	if m.CheckFunc != nil {
		return m.CheckFunc(ctx, url)
	}

	return nil, nil
}

// ========================[ID Generator]========================

type MockIDGenerator struct {
	GenerateFunc func() string
}

func (m *MockIDGenerator) Generate() string {
	if m.GenerateFunc != nil {
		return m.GenerateFunc()
	}

	return ""
}

// ----- > Helpers < -----

func newMockTargetRepository() *MockTargetRepository {
	return &MockTargetRepository{
		FindByIDFunc: func(id string) (*domain.Target, error) {
			return &domain.Target{
				ID:  "target-1",
				URL: "https://example.com",
			}, nil
		},
	}
}

func newMockResultRepository() *MockResultRepository {
	return &MockResultRepository{
		SaveFunc: func(result *domain.Result) error {
			return nil
		},
		GetLastByTargetIDFunc: func(targetID string) (*domain.Result, error) {
			return &domain.Result{
				ID:           "generated-1",
				TargetID:     "target-1",
				Status:       "OK",
				StatusCode:   200,
				ResponseTime: 30 * time.Microsecond,
			}, nil
		},
	}
}

func newMockAlertRepository() *MockAlertRepository {
	return &MockAlertRepository{
		SaveFunc: func(alert *domain.Alert) error {
			return nil
		},
		GetUnresolvedByTargetIDFunc: func(targetID string) ([]*domain.Alert, error) {
			alerts := []*domain.Alert{
				{ID: "gen-alert-1", TargetID: "target-1", Type: "Error", Message: "Internal Server Error"},
				{ID: "gen-alert-2", TargetID: "target-1", Type: "Error", Message: "Database connection failed"},
			}
			return alerts, nil
		},
	}
}

func newMockIDGenerator() *MockIDGenerator {
	return &MockIDGenerator{
		GenerateFunc: func() string {
			return "generatedID"
		},
	}
}

func newMockHTTPClient() *MockHTTPClient {
	return &MockHTTPClient{
		CheckFunc: func(ctx context.Context, url string) (*domain.HTTPResponse, error) {
			return &domain.HTTPResponse{
				StatusCode:   200,
				ResponseTime: 100 * time.Millisecond,
			}, nil
		},
	}
}

// ----- > Test cases < -----

func TestCheckTarget_StatusOK(t *testing.T) {
	mockTargetRepo := newMockTargetRepository()
	mockResultRepo := newMockResultRepository()
	mockAlertRepo := newMockAlertRepository()
	mockHTTPClient := newMockHTTPClient()
	mockIDGenerator := newMockIDGenerator()

	usecase := NewMonitorUseCase(mockTargetRepo, mockResultRepo, mockAlertRepo, mockHTTPClient, mockIDGenerator)

	ctx := context.Background()
	err := usecase.CheckTarget(ctx, "target-1")

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if len(mockResultRepo.SavedResults) == 0 {
		t.Error("Result should be saved")
	}

	if mockResultRepo.SavedResults[0].Status != "OK" {
		t.Error("Wrong status")
	}

	if mockResultRepo.SavedResults[0].StatusCode != 200 {
		t.Error("Wrong status code")
	}

	if len(mockAlertRepo.SavedAlerts) != 0 {
		t.Error("List with alerts should be empty")
	}
}

func TestCheckTarget_StatusError(t *testing.T) {
	mockTargetRepo := newMockTargetRepository()
	mockResultRepo := newMockResultRepository()
	mockAlertRepo := newMockAlertRepository()
	mockHTTPClient := newMockHTTPClient()
	mockIDGenerator := newMockIDGenerator()

	mockHTTPClient.CheckFunc = func(ctx context.Context, url string) (*domain.HTTPResponse, error) {
		return &domain.HTTPResponse{
			StatusCode:   500,
			ResponseTime: 30 * time.Millisecond,
		}, nil
	}

	usecase := NewMonitorUseCase(mockTargetRepo, mockResultRepo, mockAlertRepo, mockHTTPClient, mockIDGenerator)

	ctx := context.Background()
	err := usecase.CheckTarget(ctx, "target-1")

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if len(mockResultRepo.SavedResults) != 1 {
		t.Error("Length of returned result should be 1")
	}

	if mockResultRepo.SavedResults[0].Status != "SERVER_ERROR" {
		t.Error("Wrong status")
	}

	if mockResultRepo.SavedResults[0].StatusCode != 500 {
		t.Error("Wrong status code")
	}
}

func TestCheckTarget_AlertCreated(t *testing.T) {
	mockTargetRepo := newMockTargetRepository()
	mockResultRepo := newMockResultRepository()
	mockAlertRepo := newMockAlertRepository()
	mockHTTPClient := newMockHTTPClient()
	mockIDGenerator := newMockIDGenerator()

	mockHTTPClient.CheckFunc = func(ctx context.Context, url string) (*domain.HTTPResponse, error) {
		return &domain.HTTPResponse{
			StatusCode:   500,
			ResponseTime: 30 * time.Millisecond,
		}, nil
	}

	usecase := NewMonitorUseCase(mockTargetRepo, mockResultRepo, mockAlertRepo, mockHTTPClient, mockIDGenerator)

	ctx := context.Background()
	err := usecase.CheckTarget(ctx, "target-1")

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if len(mockResultRepo.SavedResults) != 1 {
		t.Error("Length of saved result should be 1")
	}	

	if len(mockAlertRepo.SavedAlerts) != 1 {
		t.Error("Lenght of saved alesrts should be 1")
	}

	if mockAlertRepo.SavedAlerts[0].Type != "SERVER_ERROR" {
		t.Error("Wrong alert type")
	}

	if mockAlertRepo.SavedAlerts[0].IsResolved != false {
		t.Error("Should return unResolved")
	}
}

func TestCheckTarget_AlertResolved(t *testing.T) {
	mockTargetRepo := newMockTargetRepository()
	mockResultRepo := newMockResultRepository()
	mockAlertRepo := newMockAlertRepository()
	mockHTTPClient := newMockHTTPClient()
	mockIDGenerator := newMockIDGenerator()

	mockAlertRepo.GetUnresolvedByTargetIDFunc = func(targetID string) ([]*domain.Alert, error) {
		return []*domain.Alert{
			{ID: "gen-alert-1", TargetID: "target-1", Type: "Error", Message: "Internal Server Error"},
		}, nil
	}

	usecase := NewMonitorUseCase(mockTargetRepo, mockResultRepo, mockAlertRepo, mockHTTPClient, mockIDGenerator)

	ctx := context.Background()
	err := usecase.CheckTarget(ctx, "target-1")

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if len(mockAlertRepo.UpdatedAlerts) != 1 {
		t.Error("Alert should be updated once")
	}

	if mockAlertRepo.UpdatedAlerts[0].IsResolved != true {
		t.Error("Alert after update (resolve) should get true")
	}

	if mockAlertRepo.UpdatedAlerts[0].ResolvedAt == nil {
		t.Error("Resolved alerts should get time of resolve")
	}
}