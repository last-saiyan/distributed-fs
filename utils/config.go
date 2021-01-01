package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

// util for parsing configs json

// Config contains configuration
type Config struct {
	BlockSize    int
	NameNodePort int
	DataNodePort int
	Replica      int
	EditLog      string
	DataDir      string
	NameNodeHost string
}

// GetConfig parses config.json and returns
// Config struct
func GetConfig() Config {
	var config Config
	file := "./config.json"
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	err = CreateDir(config.DataDir)
	if err != nil {
		fmt.Println(err.Error())
	}
	return config
}
