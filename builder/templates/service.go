package templates

var HelloServiceTemplate = `package service

import (
	"{{.pkgname}}/common/errorx"
	"{{.pkgname}}/models"
)

type HelloResponse struct {
	Msg string #json:"msg"#
}

func HelloService(req models.HelloRequest) (HelloResponse, errorx.ErrorCode) {
	return HelloResponse{Msg: "hello service"}, errorx.SuccessCode
}

`
