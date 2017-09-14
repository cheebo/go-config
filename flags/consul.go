package flags

import (
	"flag"
	"github.com/cheebo/go-config/types"
)

func ConsulFlags(cfg *types.ConsulConfig, f *flag.FlagSet) {
	f.StringVar(&cfg.Addr, "consul.addr", cfg.Addr, "Consul HTTP addr, default: localhost:8500")
	f.StringVar(&cfg.Scheme, "consul.scheme", cfg.Scheme, "Consul scheme, default: http")
	f.StringVar(&cfg.Datacenter, "consul.dc", cfg.Datacenter, "Consul datacenter, default: dc1")
	f.StringVar(&cfg.Token, "consul.token", cfg.Token, "Consul token")

	f.StringVar(&cfg.Tls.CAFile, "consul.tls.cafile", cfg.Tls.CAFile, "Path to the CA certificate used for Consul communication")
	f.StringVar(&cfg.Tls.CAPath, "consul.tls.capath", cfg.Tls.CAPath, "Path to a directory of CA certificates to use for Consul communication")
	f.StringVar(&cfg.Tls.CertFile, "consul.tls.certfile", cfg.Tls.CertFile, "Path to the certificate for Consul communication. If this is set then you need to also set KeyFile.")
	f.StringVar(&cfg.Tls.KeyFile, "consul.tls.keyfile", cfg.Tls.KeyFile, "Path to the private key for Consul communication. If this is set then you need to also set CertFile.")
}