package domain

type TargetRepository interface {
	Save(target *Target) error
	FindByID(id string) (*Target, error)
	GetAll() ([]*Target, error)
	Delete(id string) error
	Update(target *Target) error
}

type ResultRepository interface {
	Save(result *Result) error
	FindByTargetID(targetID string) ([]*Result, error)
	GetLastByTargetID(targetID string) (*Result, error)
}

type AlertRepository interface {
	Save(alert *Alert) error
	FindByTargetID(targetID string) ([]*Alert, error)
	GetUnresolvedByTargetID(targetID string) ([]*Alert, error)
	Update(alert *Alert) error
}
