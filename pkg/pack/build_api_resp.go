package pack

import (
	"bibi/cmd/api/biz/model/api"
	"bibi/kitex_gen/base"
	"bibi/pkg/errno"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func BuildBaseResp(err error) *base.BaseResp {
	if err == nil {
		return ErrToResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return ErrToResp(e)
	}

	_e := errno.ServiceError.WithMessage(err.Error()) //未知错误
	return ErrToResp(_e)
}

func ErrToResp(err errno.ErrNo) *base.BaseResp {
	return &base.BaseResp{
		Code: err.ErrorCode,
		Msg:  err.ErrorMsg,
	}
}

func ConvertToAPIBaseResp(baseResp *base.BaseResp) *api.BaseResp {
	return &api.BaseResp{
		Code: baseResp.Code,
		Msg:  baseResp.Msg,
	}
}

func SendRPCFailResp(c *app.RequestContext, err error) {
	c.JSON(consts.StatusOK, api.BaseResp{
		Code: -1,
		Msg:  errno.ConvertErr(err).Error(),
	})
}

func ConvertToAPIUser(kitexUser *base.User) *api.User {
	return &api.User{
		ID:            kitexUser.Id,
		Name:          kitexUser.Name,
		Email:         kitexUser.Email,
		FollowCount:   kitexUser.FollowerCount,
		FollowerCount: kitexUser.FollowerCount,
		//IsFollow: kitexUser.IsFollow,
		Avatar:     kitexUser.Avatar,
		VideoCount: kitexUser.VideoCount,
	}
}

func ToUserResp(_user interface{}) *api.User {
	//这里使用了一个及其抽象的断言
	p, _ := (_user).(*base.User)
	return &api.User{
		ID:     p.Id,
		Name:   p.Name,
		Email:  p.Email,
		Avatar: p.Avatar,
	}
}

func ConvertToAPIVideo(kitexVideo *base.Video) *api.Video {
	var isLike bool = false
	if kitexVideo.IsLike == 1 {
		isLike = true
	}
	return &api.Video{
		ID:           kitexVideo.Id,
		Title:        kitexVideo.Title,
		Author:       ConvertToAPIUser(kitexVideo.Author),
		UID:          kitexVideo.Uid,
		PlayURL:      kitexVideo.PlayUrl,
		CoverURL:     kitexVideo.CoverUrl,
		LikeCount:    kitexVideo.LikeCount,
		CommentCount: kitexVideo.CommentCount,
		IsLike:       isLike,
		PublishTime:  kitexVideo.PublishTime,
	}
}

func ConvertToAPIVideos(kitexVideos []*base.Video) []*api.Video {
	videosResp := make([]*api.Video, 0)
	for _, v := range kitexVideos {
		videosResp = append(videosResp, ConvertToAPIVideo(v))
	}
	return videosResp
}

func ConvertToAPIComment(kitexComment *base.Comment) *api.Comment {
	return &api.Comment{
		ID:          kitexComment.Id,
		VideoID:     kitexComment.VideoId,
		ParentID:    kitexComment.ParentId,
		User:        ConvertToAPIUser(kitexComment.User),
		Content:     kitexComment.Content,
		PublishTime: kitexComment.PublishTime,
	}
}

func ConvertToAPIComments(kitexComments []*base.Comment) []*api.Comment {
	commentsResp := make([]*api.Comment, len(kitexComments))
	for i, v := range kitexComments {
		commentsResp[i] = ConvertToAPIComment(v)
	}
	return commentsResp
}
