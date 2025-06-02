package hander

import (
	"LiveStreaming_srv/basic/global"
	"LiveStreaming_srv/model"
	__ "LiveStreaming_srv/proto"
	"context"
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"time"
)

type UserServer struct {
	__.UnimplementedUserServer
}

func isValidMobile(mobile string) bool {
	pattern := `^1[3-9]\d{9}$`
	re := regexp.MustCompile(pattern)
	return re.MatchString(mobile)
}

func (c *UserServer) SendSms(_ context.Context, in *__.SendSmsRequest) (*__.SendSmsResponse, error) {
	if !isValidMobile(in.Mobile) {
		return nil, fmt.Errorf("手机号格式不正确")
	}

	var userCount model.VideoUser
	global.DB.Where("mobile = ?", in.Mobile).Find(&userCount)

	key := "SendSms" + in.Mobile
	existingCode := global.Rdb.Get(context.Background(), key)
	if existingCode.Err() == nil {
		return nil, fmt.Errorf("验证码已发送，请稍后再试")
	}

	code := rand.Intn(9000) + 1000

	sendCountKey := "SendSmsCount_" + in.Mobile
	count, err := global.Rdb.Incr(context.Background(), sendCountKey).Result()
	if err == nil && count > 5 {
		return nil, fmt.Errorf("验证码发送过于频繁，请明天再试")
	}
	global.Rdb.Expire(context.Background(), sendCountKey, 24*time.Hour)

	global.Rdb.Set(context.Background(), key, code, time.Minute*2)

	return &__.SendSmsResponse{}, nil
}

func (c *UserServer) Login(_ context.Context, in *__.LoginRequest) (*__.LoginResponse, error) {

	if !isValidMobile(in.Mobile) {
		return nil, fmt.Errorf("手机号格式不正确")
	}

	if in.SendSmsCode == "" {
		return nil, fmt.Errorf("验证码不能为空")
	}

	var user model.VideoUser
	global.DB.Where("mobile=?", in.Mobile).Find(&user)

	if user.Id == 0 {
		var count int64
		global.DB.Where("name = ?", "用户"+in.Mobile).Count(&count)
		if count > 0 {
			return nil, fmt.Errorf("用户名已存在")
		}

		newUser := model.VideoUser{
			Name:   "用户" + in.Mobile,
			Mobile: in.Mobile,
			Status: strconv.Itoa(1),
		}

		result := global.DB.Create(&newUser)
		if result.Error != nil {
			return nil, fmt.Errorf("注册失败: %v", result.Error)
		}
		user = newUser
	}

	key := "SendSms" + in.Mobile
	get := global.Rdb.Get(context.Background(), key)
	if get.Err() != nil {
		return nil, fmt.Errorf("验证码已过期")
	}

	if get.Val() != in.SendSmsCode {

		errorCountKey := "LoginError_" + in.Mobile
		count, _ := global.Rdb.Incr(context.Background(), errorCountKey).Result()
		global.Rdb.Expire(context.Background(), errorCountKey, time.Hour)

		if count > 5 {
			return nil, fmt.Errorf("验证码错误次数过多，请稍后再试")
		}

		return nil, fmt.Errorf("验证码错误")
	}

	global.Rdb.Del(context.Background(), key)

	global.Rdb.Del(context.Background(), "LoginError_"+in.Mobile)

	return &__.LoginResponse{
		Id: int64(user.Id),
	}, nil
}

func (c *UserServer) PublishContent(_ context.Context, in *__.PublishContentRequest) (*__.PublishContentResponse, error) {
	work := model.VideoWorks{
		Id:          int(in.UserId),
		Title:       in.Title,
		Desc:        in.Desc,
		MusicId:     int(in.MusicId),
		WorkType:    in.WorkType,
		CheckStatus: in.CheckStatus,
		CheckUser:   int(in.CheckUser),
		IpAddress:   in.IpAddress,
	}
	if err := global.DB.Create(&work).Error; err != nil {
		return nil, fmt.Errorf("作品发布失败")
	}

	return &__.PublishContentResponse{
		ContentId: int64(work.Id),
		Status:    "待审核状态",
	}, nil

}

func (c *UserServer) UpdateStatus(_ context.Context, in *__.UpdateStatusRequest) (*__.UpdateStatusResponse, error) {
	work := model.VideoWorks{
		Id:          int(in.Id),
		CheckStatus: in.CheckStatus,
	}

	if err := global.DB.Updates(&work).Error; err != nil {
		return nil, fmt.Errorf("审核失败")
	}

	return &__.UpdateStatusResponse{}, nil

}
