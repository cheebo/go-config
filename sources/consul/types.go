package consul

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
	Scheme          string  `json:"scheme"`
	Tls             ConsulTlsConfig  `json:"tls"`
}
