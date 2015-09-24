package logger

import (
	"os"

	"github.com/op/go-logging"
)

var Log = logging.MustGetLogger("default")

func init() {
	logBackend := logging.NewLogBackend(os.Stderr, "", 0)
	logging.SetBackend(logBackend)

	format := "%{color}%{time:15:04:05.000000} %{level:.3s} %{id:03x}%{color:reset} %{message}"
	logging.SetFormatter(logging.MustStringFormatter(format))
}

func SetLevel(level logging.Level) {
	logging.SetLevel(level, "default")
}
