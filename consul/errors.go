package consul

import (
	"errors"
)

var (
	ErrCantGetConfig     = errors.New("Can't get configuration from consul")
	ErrCantUnmarshalJSON = errors.New("Can't unmarshal JSON configuration string")
	ErrCantPutConfig     = errors.New("Can't put configuration to consul")
	ErrCantMarshalJSON   = errors.New("Can't marshal struct to JSON configuration string")
)
