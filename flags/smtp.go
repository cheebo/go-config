package flags

import (
	"flag"
	"github.com/cheebo/go-config/types"
)

func SmtpFlags(cfg *types.SMTPConfig, f *flag.FlagSet) {
	f.StringVar(&cfg.Host, "smtp.host", cfg.Host, "smtp host: 'localhost' or '127.0.0.1'")
	f.UintVar(&cfg.Port, "smtp.port", cfg.Port, "port number")
	f.StringVar(&cfg.User, "smtp.user", cfg.User, "user's name")
	f.StringVar(&cfg.Password, "smtp.password", cfg.Password, "user's password")
	f.BoolVar(&cfg.SSL, "smtp.ssl", cfg.SSL, "use ssl")
	f.StringVar(&cfg.LocalName, "smtp.localname", cfg.LocalName, "the name to use in HELO/EHLO")
}