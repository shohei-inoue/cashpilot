package response

// エラーメッセージ
const (
	MsgInvalidRequest      = "invalid request"
	MsgInvalidEmail        = "invalid email"
	MsgInvalidPassword     = "password must be at least 8 characters"
	MsgUnauthorized        = "unauthorized"
	MsgInvalidToken        = "invalid token"
	MsgInvalidCredentials  = "invalid email or password"
	MsgUserNotFound        = "user not found"
	MsgEmailAlreadyExists  = "email already registered"
	MsgDatabaseError       = "database error"
	MsgInternalError       = "internal server error"
	MsgFailedToHash        = "failed to hash password"
	MsgFailedToCreateUser  = "failed to create user"
	MsgFailedToIssueToken  = "failed to issue token"
	MsgDatabaseUnreachable = "database unreachable"
)

// 成功
const (
	MsgOK     = "ok"
	MsgStatus = "status"
)
