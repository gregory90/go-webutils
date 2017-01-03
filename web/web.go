package web

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"

	"gopkg.in/gin-gonic/gin.v1"

	"github.com/gregory90/go-webutils/slice"

	. "github.com/gregory90/go-webutils/logger"
)

type Model interface {
	Validate() map[string]interface{}
}

func CreateFile(file multipart.File, path string, name string) error {
	out, err := os.Create(path + name)
	if err != nil {
		return err
	}

	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		return err
	}

	return nil
}

func GetClientIP(r *http.Request) string {
	// if cloudflare is used CF-Connecting-IP is real client IP
	if ip := r.Header.Get("CF-Connecting-IP"); ip != "" {
		return ip
	}

	if ip := r.Header.Get("X-Forwarded-For"); ip != "" {
		return ip
	}

	return strings.Split(r.RemoteAddr, ":")[0]
}

func LimitAndOffset(c *gin.Context) (int, int) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	if limit == 0 {
		limit = 10
	}

	if limit > 500 {
		limit = 500
	}

	return limit, offset
}

func SortFieldAndDirection(c *gin.Context, fields []string) (string, string) {
	sortDir := c.DefaultQuery("sortDir", "DESC")
	sortField := c.DefaultQuery("sortField", "createdAt")

	if !slice.StringInSlice(sortField, fields) {
		sortField = "createdAt"
	}

	if !slice.StringInSlice(sortDir, []string{"desc", "asc", "DESC", "ASC"}) {
		sortDir = "DESC"
	}

	return sortDir, sortField
}

func Locale(c *gin.Context, locales []string, defaultLocale string) string {
	locale := c.DefaultQuery("locale", defaultLocale)

	if !slice.StringInSlice(locale, locales) {
		locale = defaultLocale
	}

	return locale
}

func Fields(c *gin.Context) []string {
	fields := strings.Split(c.Query("fields"), ",")

	return fields
}

func FromJSONString(data string, obj interface{}) error {
	byt := []byte(data)
	if err := json.Unmarshal(byt, obj); err != nil {
		Log.Debug("%+v", err)
		return &SerializationError{ErrorType: "deserialization_error"}
	}
	return nil
}

func ToJSONString(obj interface{}) (string, error) {
	b, err := json.Marshal(obj)
	return string(b), err
}

func Whitelist(from interface{}, fields []string) map[string]interface{} {
	out := make(map[string]interface{})

	var v reflect.Value
	var t reflect.Type
	v = reflect.ValueOf(from)
	t = reflect.TypeOf(from)

	if reflect.Ptr == t.Kind() {
		v = v.Elem()
		t = t.Elem()
	}

	for i := 0; i < v.NumField(); i++ {
		valueOfField := v.Field(i)
		typeOfField := t.Field(i)

		val := valueOfField.Interface()
		name := strings.Split(typeOfField.Tag.Get("json"), ",")[0]

		if len(fields) > 0 && fields[0] != "" {
			if slice.StringInSlice(name, fields) {
				out[name] = val
			}
		} else {
			out[name] = val
		}
	}
	return out
}

func WhitelistArray(objs []interface{}, fields []string) []interface{} {
	if len(objs) > 0 {
		if len(fields) > 0 && fields[0] != "" {
			var out []interface{}

			for _, value := range objs {
				r := Whitelist(value, fields)

				out = append(out, r)
			}

			return out
		}
	}
	return objs
}

type nopCloser struct {
	io.Reader
}

func (nopCloser) Close() error { return nil }
func StringToReadCloser(data string) io.ReadCloser {

	return nopCloser{bytes.NewBufferString(data)}
}

func A(m map[string]interface{}, key string, message string) bool {
	m[key] = make([]string, 0)
	m[key] = append(m[key].([]string), message)

	return true
}
