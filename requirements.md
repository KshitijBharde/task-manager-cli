# Task Manager CLI - Requirements Specification

## Functional Requirements

### Core Task Management

**FR1: Add Task**
- User can create a new task with title, description (optional), and priority (low/medium/high)
- System assigns unique ID and creation timestamp automatically
- Task defaults to "pending" status
- Priority defaults to "medium" if not specified

**FR2: List Tasks**
- Display all tasks with ID, title, status, priority, and creation date
- Support filtering by status (pending/in-progress/completed)
- Support filtering by priority (low/medium/high)
- Display tasks sorted by creation date (newest first) by default
- Show task count at the bottom

**FR3: Update Task**
- Modify task title, description, priority, or status by ID
- Support partial updates (only change specified fields)
- Update last modified timestamp automatically

**FR4: Delete Task**
- Remove task by ID
- Request confirmation before deletion
- Display success/failure message

**FR5: Mark Task Status**
- Quick commands to mark task as:
  - In progress
  - Completed
  - Pending (reopen)
- Update completion timestamp when marked complete

**FR6: Search Tasks**
- Search tasks by keyword in title or description
- Case-insensitive search
- Display matching results with highlighting

**FR7: View Task Details**
- Show full details of a single task by ID
- Include all fields: title, description, status, priority, created date, modified date, completed date

### Data Persistence

**FR8: JSON Storage**
- Store all tasks in `~/.taskmanager/tasks.json` file
- Create directory and file if they don't exist
- Maintain valid JSON structure at all times

**FR9: Data Integrity**
- Validate JSON structure on load
- Handle corrupted file gracefully with error message
- Provide backup functionality (`task backup`)
- Provide restore from backup (`task restore`)

### CLI Interface

**FR10: Help System**
- Display usage information with `task help` or `task --help`
- Show command-specific help with `task <command> --help`
- Include examples for each command

**FR11: Version Information**
- Display version with `task version` or `task --version`

## Non-Functional Requirements

### Performance

**NFR1: Response Time**
- All commands execute in under 100ms for datasets up to 10,000 tasks
- List command renders results within 200ms

**NFR2: File Size Handling**
- Support up to 50,000 tasks without performance degradation
- Gracefully handle file size limits

### Usability

**NFR3: User-Friendly Output**
- Use colors for different priorities and statuses
- Format output in readable tables
- Clear error messages with actionable suggestions
- Success confirmations for all operations

**NFR4: Intuitive Commands**
- Follow common CLI conventions (ls, rm, etc.)
- Support both long flags (`--status`) and short flags (`-s`)
- Consistent command structure across all operations

### Reliability

**NFR5: Error Handling**
- Handle all file I/O errors gracefully
- Validate user input before processing
- Never crash - always provide meaningful error messages
- Atomic writes to prevent data corruption

**NFR6: Data Safety**
- Create backup before destructive operations (optional flag)
- Validate JSON before writing
- Use file locking for concurrent access protection

### Maintainability

**NFR7: Code Quality**
- Follow Go best practices and idioms
- Comprehensive error handling
- Package structure: cmd/, internal/, pkg/
- Unit tests for core logic (>70% coverage)
- Integration tests for CLI commands

**NFR8: Documentation**
- Inline code comments for complex logic
- README with installation and usage instructions
- Architecture decision records for key design choices

### Portability

**NFR9: Cross-Platform Support**
- Work on Linux, macOS, and Windows
- Handle path separators correctly
- Use cross-platform terminal color libraries

**NFR10: Installation**
- Single binary with no dependencies
- Install via `go install` command
- Provide pre-built binaries for major platforms

### Security

**NFR11: Data Privacy**
- Store tasks in user's home directory only
- File permissions set to user read/write only (0600)
- No network calls or external dependencies

## Command Structure Examples

```bash
# Add tasks
task add "Buy groceries" --priority high
task add "Write blog post" --desc "About Go interfaces" -p medium

# List and filter
task list
task list --status pending
task list --priority high
task ls -s completed

# Update tasks
task update 1 --title "Buy groceries and cook"
task update 2 --priority low --status in-progress

# Quick status changes
task start 3        # Mark as in-progress
task done 1         # Mark as completed
task reopen 5       # Mark as pending

# Search and view
task search "groceries"
task show 1

# Delete
task delete 4
task rm 2

# Maintenance
task backup
task restore backup-2024-12-01.json
task stats          # Show statistics

# Help
task help
task add --help
```

## Data Model

```json
{
  "tasks": [
    {
      "id": 1,
      "title": "Buy groceries",
      "description": "Milk, eggs, bread",
      "status": "pending",
      "priority": "high",
      "created_at": "2024-12-01T10:30:00Z",
      "updated_at": "2024-12-01T10:30:00Z",
      "completed_at": null
    }
  ],
  "next_id": 2
}
```

## Success Metrics

- Complete all functional requirements
- Unit test coverage >70%
- Zero crashes on invalid input
- Commands complete in <100ms
- Successfully handle 10,000+ tasks
- Positive user feedback on usability