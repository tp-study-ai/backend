package tools

import "github.com/BurntSushi/toml"

type Config struct {
	DbName   string `toml:"Dbname"`
	User     string `toml:"User"`
	Password string `toml:"Password"`
	Port     int    `toml:"Port"`
	Host     string `toml:"Host"`
	Testis   string `toml:"testis"`
	Ml       string `toml:"ml"`
}

func NewConfig() *Config {
	return &Config{}
}

func ReadConfigFile(configPath string, dst interface{}) error {
	_, err := toml.DecodeFile(configPath, dst)
	return err
}
