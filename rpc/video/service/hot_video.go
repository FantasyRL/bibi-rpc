package service

import (
	"bibi/kitex_gen/video"
	"bibi/rpc/video/dal/db"
)

func (s *VideoService) HotVideo(req *video.HotVideoRequest) ([]db.Video, error) {
	//videoIdList, err := cache.ListHotVideo(s.ctx)
	//if err != nil {
	//	return nil, err
	//}
	//videoList, err := db.GetVideoByIdList(videoIdList)
	//if err != nil {
	//	return nil, err
	//}
	//return videoList, nil
	return nil, nil
}
