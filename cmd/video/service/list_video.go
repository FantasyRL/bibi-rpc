package service

import (
	"bibi/cmd/video/dal/db"
	"bibi/kitex_gen/video"
)

func (s *VideoService) ListVideo(req *video.ListUserVideoRequest) (*[]db.Video, int64, error) {
	return db.ListVideosByID(s.ctx, int(req.PageNum), req.UserId)
}
