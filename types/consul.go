package types

type ConsulTlsConfig struct {
	CAFile          string  `json:"cafile"`
	CAPath          string  `json:"capath"`
	CertFile        string  `json:"certfile"`
	KeyFile         string  `json:"keyfile"`
}

type ConsulConfig struct {
	Addr            string  `json:"addr"`
	Datacenter      string  `json:"datacenter"`
	Token           string  `json:"token"`
	Schema          string  `json:"schema"`
	Tls             ConsulTlsConfig  `json:"tls"`
}