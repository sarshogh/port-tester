package main

import (
	"errors"
	"fmt"
	"os"
	"path"
)

func ConfigFileExists() (os.FileInfo, error) {
	configFileName := os.Getenv("PORT_TESTER_CONFIG_FILE_NAME")
	if configFileName == "" {
		configFileName = "template"
	}

	_, err := FileExists(path.Join(path.Base(""),configFileName))
	if err != nil {
		panic(fmt.Sprintf("config file `%s` not found!", configFileName))
	}

	return os.Stat(configFileName)
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
