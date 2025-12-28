package customerrors

import "errors"

var (
	ErrTaskNotFound    = errors.New("task not found")
	ErrTitleIsRequired = errors.New("title is required")
	ErrInvalidData     = errors.New("invalid task data")
	ErrInvalidID       = errors.New("invalid task id")
)
