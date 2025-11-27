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
	r.targets[target.ID] = target
	r.mu.Unlock()

	return nil
}

func (r *MemoryTargetRepository) FindByID(id string) (*domain.Target, error) {
	r.mu.RLock()
	val, exists := r.targets[id]
	r.mu.RUnlock()

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
