package routes

import (
	"fmt"
	"gin_demo/logger"
	"gin_demo/settings"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Setup() (r *gin.Engine) {
	r = gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.GET("/", func(c *gin.Context) {
		s := fmt.Sprintf("port:%d, name:%s, version:%s, JavaHome:%s", settings.Conf.Port, settings.Conf.Name, settings.Conf.Version, settings.Conf.JavaHome)

		c.String(http.StatusOK, s)
	})
	return
}
