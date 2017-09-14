package flags

import (
	"flag"
	"github.com/cheebo/go-config/types"
)

func RedisFlags(cfg *types.RedisConfig, f *flag.FlagSet) {
	f.StringVar(&cfg.Host, "redis.host", cfg.Host, "redis host: 'localhost' or '127.0.0.1'")
	f.UintVar(&cfg.Port, "redis.port", cfg.Port, "port number")
	f.StringVar(&cfg.Password, "redis.password", cfg.Password, "password")
	f.UintVar(&cfg.Database, "redis.database", cfg.Database, "port number")
	f.UintVar(&cfg.PoolSize, "redis.poolsize", cfg.PoolSize, "size of the connection pool, default 10")
	f.StringVar(&cfg.MasterName, "redis.mastername", cfg.MasterName, "master name")
	f.BoolVar(&cfg.SlaveReadOnly, "redis.slave_read_only", cfg.SlaveReadOnly, "slave read only, default true")
}
