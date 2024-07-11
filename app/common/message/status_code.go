package message

type StatusCode struct {
	Code int
	Lang string
	Desc string
}

var (
	StatusOK                      = StatusCode{2000, "ok", ""}
	StatusFail                    = StatusCode{2001, "fail", ""}
	StatusParameterError          = StatusCode{2002, "parameter error", ""}
	StatusUnAuthorized            = StatusCode{2003, "unauthorized", ""}
	StatusTokenCannotBeEmpty      = StatusCode{2004, "token cannot be empty", ""}
	StatusTokenIsInvalid          = StatusCode{2005, "token is invalid", ""}
	StatusWrongUsernameOrPassword = StatusCode{2006, "wrong user name or password", ""}

	StatusInternalError = StatusCode{5000, "internal error", ""}
)
