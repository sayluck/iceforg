package config

type DB struct {
	Mysql *Mysql `type:"mysql",yaml:"mysql",toml:"mysql"`
	Mongo *Mongo `type:"mongo",yaml:"mongo",toml:"mongo"`
	Redis *Redis `type:"redis",yaml:"redis",toml:"redis"`
}

type Mysql struct {
	URL          string `yaml:"url",toml:"url"`
	UserName     string `yaml:"userName",toml:"userName"`
	Password     string `yaml:"password",toml:"password"`
	DBName       string `yaml:"dbName",toml:"dbName"`
	MaxIdleConns int    `yaml:"maxIdleConns",toml:"maxIdleConns"`
	MaxOpenConns int    `yaml:"maxOpenConns",toml:"maxOpenConns"`
	LogMode      bool   `yaml:"logMode",toml:"logMode"`
}
type Mongo struct {
}

type Redis struct {
	Addr     string `yaml:"addr",toml:"addr"`
	Password string `yaml:"password",toml:"password"`
	DBNum    string `yaml:"dbNum",toml:"dbNum"`
	Pong     string `yaml:"pong",toml:"pong"`
}
