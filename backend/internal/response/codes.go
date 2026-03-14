package response

// エラーコード
const (
	CodeInvalidRequest  = "invalid_request"
	CodeInvalidEmail    = "invalid_email"
	CodeInvalidPassword = "invalid_password"
	CodeUnauthorized    = "unauthorized"
	CodeInvalidToken    = "invalid_token"
	CodeNotFound        = "not_found"
	CodeConflict        = "conflict"
	CodeDatabaseError   = "database_error"
	CodeInternalError   = "internal_error"
)

// 成功
const (
	CodeOK = "ok"
)
