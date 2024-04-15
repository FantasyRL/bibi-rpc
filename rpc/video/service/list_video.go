package service

import (
	"bibi/kitex_gen/video"
	"bibi/rpc/video/dal/db"
)

func (s *VideoService) ListVideo(req *video.ListUserVideoRequest) (*[]db.Video, int64, error) {
	return db.ListVideosByID(s.ctx, int(req.PageNum), req.UserId)
}
