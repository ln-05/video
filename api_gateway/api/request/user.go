package request

type SendSms struct {
	Mobile      string `form:"mobile" binding:"required"`
	SendSmsCode string `form:"sendSmsCode" binding:"required"`
}
type Login struct {
	Mobile      string `form:"mobile" binding:"required"`
	SendSmsCode string `form:"sendSmsCode" binding:"required"`
}
