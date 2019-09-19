package config

import (
	"github.com/BurntSushi/toml"
)

var (
	global *Config
)

/**
ParseConfig解析toml配置文件
*/
func ParseConfig(path string) (*Config, error) {
	var c Config
}

type Config struct {
	Log Log `toml:"log"`
}

type Log struct {
	Level      int    `toml:"level"`
	Format     string `toml:"format"`
	OutputDir  string `toml:"output_dir"`
	OutputFile string `toml:"output_file"`
}
