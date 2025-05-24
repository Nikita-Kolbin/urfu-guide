package model

type ErrorResponse struct {
	Err string `json:"error"`
}

type HealthResponse struct {
	Success bool `json:"success"`
}

func NewErrResp(err string) *ErrorResponse {
	return &ErrorResponse{Err: err}
}
