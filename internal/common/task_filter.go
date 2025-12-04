package common

import "time"

// TaskFilter defines filtering options for task queries
type TaskFilter struct {
	// Status filters
	Status *string // Exact status match

	// Priority filters
	PriorityEq  *int // Priority equals
	PriorityGte *int // Priority greater than or equal to
	PriorityLte *int // Priority less than or equal to

	// Date filters (date only, ignoring time component)
	CreatedAtDate *time.Time // Filter by creation date (YYYY-MM-DD)
	UpdatedAtDate *time.Time // Filter by update date (YYYY-MM-DD)

	// Date range filters
	CreatedAfter  *time.Time // Created after this date
	CreatedBefore *time.Time // Created before this date
	UpdatedAfter  *time.Time // Updated after this date
	UpdatedBefore *time.Time // Updated before this date
}

// NewTaskFilter returns an empty task filter
func NewTaskFilter() *TaskFilter {
	return &TaskFilter{}
}

// WithStatus sets the status filter
func (f *TaskFilter) WithStatus(status string) *TaskFilter {
	f.Status = &status
	return f
}

// WithPriorityEq sets the priority equals filter
func (f *TaskFilter) WithPriorityEq(priority int) *TaskFilter {
	f.PriorityEq = &priority
	return f
}

// WithPriorityGte sets the priority greater than or equal filter
func (f *TaskFilter) WithPriorityGte(priority int) *TaskFilter {
	f.PriorityGte = &priority
	return f
}

// WithPriorityLte sets the priority less than or equal filter
func (f *TaskFilter) WithPriorityLte(priority int) *TaskFilter {
	f.PriorityLte = &priority
	return f
}

// WithCreatedAtDate sets the created at date filter
func (f *TaskFilter) WithCreatedAtDate(date time.Time) *TaskFilter {
	f.CreatedAtDate = &date
	return f
}

// WithUpdatedAtDate sets the updated at date filter
func (f *TaskFilter) WithUpdatedAtDate(date time.Time) *TaskFilter {
	f.UpdatedAtDate = &date
	return f
}
