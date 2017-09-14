package types

type DatabaseService struct {
	Service        string  `json:"service"`
	Tag            string  `json:"tag"`
	DatabaseConfig
}

type DatabaseConfig struct {
	Host           string  `json:"host"`
	Port           uint    `json:"port"`
	User           string  `json:"user"`
	Password       string  `json:"password"`
	Database       string  `json:"database"`
	Driver         string  `json:"driver"`
}