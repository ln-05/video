package request

type SendSms struct {
	Mobile      string `form:"mobile" binding:"required"`
	SendSmsCode string `form:"sendSmsCode" binding:"required"`
}
type Login struct {
	Mobile      string `form:"mobile" binding:"required"`
	SendSmsCode string `form:"sendSmsCode" binding:"required"`
}

type PublishContents struct {
	UserId      int64  `form:"userid" binding:"required"`
	Title       string `form:"title" binding:"required"`
	Desc        string `form:"desc" binding:"required"`
	MusicId     int64  `form:"musicid" binding:"required"`
	WorkType    string `form:"worktype" binding:"required"`
	CheckStatus string `form:"checkstatus" binding:"required"`
	CheckUser   int64  `form:"checkuser" binding:"required"`
	IpAddress   string `form:"ipaddress" binding:"required"`
}
type UpdateStatus struct {
	Id          int64  `form:"id" binding:"required"`
	CheckStatus string `form:"checkstatus" binding:"required"`
}
