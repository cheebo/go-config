package types


type RedisService struct {
	Service        string           `json:"service"`
	Tag            string           `json:"tag"`
	Addr           string           `json:"addr"`
	Port           string           `json:"port"`
	Pass           string           `json:"password"`
	Database       int              `json:"database"`
	PoolSize       int              `json:"pool_size"`
}
