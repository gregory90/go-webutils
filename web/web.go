package web

import (
	"encoding/json"
	//"errors"
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/gorilla/context"

	"bitbucket.org/pqstudio/go-webutils/slice"

	. "bitbucket.org/pqstudio/go-webutils/logger"
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

	if limit > 100 {
		limit = 100
	}

	return limit, offset
}

func Fields(c *gin.Context) []string {
	fields := strings.Split(c.Query("fields"), ",")

	return fields
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
