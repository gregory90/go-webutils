package request

import (
	"github.com/parnurzeal/gorequest"
)

var Client *gorequest.SuperAgent

func Init() error {
	request := gorequest.New()
	Client = request
	return nil
}
