package config

type DB struct {
	Mysql *Mysql `type:"mysql",yaml:"mysql",toml:"mysql"`
	Mongo *Mongo `type:"mysql",yaml:"mongo",toml:"mongo"`
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
