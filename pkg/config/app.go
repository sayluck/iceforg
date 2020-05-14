package config

type App struct {
	Port  string `yaml:"port",toml:"port"`
	Mode  string `yaml:"mode",toml:"mode"`
	Log   *Log   `yaml:"log",toml:"log"`
	Token *Token `yaml:"token",toml:"token"`
}

type Log struct {
	Level       string `yaml:"level",toml:"level"`
	PrettyPrint bool   `yaml:"prettyPrint",toml:"prettyPrint"`
}

type Token struct {
	ExpiredAfter int64 `yaml:"expiredAfter",toml:"expiredAfter"`
}
