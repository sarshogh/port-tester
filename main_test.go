package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_main(t *testing.T) {
	result := true
	t.Log("Go Test-context is working!")
	assert.Equal(t, true, result)
}
