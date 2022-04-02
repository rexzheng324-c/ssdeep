package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFuzzyBytes(t *testing.T) {
	// normal file
	file := "test"
	ssdeep, err := FuzzyBytes([]byte(file))
	assert.NoError(t, err)
	assert.Equal(t, "3:Hn:Hn", ssdeep)

	// empty file
	file = ""
	ssdeep, err = FuzzyBytes([]byte(file))
	assert.NoError(t, err)
	assert.Equal(t, "3::", ssdeep)
}
