package storage

import (
	"fmt"
	"sync"

	"github.com/karoljaro/go-uptime-monitor/domain"
)

// ========== [TARGET] ==========

type MemoryTargetRepository struct {
	mu      sync.RWMutex
	targets map[string]*domain.Target
}

func NewMemoryTargetRepository() *MemoryTargetRepository {
	return &MemoryTargetRepository{
		targets: make(map[string]*domain.Target),
	}
}

func (r *MemoryTargetRepository) Save(target *domain.Target) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.targets[target.ID] = target

	return nil
}

func (r *MemoryTargetRepository) FindByID(id string) (*domain.Target, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	val, exists := r.targets[id]

	if !exists {
		return nil, fmt.Errorf("target with id: %s not found", id)
	}

	return val, nil
}

func (r *MemoryTargetRepository) GetAll() ([]*domain.Target, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	targets := make([]*domain.Target, 0, len(r.targets))

	for _, target := range r.targets {
		targets = append(targets, target)
	}

	return targets, nil
}

func (r *MemoryTargetRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.targets[id]; !exists {
		return fmt.Errorf("target with id: %s not found", id)
	}

	delete(r.targets, id)
	return nil
}

func (r *MemoryTargetRepository) Update(target *domain.Target) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.targets[target.ID]; !exists {
		return fmt.Errorf("target with id: %s not found", target.ID)
	}

	r.targets[target.ID] = target
	return nil
}

// ========== [ALERT] ==========

type MemoryAlertRepository struct {
	mu     sync.RWMutex
	alerts map[string][]*domain.Alert
}

func NewMemoryAlertRepository() *MemoryAlertRepository {
	return &MemoryAlertRepository{
		alerts: make(map[string][]*domain.Alert),
	}
}

func (r *MemoryAlertRepository) Save(alert *domain.Alert) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.alerts[alert.TargetID] = append(r.alerts[alert.TargetID], alert)
	return nil
}

func (r *MemoryAlertRepository) FindByTargetID(targetID string) ([]*domain.Alert, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if val, exists := r.alerts[targetID]; exists {
		return val, nil
	}

	return nil, fmt.Errorf("alerts with targetID: %s not found", targetID)
}

func (r  *MemoryAlertRepository) GetUnresolvedByTargetID(targetID string) ([]*domain.Alert, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	alerts := make([]*domain.Alert, 0, len(r.alerts[targetID]))

	if len(r.alerts[targetID]) == 0 {
		return nil, fmt.Errorf("alert with id: %s not exists", targetID)
	}

	for _, alert := range r.alerts[targetID] {
		if !alert.IsResolved {
			alerts = append(alerts, alert)
		}
	}

	return alerts, nil
}

// ========== [RESULT] ==========

type MemoryResultRepository struct {
	mu      sync.RWMutex
	results map[string][]*domain.Result
}

func NewMemoryResultRepository() *MemoryResultRepository {
	return &MemoryResultRepository{
		results: make(map[string][]*domain.Result),
	}
}

func (r *MemoryResultRepository) Save(result *domain.Result) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.results[result.TargetID] = append(r.results[result.TargetID], result)
	return nil
}

func (r *MemoryResultRepository) FindByTargetID(targetID string) ([]*domain.Result, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if val, exists := r.results[targetID]; exists {
		return val, nil
	} 
		
	return nil, fmt.Errorf("result with targetID: %s not found", targetID)
}

func (r *MemoryResultRepository) GetLastByTargetID(targetID string) (*domain.Result, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	val, exists := r.results[targetID]

	if !exists || len(val) == 0 {
		return nil, fmt.Errorf("last result with targetID %s not found", targetID)
	}

	return val[len(val)-1], nil
}
