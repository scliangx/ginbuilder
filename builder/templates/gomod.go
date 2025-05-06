package templates

var GoModTemplate = `module %s

go %s

require (
	github.com/gin-gonic/gin v1.10.0
	github.com/go-redis/redis/v8 v8.11.5
	github.com/natefinch/lumberjack v2.0.0+incompatible
	go.uber.org/zap v1.24.0
	gopkg.in/yaml.v2 v2.4.0
	gorm.io/driver/mysql v1.5.0
	gorm.io/gorm v1.25.0
)

replace github.com/ugorji/go v1.1.4 => github.com/ugorji/go/codec v1.2.9
`
