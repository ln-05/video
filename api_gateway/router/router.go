package router

import (
	"api_gateway/api/hander"
	"api_gateway/pkg"
	"github.com/gin-gonic/gin"
)

func LoadRouter(r *gin.Engine) {
	g := r.Group("/videouser")
	{
		user := g.Group("/user")
		{
			user.POST("/sendsms", hander.Sendsms)
			user.POST("/login", hander.Login)
			user.Use(pkg.JWTAuth("2211a"))
			user.POST("/personal", hander.Personal)
			user.POST("/updatePersonal", hander.UpdatePersonal)
		}

		work := g.Group("/work")
		{
			work.POST("/publishContent", hander.PublishContent)
			work.POST("/listWork", hander.ListWork)
			work.POST("/infoWork", hander.InfoWork)
		}

		comment := g.Group("/comment")
		{
			comment.POST("/postComment", hander.PostComment)
		}

	}

}
