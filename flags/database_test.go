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
		dbHost = "localhost"
		dbPort = uint(5432)
		dbUser = "admin"
		dbPassword = "password"
		dbDatabase = "template1"
		dbDriver = "postgresql"
	)

	args := []string{
		"-db.host="+dbHost,
		"-db.port=5432",
		"-db.user="+dbUser,
		"-db.password="+dbPassword,
		"-db.database="+dbDatabase,
		"-db.driver="+dbDriver,
	}

	f := flag.NewFlagSet("", flag.ContinueOnError)
	cfg := &types.DatabaseConfig{}
	flags.DatabaseFlags(cfg, f)

	err := f.Parse(args)

	assert.NoError(err)
	assert.Equal(dbHost, cfg.Host)
	assert.Equal(dbPort, cfg.Port)
	assert.Equal(dbUser, cfg.User)
	assert.Equal(dbPassword, cfg.Password)
	assert.Equal(dbDatabase, cfg.Database)
	assert.Equal(dbDriver, cfg.Driver)
}

