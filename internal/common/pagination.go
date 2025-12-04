package common

// SortOrder defines the sort order for queries
type SortOrder string

const (
	SortOrderAsc  SortOrder = "ASC"
	SortOrderDesc SortOrder = "DESC"
)

// PaginationParams defines pagination parameters for list operations
type PaginationParams struct {
	Limit  int       // Number of records to return (default: 10)
	Offset int       // Number of records to skip (default: 0)
	Order  SortOrder // Sort order: ASC or DESC (default: DESC)
}

// DefaultPagination returns default pagination parameters
func DefaultPagination() PaginationParams {
	return PaginationParams{
		Limit:  10,
		Offset: 0,
		Order:  SortOrderDesc,
	}
}
