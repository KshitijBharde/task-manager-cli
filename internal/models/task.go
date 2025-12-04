package models

import "time"

// Status constants for task status
const (
	StatusBacklog    = "BACKLOG"
	StatusPending    = "PENDING"
	StatusInProgress = "IN_PROGRESS"
	StatusCompleted  = "COMPLETED"
)

// Priority constants for task prioritization
const (
	PriorityCritical = 0  // Must do now, blocks everything
	PriorityHigh     = 10 // Important, do soon
	PriorityMedium   = 20 // Normal priority
	PriorityLow      = 30 // Can wait
	PriorityBacklog  = 40 // No priority set / backlog
)

// Task represents a task entity in the task manager
type Task struct {
	ID          int        `json:"id" db:"id"`
	Title       string     `json:"title" db:"title"`
	Description string     `json:"description" db:"description"`
	Status      string     `json:"status" db:"status"`           // BACKLOG, PENDING, IN_PROGRESS, COMPLETED
	Priority    int        `json:"priority" db:"priority"`
	Version     int        `json:"version" db:"version"`         // Current version, increments on each update
	CreatedAt   time.Time  `json:"createdAt" db:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt" db:"updatedAt"`
	DeletedAt   *time.Time `json:"deletedAt,omitempty" db:"deletedAt"` // Soft delete
}

// IsDeleted returns true if the task has been soft deleted
func (t *Task) IsDeleted() bool {
	return t.DeletedAt != nil
}

// IsValidStatus returns true if the status is valid
func (t *Task) IsValidStatus() bool {
	return t.Status == StatusBacklog ||
		t.Status == StatusPending ||
		t.Status == StatusInProgress ||
		t.Status == StatusCompleted
}

// PriorityName returns the human-readable name for the task's priority
func (t *Task) PriorityName() string {
	switch t.Priority {
	case PriorityCritical:
		return "critical"
	case PriorityHigh:
		return "high"
	case PriorityMedium:
		return "medium"
	case PriorityLow:
		return "low"
	case PriorityBacklog:
		return "backlog"
	default:
		return "unknown"
	}
}
