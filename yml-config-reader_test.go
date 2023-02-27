package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_access_to_template_file(t *testing.T) {
	templatePath := "./template.yml"
	assert.FileExists(t, templatePath)
}
