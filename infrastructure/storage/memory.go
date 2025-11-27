package storage

import (
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
