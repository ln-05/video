package hander

import (
	"api_gateway/api/request"
	__ "api_gateway/proto"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
)

func PostComment(c *gin.Context) {
	var req request.PostComment
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "参数验证失败",
			"data": nil,
		})
		return
	}

	conn, err := grpc.NewClient("127.0.0.1:8300", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := __.NewUserClient(conn)

	content, _ := client.PostComment(c, &__.PostCommentRequest{
		WorkId:  req.WorkId,
		UserId:  req.UserId,
		Content: req.Content,
		Tag:     int32(req.Tag),
		Pid:     req.Pid,
	})

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": content,
	})
}
