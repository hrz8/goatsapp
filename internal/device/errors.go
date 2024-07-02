package device

import "errors"

var (
	ErrProjectIDNotFound = errors.New("project id is not defined")
	ErrStoreDevice       = errors.New("something when wrong on us, not you")
)
