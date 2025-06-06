package request

type SendSms struct {
	Mobile      string `form:"mobile" binding:"required"`
	SendSmsCode string `form:"sendSmsCode" binding:"required"`
}
type Login struct {
	Mobile      string `form:"mobile" binding:"required"`
	SendSmsCode string `form:"sendSmsCode" binding:"required"`
}

type Personal struct {
	Id int64 `form:"id" binding:"required"`
}

type UpdatePersonal struct {
	Id            int64  `form:"id" binding:"required"`
	Name          string `form:"name" binding:"required"`
	NickName      string `form:"nick_name" binding:"required"`
	Signature     string `form:"signature" binding:"required"`
	Sex           string `form:"sex" binding:"required"`
	Constellation string `form:"constellation" binding:"required"`
	AvatorFileId  int64  `form:"avator_file_id" binding:"required"`
	Mobile        string `form:"mobile" binding:"required"`
	Age           int64  `form:"age" binding:"required"`
	OnlineStatus  string `form:"online_status" binding:"required"`
}
