package main

import (
	"errors"
	"fmt"
	"os"
)

func ConfigFileExists() (os.FileInfo, error) {
	config := os.Getenv("PORT_TESTER_CONFIG_FILE_NAME")
	if config == "" {
		config = "template"
	}

	_, err := FileExists(config)
	if err != nil {
		panic(fmt.Sprintf("config file %s not found!", config))
	}

	return os.Stat(config)
}

func FileExists(filePath string) (bool, error) {
	info, err := os.Stat(filePath)
	if err == nil {
		return !info.IsDir(), nil
	}
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	return false, err
}
