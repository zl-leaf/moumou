package internal

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"strconv"
)

// EncryptPassword 加密
func EncryptPassword(passwordPlainText string, userCreatedAt int64) string {
	createdAtStr := strconv.FormatInt(userCreatedAt, 10)
	h := hmac.New(sha256.New, []byte(createdAtStr))

	h.Write([]byte(passwordPlainText))
	encryptPwd := hex.EncodeToString(h.Sum(nil))

	return encryptPwd
}
