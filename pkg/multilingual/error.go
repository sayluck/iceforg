package multilingual

import "errors"

var (
	// system error
	SystemOperationError = errors.New("systemOptionError")
	SystemPanicError     = errors.New("systemPanicError")

	// user error
	UserAlreadyExisted = errors.New("userAlreadyExisted")
	UserNotExisted     = errors.New("userNotExisted")
	UserLoginErr       = errors.New("userLoginErr")
	UserInvaildToken   = errors.New("userInvaildToken")
	UserTokenIsExpired = errors.New("userTokenIsExpired")
)
