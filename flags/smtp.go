package flags

import (
	"flag"
	"github.com/cheebo/go-config/types"
)

func SmtpFlags(cfg *types.SMTPConfig, f *flag.FlagSet) {
	f.StringVar(&cfg.Host, "smtp.host", "", "smtp host: 'localhost' or '127.0.0.1'")
	f.UintVar(&cfg.Port, "smtp.port", 0, "port number")
	f.StringVar(&cfg.User, "smtp.user", "", "user's name")
	f.StringVar(&cfg.Password, "smtp.password", "", "user's password")
	f.BoolVar(&cfg.SSL, "smtp.ssl", false, "use ssl")
	f.StringVar(&cfg.LocalName, "smtp.localname", "", "the name to use in HELO/EHLO")
}