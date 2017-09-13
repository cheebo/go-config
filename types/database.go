package types

type DatabaseService struct {
	Service        string  `json:"service"`
	Tag            string  `json:"tag"`
	DatabaseConfig
}

type DatabaseConfig struct {
	Addr           string  `json:"addr"`
	Port           string  `json:"port"`
	User           string  `json:"user"`
	Pass           string  `json:"password"`
	Database       string  `json:"database"`
	Driver         string  `json:"driver"`
}