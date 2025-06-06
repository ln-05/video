package hander

import (
	"LiveStreaming_srv/basic/global"
	"LiveStreaming_srv/model"
	__ "LiveStreaming_srv/proto"
	"context"
	"gorm.io/gorm"
	"log"
)

func (c *UserServer) PostComment(_ context.Context, in *__.PostCommentRequest) (*__.PostCommentResponse, error) {
	log.Printf("收到发布评论请求: work_id=%d, user_id=%d", in.WorkId, in.UserId)

	if in.WorkId <= 0 {
		return &__.PostCommentResponse{
			Success: false,
			Message: "无效的作品ID",
		}, nil
	}
	if in.UserId <= 0 {
		return &__.PostCommentResponse{
			Success: false,
			Message: "无效的用户ID",
		}, nil
	}
	if in.Content == "" {
		return &__.PostCommentResponse{
			Success: false,
			Message: "评论内容不能为空",
		}, nil
	}

	work := model.VideoWorkComment{
		WorkId:  int(in.WorkId),
		UserId:  int(in.UserId),
		Content: in.Content,
		Tag:     int(in.Tag),
		Pid:     int(in.Pid),
	}

	if err := global.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&work).Error; err != nil {
			return err
		}

		if err := tx.Model(&model.VideoWorks{}).Where("id = ?", in.WorkId).Update("comment_count", gorm.Expr("comment_count + 1")).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		log.Printf("发布评论失败: %v", err)
		return &__.PostCommentResponse{
			Success: false,
			Message: "发布评论失败: " + err.Error(),
		}, nil
	}

	return &__.PostCommentResponse{
		Success: true,
		Message: "发布评论成功",
	}, nil
}
