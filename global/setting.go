package global
import (
	"github.com/gin-blog-service/pkg/setting"
	"github.com/gin-blog-service/pkg/logger"
)
var (
	ServerSetting	*setting.ServerSettings
	AppSetting		*setting.AppSettings
	DatabaseSetting	*setting.DatabaseSettings
	Logger			*logger.Logger
)