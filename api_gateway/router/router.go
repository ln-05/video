package router

import (
	"api_gateway/api/hander"
	"github.com/gin-gonic/gin"
)

func LoadRouter(r *gin.Engine) {
	g := r.Group("/videouser")
	{
		g.POST("/sendsms", hander.Sendsms)
		g.POST("/login", hander.Login)
	}

}
