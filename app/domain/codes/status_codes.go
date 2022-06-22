package codes

type ErrorCode string

const (
	OK             ErrorCode = "OK"
	BadParams      ErrorCode = "bad_params"
	EmptyBody      ErrorCode = "empty_body"
	InvalidRequest ErrorCode = "invalid_request"
	Unauthorized   ErrorCode = "unauthorized"
	NotFound       ErrorCode = "not_found"
	Database       ErrorCode = "database_error"
	Internal       ErrorCode = "internal_error"
	Forbidden      ErrorCode = "forbidden"
	Unknown        ErrorCode = "unknown"
)
