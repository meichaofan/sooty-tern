package config

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	Log Log `toml:"log"`
}

type Log struct {
	Level      int    `toml:"level"`
	Format     string `toml:"format"`
	OutputDir  string `toml:"output_dir"`
	OutputFile string `toml:"output_file"`
}
