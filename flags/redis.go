package flags

import (
	"flag"
	"github.com/cheebo/go-config/types"
)

func RedisFlags(cfg *types.RedisConfig, f *flag.FlagSet) {
	f.StringVar(&cfg.Addr, "redis.addr", cfg.Addr, "redis addr")
	f.StringVar(&cfg.Password, "redis.password", cfg.Password, "password")
	f.UintVar(&cfg.Database, "redis.database", cfg.Database, "port number")
	f.UintVar(&cfg.PoolSize, "redis.poolsize", cfg.PoolSize, "size of the connection pool, default 10")
	f.StringVar(&cfg.MasterName, "redis.mastername", cfg.MasterName, "master name")
	f.BoolVar(&cfg.SlaveReadOnly, "redis.slave_read_only", cfg.SlaveReadOnly, "slave read only, default true")
}
