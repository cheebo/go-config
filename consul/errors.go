package consul

var (
	ErrCantGetConfig = errors.New("Can't get configuration from consul")
	ErrCantUnmarshalJSON = errors.New("Can't unmarshal JSON configuration string")
)
