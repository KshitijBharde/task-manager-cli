package models

import (
	"encoding/json"
	"time"
)

// Operation constants for task history operations
const (
	OperationCreate = "CREATE"
	OperationUpdate = "UPDATE"
	OperationDelete = "DELETE"
)

// TaskHistory represents a historical record of a task's state
type TaskHistory struct {
	ID        int       `json:"id" db:"history_id"`
	TaskID    int       `json:"taskId" db:"task_id"`
	Version   int       `json:"version" db:"version"`
	Snapshot  Task      `json:"snapshot" db:"snapshot"`   // Full task snapshot
	Operation string    `json:"operation" db:"operation"` // CREATE, UPDATE, DELETE
	CreatedAt time.Time `json:"createdAt" db:"createdAt"`
}

// IsValidOperation returns true if the operation is valid
func (th *TaskHistory) IsValidOperation() bool {
	return th.Operation == OperationCreate ||
		th.Operation == OperationUpdate ||
		th.Operation == OperationDelete
}

// FromTask creates a TaskHistory record from a Task
func (th *TaskHistory) FromTask(task *Task, operation string) {
	th.TaskID = task.ID
	th.Version = task.Version
	th.Snapshot = *task // Store full task snapshot
	th.Operation = operation
	th.CreatedAt = time.Now()
}

// MarshalSnapshot converts the task snapshot to JSON for database storage
func (th *TaskHistory) MarshalSnapshot() ([]byte, error) {
	return json.Marshal(th.Snapshot)
}

// UnmarshalSnapshot converts JSON from database back to Task struct
func (th *TaskHistory) UnmarshalSnapshot(data []byte) error {
	return json.Unmarshal(data, &th.Snapshot)
}
