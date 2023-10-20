package test

import (
	"belajar-golang-restful-api/simple"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimpleServiceError(t *testing.T) {
	simpleservice, err := simple.InitializedService(true)
	assert.NotNil(t, err)
	assert.Nil(t, simpleservice)
}

func TestSimpleServiceSuccess(t *testing.T) {
	simpleService, err := simple.InitializedService(false)
	assert.Nil(t, err)
	assert.NotNil(t, simpleService)
}
