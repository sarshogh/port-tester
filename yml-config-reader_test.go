package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_DefaultConfigFileExists(t *testing.T) {
	templatePath := "./template.yml"
	assert.FileExists(t, templatePath)
}

func Test_ConfigFileNameViaEnvVars(t *testing.T) {
	_, err := ConfigFileExists()

	if err != nil {
		t.Error("Loading config template file failed!")
	}
}
