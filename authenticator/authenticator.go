package authenticator

import (
	"crypto/rand"
	"encoding/base32"
	"encoding/base64"
	"strings"

	"code.google.com/p/rsc/qr"
	"github.com/dgryski/dgoogauth"
)

func randStr(strSize int, randType string) string {
	var dictionary string
	if randType == "alphanum" {
		dictionary = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	}

	if randType == "alpha" {
		dictionary = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	}

	if randType == "number" {
		dictionary = "0123456789"
	}

	var bytes = make([]byte, strSize)
	rand.Read(bytes)
	for k, v := range bytes {
		bytes[k] = dictionary[v%byte(len(dictionary))]
	}
	return string(bytes)
}

func GenerateNewSecretAndImage(issuer string) (string, string, error) {
	randomStr := randStr(6, "alphanum")
	secret := base32.StdEncoding.EncodeToString([]byte(randomStr))

	authLink := "otpauth://totp/" + issuer + "?secret=" + secret + "&issuer=" + issuer

	code, err := qr.Encode(authLink, qr.H)

	if err != nil {
		return "", "", err
	}

	imgByte := code.PNG()

	str := base64.StdEncoding.EncodeToString(imgByte)

	return secret, str, nil
}

func IsTokenValid(tokenSecret string, token string) bool {
	otpConfig := &dgoogauth.OTPConfig{
		Secret:      strings.TrimSpace(tokenSecret),
		WindowSize:  3,
		HotpCounter: 0,
	}

	trimmedToken := strings.TrimSpace(token)

	// Validate token
	ok, err := otpConfig.Authenticate(trimmedToken)

	if err != nil || !ok {
		return false
	}
	return true
}
