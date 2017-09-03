package types


type RedisService struct {
	Service        string           `json:"service"`
	Tag            string           `json:"tag"`
	Addr           string           `json:"addr"`
	Port           string           `json:"port"`
	Password       string           `json:"password"`
	Database       int              `json:"database"`
	PoolSize       int              `json:"pool_size"`
	MasterName     string           `json:"master_name"`
	SlaveReadOnly  bool             `json:"slave_read_only"`
}
