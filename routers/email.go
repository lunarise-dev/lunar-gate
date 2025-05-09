package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/lunarise-dev/lunar-gate/api"
)

func EmailRouter(g *gin.RouterGroup) {
	r := g.Group("email")
	app := api.App
	r.POST("send", app.SendEmail)
}
