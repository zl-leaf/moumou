package internal

import (
	"encoding/base64"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomImage(t *testing.T) {
	svc := NewCaptchaService()
	randomId, digits, err := svc.RandomImage()
	assert.Nil(t, err)
	assert.NotEmpty(t, randomId)
	assert.NotEmpty(t, digits)

	data := base64.StdEncoding.EncodeToString(digits)
	t.Log(data)
}
