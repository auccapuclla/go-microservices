package security

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncryptPassword(t *testing.T) {
	password, err := EncryptPassword("123411156")
	assert.NoError(t, err)
	assert.NotEmpty(t, password)
	assert.Len(t, password, 60)

}

func TestVerifyPassword(t *testing.T) {
	password, err := EncryptPassword("123456")
	assert.NoError(t, err)
	assert.NotEmpty(t, password)

	assert.NoError(t, VerifyPassword("123456", password))

	assert.Error(t, VerifyPassword("1234567", password))
	assert.Error(t, VerifyPassword(password, password))
	assert.Error(t, VerifyPassword(password, "1234567"))
}
