package model

type Realname struct {
	Userid   int    `gorm:"column:userid;type:int;primaryKey;not null;" json:"userid"`
	NickName string `gorm:"column:nick_name;type:varchar(255);default:NULL;" json:"nick_name"`
	Mobile   string `gorm:"column:mobile;type:char(11);default:NULL;" json:"mobile"`
}

func (r *Realname) TableName() string {
	return "realname"
}
