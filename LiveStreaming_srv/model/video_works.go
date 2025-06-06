package model

import (
	"gorm.io/gorm"
	"time"
)

type VideoWorks struct {
	Id             int            `gorm:"column:id;type:int;primaryKey;not null;" json:"id"`
	Title          string         `gorm:"column:title;type:varchar(100);comment:标题;default:NULL;" json:"title"`                      // 标题
	Desc           string         `gorm:"column:desc;type:varchar(255);comment:描述;default:NULL;" json:"desc"`                        // 描述
	MusicId        int            `gorm:"column:music_id;type:bigint;comment:选择音乐;default:NULL;" json:"music_id"`                    // 选择音乐
	WorkType       string         `gorm:"column:work_type;type:varchar(20);comment:作品类型;default:NULL;" json:"work_type"`             // 作品类型
	CheckStatus    string         `gorm:"column:check_status;type:varchar(10);comment:审核状态;default:NULL;" json:"check_status"`       // 审核状态
	CheckUser      int            `gorm:"column:check_user;type:bigint;comment:审核人;default:NULL;" json:"check_user"`                 // 审核人
	IpAddress      string         `gorm:"column:ip_address;type:varchar(20);comment:IP地址;default:NULL;" json:"ip_address"`           // IP地址
	WorkPermission string         `gorm:"column:work_permission;type:varchar(20);comment:作品权限;default:NULL;" json:"work_permission"` // 作品权限
	LikeCount      int            `gorm:"column:like_count;type:bigint;comment:喜欢数量;default:NULL;" json:"like_count"`                // 喜欢数量
	CommentCount   int            `gorm:"column:comment_count;type:bigint;comment:评论数;default:NULL;" json:"comment_count"`           // 评论数
	ShareCount     int            `gorm:"column:share_count;type:bigint;comment:分享数;default:NULL;" json:"share_count"`               // 分享数
	CollectCount   int            `gorm:"column:collect_count;type:bigint;comment:收藏数;default:NULL;" json:"collect_count"`           // 收藏数
	BrowseCount    int            `gorm:"column:browse_count;type:bigint;comment:浏览量;default:NULL;" json:"browse_count"`             // 浏览量
	CreatedAt      time.Time      `gorm:"column:created_at;type:datetime(3);default:NULL;" json:"created_at"`
	UpdatedAt      time.Time      `gorm:"column:updated_at;type:datetime(3);default:NULL;" json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"column:deleted_at;type:datetime(3);default:NULL;" json:"deleted_at"`
	CreatedBy      uint64         `gorm:"column:created_by;type:bigint UNSIGNED;comment:创建者;default:NULL;" json:"created_by"` // 创建者
	UpdatedBy      uint64         `gorm:"column:updated_by;type:bigint UNSIGNED;comment:更新者;default:NULL;" json:"updated_by"` // 更新者
	DeletedBy      uint64         `gorm:"column:deleted_by;type:bigint UNSIGNED;comment:删除者;default:NULL;" json:"deleted_by"` // 删除者
}

func (v *VideoWorks) TableName() string {
	return "video_works"

}
