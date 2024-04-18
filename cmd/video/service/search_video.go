package service

import (
	"bibi/cmd/video/dal/db"
	"bibi/kitex_gen/video"
)

func (s *VideoService) SearchVideo(req *video.SearchVideoRequest) (*[]db.Video, int64, error) {
	return db.SearchVideo(s.ctx, int(req.PageNum), req.Param)
}
