package utils

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"net/http"
	"strings"

	"github.com/gregory90/go-webutils/web"

	"code.google.com/p/go-uuid/uuid"
)

func RateLimitCustom(r *http.Request) string {
	IP := web.GetClientIP(r)

	p := r.URL.Path

	return IP + p
}

func NewUUID() string {
	uuid := uuid.New()

	return strings.Replace(uuid, "-", "", -1)
}

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func Base64Encode(text string) string {
	return base64.StdEncoding.EncodeToString([]byte(data))
}
