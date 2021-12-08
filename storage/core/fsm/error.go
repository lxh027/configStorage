package fsm

import "errors"

var (
	ErrPathNotMatch   = errors.New("value path doesn't match")
	ErrItemExisted    = errors.New("item key existed")
	ErrItemNotExisted = errors.New("item key not existed")

	ErrWrongArgs          = errors.New("wrong args passed")
	ErrOperationUndefined = errors.New("operation is not defined")

	ErrCommitNotFound = errors.New("can't find the committed id")
)
