package hander

import (
	"api_gateway/api/request"
	"api_gateway/pkg"
	__ "api_gateway/proto"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
)

func Sendsms(c *gin.Context) {
	var req request.SendSms
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

	client.SendSms(c, &__.SendSmsRequest{
		Mobile:      req.Mobile,
		SendSmsCode: req.SendSmsCode,
	})

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "发送验证码成功",
	})
}

func Login(c *gin.Context) {
	var req request.Login
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
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "连接服务失败",
			"data": nil,
		})
		return
	}
	defer conn.Close()

	client := __.NewUserClient(conn)

	re, err := client.Login(c, &__.LoginRequest{
		Mobile:      req.Mobile,
		SendSmsCode: req.SendSmsCode,
	})

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}

	// 检查返回值是否为nil
	if re == nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "登录服务返回空数据",
			"data": nil,
		})
		return
	}

	// 生成JWT令牌
	token, err := pkg.NewJWT("2211a").CreateToken(pkg.CustomClaims{
		ID: uint(re.Id),
	})

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "生成令牌失败",
			"data": nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "登录成功",
		"data": token,
	})
}

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
		UserId:      req.UserId,
		Title:       req.Title,
		Desc:        req.Desc,
		MusicId:     req.MusicId,
		WorkType:    req.WorkType,
		CheckStatus: req.CheckStatus,
		CheckUser:   req.CheckUser,
		IpAddress:   req.IpAddress,
	})

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "发布成功",
		"data": content,
	})
}

func UpdateStatus(c *gin.Context) {
	var req request.UpdateStatus
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

	client.UpdateStatus(c, &__.UpdateStatusRequest{
		Id:          req.Id,
		CheckStatus: req.CheckStatus,
	})

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "审核成功",
	})
}
