package flags_test

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
		smtpHost = "localhost"
		smtpPort = uint(25)
		smtpUser = "admin"
		smtpPassword = "password"
		smtpSSL = false
		smtpLocalname = "localname"
	)

	args := []string{
		"-smtp.host="+smtpHost,
		"-smtp.port=25",
		"-smtp.user="+smtpUser,
		"-smtp.password="+smtpPassword,
		"-smtp.ssl=false",
		"-smtp.localname="+smtpLocalname,
	}

	f := flag.NewFlagSet("", flag.ContinueOnError)
	cfg := &types.SMTPConfig{}
	flags.SmtpFlags(cfg, f)

	err := f.Parse(args)

	assert.NoError(err)
	assert.Equal(smtpHost, cfg.Host)
	assert.Equal(smtpPort, cfg.Port)
	assert.Equal(smtpUser, cfg.User)
	assert.Equal(smtpPassword, cfg.Password)
	assert.Equal(smtpSSL, cfg.SSL)
	assert.Equal(smtpLocalname, cfg.LocalName)
}
