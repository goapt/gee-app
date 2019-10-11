package util

import (
	"encoding/base64"

	"github.com/goapt/golib/cryptoutil"
)

func AesEncrypt(token, data string) (string, error) {
	s, err := cryptoutil.AesEncrypt([]byte(token), []byte(data))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(s), nil
}

func AesDecrypt(token, data string) (string, error) {
	bt, err := base64.StdEncoding.DecodeString(data)

	if err != nil {
		return "", err
	}

	s, err := cryptoutil.AesDecrypt([]byte(token), bt)

	if err != nil {
		return "", err
	}

	return string(s), nil
}
