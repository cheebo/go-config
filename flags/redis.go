package flags

import (
	"flag"
	"github.com/cheebo/go-config/types"
)

func RedisFlags(cfg *types.RedisConfig, f *flag.FlagSet) {
	f.StringVar(&cfg.Host, "redis.host", "localhost", "redis host: 'localhost' or '127.0.0.1'")
	f.UintVar(&cfg.Port, "redis.port", 6379, "port number")
	f.StringVar(&cfg.Password, "redis.password", "", "password")
	f.UintVar(&cfg.Database, "redis.database", 0, "port number")
	f.UintVar(&cfg.PoolSize, "redis.poolsize", 10, "size of the connection pool, default 10")
	f.StringVar(&cfg.MasterName, "redis.mastername", "", "master name")
	f.BoolVar(&cfg.SlaveReadOnly, "redis.slave_read_only", true, "slave read only, default true")
}
