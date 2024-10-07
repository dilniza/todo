package models

type ErrorResponse struct {
    Error   string `json:"error"`
    Message string `json:"message,omitempty"`
}

type SuccessResponse struct {
    Message string `json:"message"`
}

type PagedResponse struct {
    Data       interface{} `json:"data"`
    TotalCount int64       `json:"total_count"`
    Page       int64       `json:"page"`
    Limit      int64       `json:"limit"`
}
