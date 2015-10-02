package utils

import (
	"database/sql"
	"net/http"

	"bitbucket.org/pqstudio/go-webutils/web"

	. "bitbucket.org/pqstudio/go-webutils/logger"
)

type handleErr func(http.ResponseWriter, *http.Request) error

// handle all errors from application here
func (fn handleErr) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := fn(w, r); err != nil {

		switch err {
		case sql.ErrNoRows:
			Log.Debug(err.Error())
			HttpError(w, nil, 404)
			return
		}

		switch e := err.(type) {
		case *web.ValidationError:
			err := err.(*web.ValidationError)
			err.ErrorType = "validation_error"
			Log.Debug(e.Error())
			HttpError(w, err, 422)
			return
		case *web.SerializationError:
			Log.Debug(e.Error())
			HttpError(w, err, 400)
			return
		case *web.NotFound:
			Log.Debug(e.Error())
			HttpError(w, err, 404)
			return
		case *web.Forbidden:
			Log.Debug(e.Error())
			HttpError(w, err, 403)
			return
		case *web.NoContent:
			Log.Debug(e.Error())
			HttpError(w, err, 204)
			return
		case *web.BadRequest:
			Log.Debug(e.Error())
			HttpError(w, err, 400)
			return
		}

		Log.Critical(err.Error())
		HttpError(w, err, 500)

		return
	}
}

func HttpError(w http.ResponseWriter, err error, status int) {
	if err != nil {
		w.WriteHeader(status)
	} else {
		w.WriteHeader(status)
	}
}
