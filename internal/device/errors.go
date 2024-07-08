package device

import "errors"

var (
	ErrProjectIDNotFound = errors.New("project id is not defined")
	ErrRepository        = errors.New("something when wrong on us, not you")
	ErrWhatsappClient    = errors.New("errors from whatsapp client")
	ErrWhatsappLibrary   = errors.New("errors from whatsmeow library")
	ErrAlreadyLoggedIn   = errors.New("client already logged in")
)
