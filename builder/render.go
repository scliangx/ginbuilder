package builder

import (
	"fmt"
	"github.com/scliangx/ginbuilder/builder/templates"
	"github.com/scliangx/ginbuilder/builder/tools"
	"os"
	"path"
)

func RenderFile(projectName, packageName, projectPath string) (string, error) {

	basepath := path.Join(projectPath, projectName)
	fmt.Printf("[INFO]: Project Name        : %v\n", projectName)
	fmt.Printf("[INFO]: Project Package Name: %v\n", packageName)
	fmt.Printf("[INFO]: Project Path        : %v\n", basepath)
	_, err := os.ReadDir(basepath)
	if err == nil || !os.IsNotExist(err) {
		fmt.Printf("[ERROR]: [%v] project already exists", basepath)
		return "", fmt.Errorf("[%v] project already exists", basepath)
	}
	tools.IsExistsDirectoryAndCreate(basepath)
	// go.mod
	tools.WriteToFile(path.Join(basepath, "go.mod"), fmt.Sprintf(templates.GoModTemplate, packageName, tools.GetGoVersion()))
	// main.go
	tools.WriteToFile(path.Join(basepath, "main.go"), tools.ReplaceAppNameTemplate(templates.MainTemplate, packageName))

	fs := path.Join(basepath, "models")
	tools.IsExistsDirectoryAndCreate(fs)
	tools.WriteToFile(path.Join(fs, "hello.go"), tools.ReplaceAppNameTemplate(templates.HelloModes, packageName))

	// config
	fs = path.Join(basepath, "config")
	tools.IsExistsDirectoryAndCreate(fs)
	tools.WriteToFile(path.Join(fs, "config.go"), tools.ReplaceAppNameTemplate(templates.ConfigBaseTemplate, packageName))
	tools.WriteToFile(path.Join(fs, "config.yaml"), tools.ReplaceAppNameTemplate(templates.ConfigYamlTemplate, packageName))

	fs = path.Join(fs, "internal_config")
	tools.IsExistsDirectoryAndCreate(fs)
	tools.WriteToFile(path.Join(fs, "mysql.go"), tools.ReplaceAppNameTemplate(templates.ConfigMysqlTemplate, packageName))
	tools.WriteToFile(path.Join(fs, "redis.go"), tools.ReplaceAppNameTemplate(templates.ConfigRedisTemplate, packageName))
	tools.WriteToFile(path.Join(fs, "system.go"), tools.ReplaceAppNameTemplate(templates.ConfigSysTempalte, packageName))
	tools.WriteToFile(path.Join(fs, "logger.go"), tools.ReplaceAppNameTemplate(templates.ConfigLoggerTemplate, packageName))

	// errorx
	fs = path.Join(basepath, "common/errorx")
	tools.IsExistsDirectoryAndCreate(fs)
	tools.WriteToFile(path.Join(fs, "errorx.go"), tools.ReplaceAppNameTemplate(templates.ErrorxTempalte, packageName))

	fs = path.Join(basepath, "common/logx")
	tools.IsExistsDirectoryAndCreate(fs)
	tools.WriteToFile(path.Join(fs, "logx.go"), tools.ReplaceAppNameTemplate(templates.LoggerxTemplate, packageName))

	fs = path.Join(basepath, "common/responsex")
	tools.IsExistsDirectoryAndCreate(fs)
	tools.WriteToFile(path.Join(fs, "responsex.go"), tools.ReplaceAppNameTemplate(templates.ResponsexTemplate, packageName))

	fs = path.Join(basepath, "global")
	tools.IsExistsDirectoryAndCreate(fs)
	tools.WriteToFile(path.Join(fs, "global.go"), tools.ReplaceAppNameTemplate(templates.GlobalFileTempalte, packageName))

	// internal
	fs = path.Join(basepath, "internal")
	tools.IsExistsDirectoryAndCreate(fs)
	tools.WriteToFile(path.Join(fs, "mysql.go"), tools.ReplaceAppNameTemplate(templates.InitMysqlTemplate, packageName))
	tools.WriteToFile(path.Join(fs, "redis.go"), tools.ReplaceAppNameTemplate(templates.InitRedisTemplate, packageName))

	// router
	fs = path.Join(basepath, "apps/routers")
	tools.IsExistsDirectoryAndCreate(fs)
	tools.WriteToFile(path.Join(fs, "init_router.go"), tools.ReplaceAppNameTemplate(templates.InitRoutersTemplate, packageName))
	tools.WriteToFile(path.Join(fs, "hello_router.go"), tools.ReplaceAppNameTemplate(templates.HelloRouterTemplate, packageName))

	// api
	fs = path.Join(basepath, "apps/apis")
	tools.IsExistsDirectoryAndCreate(fs)
	tools.WriteToFile(path.Join(fs, "api.go"), tools.ReplaceAppNameTemplate(templates.ApiTemplate, packageName))
	fs = path.Join(fs, "/hello")
	tools.IsExistsDirectoryAndCreate(fs)
	tools.WriteToFile(path.Join(fs, "hello.go"), tools.ReplaceAppNameTemplate(templates.HelloApiTemplate, packageName))

	// service
	fs = path.Join(basepath, "apps/service")
	tools.IsExistsDirectoryAndCreate(fs)
	tools.WriteToFile(path.Join(fs, "hello.go"), tools.ReplaceAppNameTemplate(templates.HelloServiceTemplate, packageName))

	fmt.Printf("[INFO]: %v project create success \n", basepath)
	return basepath, nil

}
