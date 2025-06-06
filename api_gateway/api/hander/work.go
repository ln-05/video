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

func PublishContent(c *gin.Context) {
	var req request.PublishContents
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

	content, _ := client.PublishContent(c, &__.PublishContentRequest{
		Title:     req.Title,
		Desc:      req.Desc,
		MusicId:   req.MusicId,
		WorkType:  req.WorkType,
		IpAddress: req.IpAddress,
	})

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": content,
	})
}

func ListWork(c *gin.Context) {
	var req request.ListWork
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

	content, _ := client.ListWork(c, &__.ListWorkRequest{
		Page: req.Page,
		Size: req.Size,
	})

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": content,
	})
}

func InfoWork(c *gin.Context) {
	var req request.InfoWork
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

	content, _ := client.InfoWork(c, &__.InfoWorkRequest{
		Id: req.Id,
	})

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": content,
	})
}
