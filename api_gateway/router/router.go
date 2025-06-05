package router

import (
	"api_gateway/api/hander"
	"api_gateway/pkg"
	"github.com/gin-gonic/gin"
)

func LoadRouter(r *gin.Engine) {
	g := r.Group("/videouser")
	{
		g.POST("/sendsms", hander.Sendsms)
		g.POST("/login", hander.Login)
		g.Use(pkg.JWTAuth("2211a"))
		g.POST("/publishContent", hander.PublishContent)
		//g.POST("/updateStatus", hander.UpdateStatus)
		//g.POST("/realname", hander.Realname)
		g.POST("/personal", hander.Personal)
		g.POST("/updatePersonal", hander.UpdatePersonal)
		g.POST("/listWork", hander.ListWork)
		g.POST("/infoWork", hander.InfoWork)
	}

}
