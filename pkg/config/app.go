package config

type App struct {
	Port string `yaml:"port",toml:"port"`
	Mode string `yaml:"mode",toml:"mode"`
	Log  *Log   `yaml:"log",toml:"log"`
}

type Log struct {
	Level       string `yaml:"level",toml:"level"`
	PrettyPrint bool   `yaml:"prettyPrint",toml:"prettyPrint"`
}
