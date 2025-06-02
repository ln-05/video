package model

type VideoWorkObject struct {
	Id       int `gorm:"column:id;type:int;primaryKey;not null;" json:"id"`
	WorkId   int `gorm:"column:work_id;type:int;not null;" json:"work_id"`
	ObjectId int `gorm:"column:object_id;type:int;comment:作品对象ID;not null;" json:"object_id"` // 作品对象ID
}

func (v *VideoWorkObject) TableName() string {
	return "video_work_object"
}
