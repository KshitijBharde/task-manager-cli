package repository

import (
	"context"

	"github.com/KshitijBharde/task-manager-cli/internal/common"
	"github.com/KshitijBharde/task-manager-cli/internal/models"
)

// TaskHistoryRepository defines the interface for task history operations
type TaskHistoryRepository interface {
	// Create creates a new task history record
	Create(ctx context.Context, history *models.TaskHistory) error

	// GetByTaskID retrieves history records for a specific task ordered by version DESC (latest first)
	// Uses pagination with default limit of 10
	GetByTaskID(ctx context.Context, taskID int, params common.PaginationParams) ([]*models.TaskHistory, error)

	// GetByVersion retrieves a specific version of a task
	GetByVersion(ctx context.Context, taskID, version int) (*models.TaskHistory, error)

	// GetLatestVersion retrieves the latest version number for a task
	// Returns 0 if no history exists for the task
	GetLatestVersion(ctx context.Context, taskID int) (int, error)

	// ListVersions retrieves all version numbers for a task ordered DESC (latest first)
	ListVersions(ctx context.Context, taskID int) ([]int, error)

	// Delete deletes all history records for a specific task
	// This should typically only be used when hard deleting a task
	Delete(ctx context.Context, taskID int) error
}
