package pagination

type Pagination struct {
	Limit     int       `json:"limit" form:"limit"`
	Offset    int       `json:"offset" form:"offset"`
	Total     int       `json:"total"`
	SortField string    `json:"sortField" form:"sortField"`
	SortOrder SortOrder `json:"sortOrder" form:"sortOrder"`
}
