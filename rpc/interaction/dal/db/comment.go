package db

import (
	"bibi/kitex_gen/interaction"
	"context"
	"errors"
	"gorm.io/gorm"
	"time"
)

//go:generate msgp -tests=false -o=comment_msgp.go -io=false
type Comment struct {
	ID        int64          `msg:"i"`
	VideoID   int64          `msg:"v"`
	ParentID  int64          `msg:"p"`
	Uid       int64          `msg:"u"`
	Content   string         `msg:"c"`
	CreatedAt time.Time      `msg:"pu"`
	UpdatedAt time.Time      `msg:"-"`             //ignore
	DeletedAt gorm.DeletedAt `sql:"index" msg:"-"` //ignore
}

func IsParentExist(ctx context.Context, commentModel *interaction.Comment) (bool, error) {
	var comment = &Comment{
		ID:      *commentModel.ParentId,
		VideoID: commentModel.VideoId,
	}
	err := DBComment.WithContext(ctx).Take(comment).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}
	return true, err
}

func IsCommentExist(ctx context.Context, commentModel *interaction.Comment) (bool, error) {
	var comment = &Comment{
		ID:      commentModel.Id,
		VideoID: commentModel.VideoId,
		Uid:     commentModel.User.Id,
	}
	err := DBComment.WithContext(ctx).Take(comment).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}
	return true, err
}

func CreateComment(ctx context.Context, commentModel *interaction.Comment) (*Comment, error) {
	var comment = &Comment{
		VideoID: commentModel.VideoId,
		Uid:     commentModel.User.Id,
		Content: commentModel.Content,
	}
	if commentModel.ParentId != nil {
		comment.ParentID = *commentModel.ParentId
	}

	if err := DBComment.WithContext(ctx).Create(comment).Error; err != nil {
		return &Comment{}, err
	}
	return comment, nil
}

func DeleteComment(ctx context.Context, commentModel *interaction.Comment) (*Comment, error) {
	var comment = &Comment{
		ID:      commentModel.Id,
		VideoID: commentModel.VideoId,
		Uid:     commentModel.User.Id,
	}
	if err := DBComment.WithContext(ctx).Take(comment).Delete(comment).Error; err != nil {
		return nil, err
	}
	return comment, nil
}

func GetCommentCount(ctx context.Context, videoId int64) (count int64, err error) {
	if err = DBComment.WithContext(ctx).Where("video_id = ?", videoId).Count(&count).Error; err != nil {
		return 0, err
	}
	return
}

func GetCommentsByVideoID(ctx context.Context, videoId int64) ([]Comment, int64, error) {
	comments := new([]Comment)
	var count int64
	if err := DBComment.WithContext(ctx).Where("video_id = ?", videoId).Count(&count).Find(comments).Error; err != nil {
		return nil, 0, err
	}
	return *comments, count, nil
}
