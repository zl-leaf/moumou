package internal

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func Test_EncryptPassword(t *testing.T) {
	encryptPwd := EncryptPassword("demo", 1721317838000)
	t.Logf("got pwd:%s", encryptPwd)
	assert.Equal(t, "ca397c8980ab6baf1661c1086411d14918be6da76c9ad5e9135c7375c5053fa2", encryptPwd)
}
