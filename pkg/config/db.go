package config

type db struct {
	Mysql *Mysql `yaml:"mysql",toml:"mysql"`
	Mongo *mongo `yaml:"mongo",toml:"mongo"`
}

type Mysql struct {
	URL      string `yaml:"url",toml:"url"`
	User     string `yaml:"user",toml:"user"`
	Password string `yaml:"password",toml:"password"`
}
type mongo struct {
}
