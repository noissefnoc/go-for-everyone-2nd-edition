package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
)

type config struct {
	ApiKey   string `json:"api_key"`
	MaxCount int    `json:"max_count"`
}

func loadConfig() (*config, error) {
	// Get current user
	u, err := user.Current()
	if err != nil {
		return nil, err
	}

	// Refer home directory
	fname := filepath.Join(u.HomeDir, ".config", "my-app", "config.json")
	f, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var cfg config
	err = json.NewDecoder(f).Decode(&cfg)
	return &cfg, err
}

func main() {
	cfg, err := loadConfig()
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	fmt.Printf("%s, %d\n", cfg.ApiKey, cfg.MaxCount)
}
