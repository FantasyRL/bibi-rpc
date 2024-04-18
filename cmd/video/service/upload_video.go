package service

import (
	"bibi/cmd/video/dal/db"
	"bibi/config"
	"bytes"
	"log"
)

func (s *VideoService) UploadCover(cover []byte, name string) error {
	coverReader := bytes.NewReader(cover)
	err := s.bucket.PutObject(config.OSS.MainDirectory+"/video/"+name, coverReader)
	if err != nil {
		log.Fatalf("upload file error:%video\n", err)
	}
	return err
}

func (s *VideoService) UploadVideo(video []byte, name string) error {
	videoReader := bytes.NewReader(video)
	err := s.bucket.PutObject(config.OSS.MainDirectory+"/video/"+name, videoReader)
	if err != nil {
		log.Fatalf("upload file error:%video\n", err)
	}
	return err
}

func (s *VideoService) PutVideo(video *db.Video) (*db.Video, error) {
	return db.CreateVideo(s.ctx, video)
}
