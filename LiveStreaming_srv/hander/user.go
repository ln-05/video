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

//func (c *UserServer) SendSms(ctx context.Context, in *__.SendSmsRequest) (*__.SendSmsResponse, error) {
//	code := rand.Intn(9000) + 1000
//	global.Rdb.Set(context.Background(), "SendSms"+in.Mobile, code, time.Minute*2)
//	return &__.SendSmsResponse{}, nil
//}
//
//func (c *UserServer) Login(_ context.Context, in *__.LoginRequest) (*__.LoginResponse, error) {
//	var user model.VideoUser
//	global.DB.Where("mobile=?", in.Mobile).Find(&user)
//
//	if user.Id == 0 {
//
//		newUser := model.VideoUser{
//			Name:   "用户" + in.Mobile,
//			Mobile: in.Mobile,
//		}
//
//		result := global.DB.Create(&newUser)
//		if result.Error != nil {
//			return nil, fmt.Errorf("注册失败: %v", result.Error)
//		}
//
//		user = newUser
//	}
//
//	get := global.Rdb.Get(context.Background(), "SendSms"+in.Mobile)
//	if get.Val() != in.SendSmsCode {
//		return nil, fmt.Errorf("验证码错误")
//	}
//
//	return &__.LoginResponse{
//		Id: int64(user.Id),
//	}, nil
//}

func isValidMobile(mobile string) bool {
	pattern := `^1[3-9]\d{9}$`
	re := regexp.MustCompile(pattern)
	return re.MatchString(mobile)
}

func (c *UserServer) SendSms(ctx context.Context, in *__.SendSmsRequest) (*__.SendSmsResponse, error) {
	if !isValidMobile(in.Mobile) {
		return nil, fmt.Errorf("手机号格式不正确")
	}

	var userCount int64
	global.DB.Model(&model.VideoUser{}).Where("mobile = ?", in.Mobile).Count(&userCount)

	key := "SendSms" + in.Mobile
	existingCode := global.Rdb.Get(context.Background(), key)
	if existingCode.Err() == nil {
		return nil, fmt.Errorf("验证码已发送，请稍后再试")
	}

	code := rand.Intn(9000) + 1000

	sendCountKey := "SendSmsCount_" + in.Mobile
	count, err := global.Rdb.Incr(context.Background(), sendCountKey).Result()
	if err == nil && count > 5 { // 限制每天发送次数
		return nil, fmt.Errorf("验证码发送过于频繁，请明天再试")
	}
	global.Rdb.Expire(context.Background(), sendCountKey, 24*time.Hour) // 设置24小时过期

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
		global.DB.Model(&model.VideoUser{}).Where("name = ?", "用户"+in.Mobile).Count(&count)
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
