package utils

import (
	"crypto/md5"
	"encoding/hex"
	"net/http"
	"strings"
	"time"

	"bitbucket.org/pqstudio/go-webutils/web"

	"code.google.com/p/go-uuid/uuid"
	"github.com/PuerkitoBio/throttled"
	"github.com/PuerkitoBio/throttled/store"

	. "bitbucket.org/pqstudio/go-webutils/logger"
)

func RateLimitCustom(r *http.Request) string {
	IP := web.GetClientIP(r)

	Log.Debug("%s", IP)
	p := r.URL.Path

	return IP + p
}

func RateLimit(i int, minutes time.Duration) *throttled.Throttler {
	//pool := redis.GetPool()
	//keyPrefix := "rate-limit:"
	//db := 0
	return throttled.RateLimit(throttled.Q{i, minutes * time.Minute}, &throttled.VaryBy{Custom: RateLimitCustom}, store.NewMemStore(1000))
	//return throttled.RateLimit(throttled.Q{i, minutes * time.Minute}, &throttled.VaryBy{Custom: RateLimitCustom}, store.NewRedisStore(pool, keyPrefix, db))
}

func NewUUID() string {
	uuid := uuid.New()

	return strings.Replace(uuid, "-", "", -1)
}

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
