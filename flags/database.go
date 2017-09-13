package flags

import (
	"flag"
	"github.com/cheebo/go-config/types"
)

func DatabaseFlags(cfg *types.DatabaseConfig, f *flag.FlagSet) {
	f.StringVar(&cfg.Host, "db.host", "", "database host: 'localhost' or '127.0.0.1'")
	f.UintVar(&cfg.Port, "db.port", 0, "port number")
	f.StringVar(&cfg.User, "db.user", "", "user's name")
	f.StringVar(&cfg.Password, "db.password", "", "user's password")
	f.StringVar(&cfg.Database, "db.database", "", "database name")
	f.StringVar(&cfg.Driver, "db.driver", "", "database driver")
}