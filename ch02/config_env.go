package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
)

type target struct {
	Name      string `json:"name"`
	Threshold int    `json:"threshold"`
}

type config2 struct {
	Addr   string   `json:"addr"`
	Target []target `json:"target"`
}

func loadConfig2() (*config2, error) {
	// refer home directory
	var configDir string
	home := os.Getenv("HOME")
	if home == "" && runtime.GOOS == "windows" {
		// on windows environment
		configDir = os.Getenv("APPDATA")
	} else {
		configDir = filepath.Join(home, ".config")
	}
	fname := filepath.Join(configDir, "my-app", "config2.json")
	f, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var cfg config2
	err = json.NewDecoder(f).Decode(&cfg)
	return &cfg, err
}
