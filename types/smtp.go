package types

type SMTPConfig struct {
	Host            string          `json:"host"`
	Port            int             `json:"port"`
	Username        string          `json:"username"`
	Password        string          `json:"password"`
	SSL             bool            `json:"ssl"`
	LocalName       string          `json:"local_name"`
}
