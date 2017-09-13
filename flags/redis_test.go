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
		rHost = "localhost"
		rPort = uint(6379)
		rPassword = "password"
		rDatabase = uint(0)
		rPoolSize = uint(10)
		rMasterName = "master"
		rSlaveReadOnly = true
	)

	args := []string{
		"-redis.host="+rHost,
		"-redis.port=6379",
		"-redis.password="+rPassword,
		"-redis.database=0",
		"-redis.poolsize=10",
		"-redis.mastername="+rMasterName,
		"-redis.slave_read_only=true",
	}

	f := flag.NewFlagSet("", flag.ContinueOnError)
	cfg := &types.RedisConfig{}
	flags.RedisFlags(cfg, f)

	err := f.Parse(args)

	assert.NoError(err)
	assert.Equal(rHost, cfg.Host)
	assert.Equal(rPort, cfg.Port)
	assert.Equal(rPassword, cfg.Password)
	assert.Equal(rDatabase, cfg.Database)
	assert.Equal(rPoolSize, cfg.PoolSize)
	assert.Equal(rMasterName, cfg.MasterName)
	assert.Equal(rSlaveReadOnly, cfg.SlaveReadOnly)
}
