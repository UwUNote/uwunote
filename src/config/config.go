package config

import (
	"os"
	"path/filepath"
)

var (
	configPath = filepath.FromSlash(os.Getenv("HOME") + string(os.PathSeparator) + ".uwunote")
)

//CreateNeccessaryFiles creates the config folder
func CreateNeccessaryFiles() {
	os.MkdirAll(configPath, os.ModePerm)
}
