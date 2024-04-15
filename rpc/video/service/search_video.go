package service

import (
	"bibi/kitex_gen/video"
	"bibi/rpc/video/dal/db"
)

func (s *VideoService) SearchVideo(req *video.SearchVideoRequest) (*[]db.Video, int64, error) {
	return db.SearchVideo(s.ctx, int(req.PageNum), req.Param)
}
