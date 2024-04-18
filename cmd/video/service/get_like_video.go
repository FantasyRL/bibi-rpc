package service

import (
	"bibi/cmd/video/dal/db"
)

//todo:rpc
func (s *VideoService) GetLikeVideoList(videoIdList []int64) ([]db.Video, error /*,likeList []int64, isLikeList []int64*/) {
	return db.GetVideoByIdList(s.ctx, videoIdList)
}
