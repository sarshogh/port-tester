package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_DefaultConfigFileExists(t *testing.T) {
	templatePath := "./template.yml"
	assert.FileExists(t, templatePath)
}

func Test_ConfigFileNameWithDefaultValue(t *testing.T) {
	_, err := ConfigFileExists()

	if err != nil {
		t.Error("Loading config template file failed!")
	}
}

func Test_ConfigFileNameViaEnvironmentVar(t *testing.T) {
	dummy_template := "dummy_template.yml"
	os.Setenv("PORT_TESTER_CONFIG_FILE_NAME", dummy_template)

	_, err := ConfigFileExists()
	fmt.Println(err)
	if err == nil {
		t.Error("Invalid config file path founded")
	}
	expectedError := fmt.Sprintf("config file `%s` not found", dummy_template)
	if(err.Error() !=expectedError){
		t.Error("Invalid config file path founded")
	}
}
