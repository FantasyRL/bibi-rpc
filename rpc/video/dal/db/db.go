package db

import (
	"bibi/pkg/constants"
	"bibi/pkg/errno"
	"context"
	"gorm.io/gorm"
	"time"
)

type Video struct {
	ID        int64
	Uid       int64
	Title     string
	PlayUrl   string
	CoverUrl  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `sql:"index"`
}

func CreateVideo(ctx context.Context, video *Video) (*Video, error) {
	if err := DB.WithContext(ctx).Create(video).Error; err != nil {
		return nil, err
	}
	return video, nil
}

func GetVideoCountByID(uid int64) (count int64, err error) {
	if err = DB.Model(Video{}).Where("uid = ?", uid).Count(&count).Error; err != nil {
		return 114514, err
	}
	return
}

func ListVideosByID(ctx context.Context, pageNum int, uid int64) (*[]Video, int64, error) {
	videos := new([]Video)
	var count int64 = 0
	//按创建时间降序
	if err := DB.Model(&Video{}).Where("uid = ?", uid).Count(&count).Order("created_at DESC").
		Limit(constants.PageSize).Offset((pageNum - 1) * constants.PageSize).Find(videos).
		Error; err != nil {
		return nil, 114514, err
	}
	return videos, count, nil
}

func SearchVideo(ctx context.Context, pageNum int, param string) (*[]Video, int64, error) {
	videos := new([]Video)
	var count int64 = 0
	if err := DB.Model(&Video{}).
		Where("title LIKE ? ", "%"+param+"%").
		Count(&count).Limit(constants.PageSize).Offset((pageNum - 1) * constants.PageSize).Find(&videos).
		Error; err != nil {
		return nil, 114514, err
	}
	return videos, count, nil
}

func CheckVideoExistById(videoId int64) error {
	video := new(Video)
	if err := DB.Model(Video{}).Where("id = ?", videoId).Find(video).Error; err != nil {
		return err
	}
	if video != (&Video{}) {
		return errno.VideoNotExistError
	}
	return nil
}

func GetVideoByIdList(videoIdList []int64) ([]Video, error) {
	videos := new([]Video)
	//.Order("created_at DESC")没必要
	if err := DB.Model(Video{}).Where("id IN ?", videoIdList).Find(videos).Error; err != nil {
		return nil, err
	}
	return *videos, nil

}
