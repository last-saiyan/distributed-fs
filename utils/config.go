package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

// util for parsing configs json

// Config contains configuration
type Config struct {
	BlockSize         int64
	Replica           int
	IoSize            int
	NameNodePort      string
	DataNodePort      string
	EditLog           string
	DataDir           string
	NameNodeHost      string
	HeartbeatInterval int
	HeartbeatTimeout  int
}

// GetConfig parses config.json and returns
// Config struct
func GetConfig() Config {
	var config Config
	file := "./config.json"
	configFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer configFile.Close()
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	CreateDir(config.DataDir)
	return config
}
