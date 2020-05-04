package cafa

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

//Special access without the session key (HMAC)
func NewHMAC(clientCode, apiRoute, secret, timestamp string) (string, error) {
	var (
		//get mac using internal secret
		mac = hmac.New(sha256.New, []byte(secret))
		//compose a message to sign
		msg = clientCode + apiRoute + timestamp
	)
	if _, err := mac.Write([]byte(msg)); err != nil {
		return "", err
	}
	//get the hmac-auth content
	return hex.EncodeToString(mac.Sum(nil)), nil
}
