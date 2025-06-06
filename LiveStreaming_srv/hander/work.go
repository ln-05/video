package hander

import (
	"LiveStreaming_srv/basic/global"
	"LiveStreaming_srv/model"
	__ "LiveStreaming_srv/proto"
	"context"
	"fmt"
	"log"
)

func (c *UserServer) PublishContent(_ context.Context, in *__.PublishContentRequest) (*__.PublishContentResponse, error) {
	work := model.VideoWorks{
		Title:     in.Title,
		Desc:      in.Desc,
		MusicId:   int(in.MusicId),
		WorkType:  in.WorkType,
		IpAddress: in.IpAddress,
	}
	if err := global.DB.Create(&work).Error; err != nil {
		return nil, fmt.Errorf("作品发布失败")
	}

	return &__.PublishContentResponse{
		ContentId: int64(work.Id),
		Status:    "待审核状态",
	}, nil

}

func (c *UserServer) ListWork(_ context.Context, in *__.ListWorkRequest) (*__.ListWorkResponse, error) {
	var work []model.VideoWorks
	var worklist []*__.ListWork

	page := in.Page
	if page <= 0 {
		page = 1
	}

	pageSize := in.Size
	switch {
	case pageSize > 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	offset := (page - 1) * pageSize
	global.DB.Offset(int(offset)).Limit(int(pageSize)).Find(&work)

	for _, m := range work {
		worklist = append(worklist, &__.ListWork{
			Title:        m.Title,
			Desc:         m.Desc,
			MusicId:      int64(m.MusicId),
			WorkType:     m.WorkType,
			IpAddress:    m.IpAddress,
			LikeCount:    int64(m.LikeCount),
			CommentCount: int64(m.CommentCount),
			ShareCount:   int64(m.ShareCount),
			CollectCount: int64(m.CollectCount),
		})
	}

	return &__.ListWorkResponse{List: worklist}, nil

}

func (c *UserServer) InfoWork(_ context.Context, in *__.InfoWorkRequest) (*__.InfoWorkResponse, error) {
	if in.Id <= 0 {
		log.Printf("无效的作品ID: %d", in.Id)
		return nil, nil
	}
	var work model.VideoWorks

	if err := global.DB.Where("id=?", in.Id).First(&work).Error; err != nil {
		return nil, fmt.Errorf("没有该作品")
	}

	return &__.InfoWorkResponse{
		Title:        work.Title,
		Desc:         work.Desc,
		MusicId:      int64(work.MusicId),
		WorkType:     work.WorkType,
		IpAddress:    work.IpAddress,
		LikeCount:    int64(work.LikeCount),
		CommentCount: int64(work.CommentCount),
		ShareCount:   int64(work.ShareCount),
		CollectCount: int64(work.CollectCount),
	}, nil

}
