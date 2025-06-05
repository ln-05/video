package request

type PublishContents struct {
	Title     string `form:"title" binding:"required"`
	Desc      string `form:"desc" binding:"required"`
	MusicId   int64  `form:"musicid" binding:"required"`
	WorkType  string `form:"worktype" binding:"required"`
	IpAddress string `form:"ipaddress" binding:"required"`
}
type ListWork struct {
	Page int64 `form:"page" binding:"required"`
	Size int64 `form:"size" binding:"required"`
}
type InfoWork struct {
	Id int64 `form:"id" binding:"required"`
}
