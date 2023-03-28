package svcerr

const (
	InvalidCredentials         = "invalid username or password"
	InvalidRegisterCredentials = "username, password and Email are required"
	NotRegistered              = "not registered"
	Unauthorized               = "unauthorized"
	InternalServerError        = "internal server error"
	NotFound                   = "resource not found"
	AccessDenied               = "access denied"
	BadRequest                 = "bad request"
	DuplicateEntry             = "duplicate entry"
	InvalidAuthentication      = "invalid authentication"
	UserAlreadyExists          = "user already exists"
)
