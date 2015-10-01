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

func LimitAndOffset(r *http.Request) (int, int) {
	limit, _ := strconv.Atoi(r.FormValue("limit"))
	offset, _ := strconv.Atoi(r.FormValue("offset"))

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

func Bind(r io.ReadCloser, obj Model) error {
	err := FromJSON(r, obj)
	if err != nil {
		return err
	}

	if errors := obj.Validate(); errors != nil {
		return &ValidationError{
			Errors: errors,
		}
	}
	return nil
}

func FromJSON(r io.ReadCloser, obj Model) error {
	defer r.Close()
	if err := json.NewDecoder(r).Decode(obj); err != nil {
		return &SerializationError{Message: "deserialization_error"}
	}
	return nil
}

func FromJSONStrict(r io.ReadCloser, obj interface{}) error {
	defer r.Close()
	if err := json.NewDecoder(r).Decode(obj); err != nil {
		Log.Debug("%+v", err)
		return &SerializationError{Message: "deserialization_error"}
	}
	return nil
}

func FromJSONString(data string, obj interface{}) error {
	byt := []byte(data)
	if err := json.Unmarshal(byt, obj); err != nil {
		Log.Debug("%+v", err)
		return &SerializationError{Message: "deserialization_error"}
	}
	return nil
}

func ToJSON(w http.ResponseWriter, obj interface{}) error {
	if err := json.NewEncoder(w).Encode(obj); err != nil {
		return &SerializationError{Message: "serialization_error"}
	}
	return nil
}

func ToWhitelistedJSON(w http.ResponseWriter, obj interface{}, fields []string) error {
	if len(fields) > 0 && fields[0] != "" {
		res, err := Whitelist(obj, fields)
		if err != nil {
			return err
		}

		ToJSON(w, &res)
		return nil
	}
	ToJSON(w, &obj)
	return nil
}

func Whitelist(from interface{}, fields []string) (map[string]interface{}, error) {
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
	return out, nil
}

func WhitelistArray(objs []interface{}, fields []string) ([]interface{}, error) {
	if len(objs) > 0 {
		if len(fields) > 0 && fields[0] != "" {
			var out []interface{}

			for _, value := range objs {
				r, err := Whitelist(value, fields)
				if err != nil {
					return nil, err
				}

				out = append(out, r)
			}

			return out, nil
		}
	}
	return objs, nil
}

func ToWhitelistedArrayJSON(w http.ResponseWriter, objs []interface{}, fields []string) error {
	if len(objs) > 0 {
		if len(fields) > 0 && fields[0] != "" {
			var out []interface{}

			for _, value := range objs {
				r, err := Whitelist(value, fields)
				if err != nil {
					return err
				}

				out = append(out, r)
			}

			ToJSON(w, &out)
			return nil
		}
	}
	ToJSON(w, &objs)
	return nil
}

type nopCloser struct {
	io.Reader
}

func (nopCloser) Close() error { return nil }
func StringToReadCloser(data string) io.ReadCloser {

	return nopCloser{bytes.NewBufferString(data)}
}

func ContextS(r *http.Request, key string) string {
	return string(context.Get(r, key).(string))
}

func ContextI(r *http.Request, key string) int {
	i, err := strconv.Atoi(string(context.Get(r, key).(string)))
	Log.Error(err.Error())
	return i
}

func A(m map[string]interface{}, key string, message string) bool {
	m[key] = make([]string, 0)
	m[key] = append(m[key].([]string), message)

	return true
}
