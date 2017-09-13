package flags

import (
	"flag"
	"github.com/cheebo/go-config/types"
)

func ConsulFlags(cfg *types.ConsulConfig, f *flag.FlagSet) {
	f.StringVar(&cfg.Addr, "consul.addr", "http://localhost:8500", "Consul schema and addr, default: http://localhost:8500")
	f.StringVar(&cfg.Datacenter, "consul.dc", "dc1", "Consul datacenter, default: dc1")
	f.StringVar(&cfg.Token, "consul.token", "", "Consul token")

	f.StringVar(&cfg.Tls.CAFile, "consul.tls.cafile", "", "Path to the CA certificate used for Consul communication")
	f.StringVar(&cfg.Tls.CAPath, "consul.tls.capath", "", "Path to a directory of CA certificates to use for Consul communication")
	f.StringVar(&cfg.Tls.CertFile, "consul.tls.certfile", "", "Path to the certificate for Consul communication. If this is set then you need to also set KeyFile.")
	f.StringVar(&cfg.Tls.KeyFile, "consul.tls.keyfile", "", "Path to the private key for Consul communication. If this is set then you need to also set CertFile.")
}