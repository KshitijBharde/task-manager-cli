package repository

import (
	"context"

	"github.com/KshitijBharde/task-manager-cli/internal/common"
	"github.com/KshitijBharde/task-manager-cli/internal/models"
)

// TaskRepository defines the interface for task data operations
type TaskRepository interface {
	// Create creates a new task and returns it with the generated ID
	Create(ctx context.Context, task *models.Task) (*models.Task, error)

	// Update updates an existing task, increments its version, and returns the updated task
	Update(ctx context.Context, task *models.Task) (*models.Task, error)

	// Delete soft deletes a task by setting deletedAt timestamp
	// Returns true if the task was deleted, false if it didn't exist
	Delete(ctx context.Context, id int) (bool, error)

	// GetByID retrieves a task by its ID (excludes soft deleted)
	GetByID(ctx context.Context, id int) (*models.Task, error)

	// List retrieves non-deleted tasks with optional filters, ordered by updatedAt
	// Supports filtering by status, priority ranges, and date ranges
	// Uses pagination with default limit of 10 and order DESC
	List(ctx context.Context, filter *common.TaskFilter, params common.PaginationParams) ([]*models.Task, error)

	// ListDeleted retrieves soft deleted tasks ordered by deletedAt
	// Uses pagination with default limit of 10 and order DESC
	ListDeleted(ctx context.Context, params common.PaginationParams) ([]*models.Task, error)

	// Restore restores a soft deleted task
	// Returns true if the task was restored, false if it didn't exist
	Restore(ctx context.Context, id int) (bool, error)

	// HardDelete permanently deletes a task from the database
	// Returns true if the task was deleted, false if it didn't exist
	HardDelete(ctx context.Context, id int) (bool, error)
}
