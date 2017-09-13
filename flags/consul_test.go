package flags

import (
	"testing"
	"flag"
	"github.com/cheebo/go-config/flags"
	"github.com/cheebo/go-config/types"
	"github.com/stretchr/testify/assert"
)

func TestDatabaseFlags(t *testing.T) {
	assert := assert.New(t)

	var (
		addr = "localhost:8500"
		scheme = "http"
		dc = "dc2"
		token = "secret"
		cafile = "/cafile/file"
		capath = "/capth"
		certfile = "/ca/cert"
		keyfile = "/ca/key"
	)

	args := []string{
		"-consul.addr="+addr,
		"-consul.scheme="+scheme,
		"-consul.dc="+dc,
		"-consul.token="+token,
		"-consul.tls.cafile="+cafile,
		"-consul.tls.capath="+capath,
		"-consul.tls.certfile="+certfile,
		"-consul.tls.keyfile="+keyfile,
	}

	f := flag.NewFlagSet("", flag.ContinueOnError)
	cfg := &types.ConsulConfig{}
	flags.ConsulFlags(cfg, f)

	err := f.Parse(args)

	assert.NoError(err)
	assert.Equal(addr, cfg.Addr)
	assert.Equal(scheme, cfg.Scheme)
	assert.Equal(dc, cfg.Datacenter)
	assert.Equal(token, cfg.Token)
	assert.Equal(cafile, cfg.Tls.CAFile)
	assert.Equal(capath, cfg.Tls.CAPath)
	assert.Equal(certfile, cfg.Tls.CertFile)
	assert.Equal(keyfile, cfg.Tls.KeyFile)
}