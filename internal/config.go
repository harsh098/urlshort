package internal

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

var configData *Config = nil

func getConfig() ([]byte, error) {
	var configFilePath string
	path, ok := os.LookupEnv("REDIRECT_CONFIG_DIR")

	if !ok {
		altpath, _ := os.Getwd()
		path = altpath
	}

	configFilePath = filepath.Join(path, "config.yml")
	if _, err := os.Stat(configFilePath); errors.Is(err, os.ErrNotExist) {
		configFilePath = filepath.Join(path, "config.yaml")
	}

	file, err := os.Open(configFilePath)

	defer func() {
		if err := file.Close(); err != nil {
			log.Fatalf("Cannot Close Stream\n")
		}
	}()

	if err != nil {
		log.Fatalf("Cannot Open File in Specified Location: %v\n", err.Error())
		return nil, err
	}

	var data []byte
	var buffer *bytes.Buffer = new(bytes.Buffer)
	_, err = buffer.ReadFrom(file)

	if err != nil {
		log.Fatalf("Could Not Read File: %v\n", err.Error())
		return nil, err
	}
	data = buffer.Bytes()
	return data, err
}

func getConfigStruct() (*Config, error) {
	var config Config
	data, err := getConfig()

	if err != nil {
		log.Fatalf("Could not Read YAML Data: %v\n", err.Error())
		return nil, err
	}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("Invalid YAML Format: %v\n", err.Error())
		return nil, err
	}
	return &config, err
}

func GetConfig() (*Config, error) {
	var err error
	if configData == nil {
		configData, err = getConfigStruct()
	}

	return configData, err
}

func GetHost() (string, error) {
	var configStruct *Config
	var err error
	configStruct, err = GetConfig()
	return configStruct.Host, err

}

func GetSocketAddress() (string, error) {
	var configStruct *Config
	var err error
	configStruct, err = GetConfig()
	var socketAddress string = fmt.Sprintf("%v:%v", configStruct.Host, configStruct.Port)
	return socketAddress, err
}
