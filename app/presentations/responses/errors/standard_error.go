package errors

type StandardErrorResponse struct {
	Error Error `json:"error"`
}

type Error struct {
	Message string `json:"message"`
	Detail  string `json:"detail"`
}
