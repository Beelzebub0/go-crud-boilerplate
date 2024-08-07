package entity

import "time"

type Pagination struct {
	CurrentPage     int64 `json:"current_page"`
	CurrentElements int64 `json:"current_elements"`
	TotalPages      int64 `json:"total_pages"`
	TotalElements   int64 `json:"total_elements"`
}

type HttpResponse struct {
	Code       int           `json:"code" example:"200"`
	Data       interface{}   `json:"data,omitempty"`
	Pagination *Pagination   `json:"pagination,omitempty"`
	Message    string        `json:"message" example:"string"`
	Metadata   *HttpMetadata `json:"metadata,omitempty"`
}

type HttpMetadata struct {
	Error *HttpError `json:"error,omitempty"`
}

type HttpError struct {
	SystemMessage  string `json:"system_message,omitempty"`
	Stacktrace     string `json:"stacktrace,omitempty"`
	FullStacktrace string `json:"full_stacktrace,omitempty"`
}

type CreatedAndUpdatedTime struct {
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
