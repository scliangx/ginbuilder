package templates

var ApiTemplate = `package apis

import (
	"{{.PLACEHOLDER}}/apps/apis/hello"
)

type ApiGroup struct{
	hello.HelloApi
}

var ApiGroupApi = new(ApiGroup)
`

var HelloApiTemplate = `package hello


import (
	"{{.PLACEHOLDER}}/common/responsex"
	"{{.PLACEHOLDER}}/models"
	"{{.PLACEHOLDER}}/common/errorx"
	"{{.PLACEHOLDER}}/apps/service"
	"github.com/gin-gonic/gin"
)

type HelloApi struct{}

func (h *HelloApi) HelloWorld(c *gin.Context) {
	req := models.HelloRequest{}
	err := c.ShouldBind(&req)
	if err != nil {
		responsex.FailWithCode(errorx.ArgumentError,"",c)
		return
	}
	helloResponse, code := service.HelloService(req)
	if code != errorx.SuccessCode {
		responsex.FailWithCode(code, "", c)
		return
	}

	responsex.OkWithData(helloResponse, c)
	return
}
`
