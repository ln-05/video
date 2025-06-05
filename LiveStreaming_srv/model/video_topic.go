package model

type VideoTopic struct {
	Id         int    `gorm:"column:id;type:int;primaryKey;not null;" json:"id"`
	Title      string `gorm:"column:title;type:varchar(50);comment:话题标题;default:NULL;" json:"title"`      // 话题标题
	CreateUser int    `gorm:"column:create_user;type:int;comment:话题发起人;default:NULL;" json:"create_user"` // 话题发起人
}

func (v *VideoTopic) TableName() string {
	return "video_topic"
}
