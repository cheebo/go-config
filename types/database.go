package types

type DatabaseConfig struct {
	Host     string `json:"host"`
	Port     uint   `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
	Driver   string `json:"driver"`
}
