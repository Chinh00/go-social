package redis

type RedisOption struct {
	Host string `json:"host" yaml:"host"`
	Port int    `json:"port" yaml:"port"`
}
