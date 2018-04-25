package types

type ConsulTlsConfig struct {
	CAFile   string `json:"cafile" cfg:"cafile"`
	CAPath   string `json:"capath" cfg:"capath"`
	CertFile string `json:"certfile" cfg:"certfile"`
	KeyFile  string `json:"keyfile" cfg:"keyfile"`
}

type ConsulConfig struct {
	Addr       string          `json:"addr" default:"localhost:8500"`
	Datacenter string          `json:"dc" cfg:"dc" default:"dc1"`
	Token      string          `json:"token"`
	Scheme     string          `json:"scheme" default:"http"`
	Tls        ConsulTlsConfig `json:"tls"`
}
