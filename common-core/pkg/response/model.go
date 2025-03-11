package response

import "github.com/Mohammed-Aadil/common-core/pkg/pagination"

type HttpPaginatedResponse struct {
	Data       any
	Pagination pagination.Pagination
}

type HttpErrorResponse struct {
	Message string `json:"message"`
	Details string `json:"details"`
}

type RuntimeResponse struct {
	Hostname     string `json:"hostname"`
	GOOS         string `json:"goos"`
	GOARCH       string `json:"goarch"`
	Runtime      string `json:"runtime"`
	NumGoroutine int    `json:"numgoroutine"`
	NumCPU       int    `json:"numcpu"`
}
