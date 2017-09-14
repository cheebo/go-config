package flags

import (
	"flag"
	"github.com/cheebo/go-config/types"
)

func DatabaseFlags(cfg *types.DatabaseConfig, f *flag.FlagSet) {
	f.StringVar(&cfg.Host, "db.host", cfg.Host, "database host: 'localhost' or '127.0.0.1'")
	f.UintVar(&cfg.Port, "db.port", cfg.Port, "port number")
	f.StringVar(&cfg.User, "db.user", cfg.User, "user's name")
	f.StringVar(&cfg.Password, "db.password", cfg.Password, "user's password")
	f.StringVar(&cfg.Database, "db.database", cfg.Database, "database name")
	f.StringVar(&cfg.Driver, "db.driver", cfg.Driver, "database driver")
}