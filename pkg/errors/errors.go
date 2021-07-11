package errors

import "errors"

var (
	NoVariablesInitialised = errors.New("no variables initialised")
	NotAStructPtr          = errors.New("expects pointer to a struct")
	NotAStruct             = errors.New("expects a struct")
)
