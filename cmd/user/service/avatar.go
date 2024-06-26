package service

import (
	"bibi/cmd/user/dal/db"
	"bibi/config"
	"bibi/kitex_gen/user"
	"bytes"
	"log"
	"strconv"
)

func (s *AvatarService) UploadAvatar(req *user.AvatarRequest) error {
	avatarReader := bytes.NewReader(req.AvatarFile)
	err := s.bucket.PutObject(config.OSS.MainDirectory+"/"+strconv.FormatInt(req.UserId, 10)+".jpg", avatarReader)
	if err != nil {
		log.Fatalf("upload file error:%video\n", err)
	}
	return err
}

func (s *AvatarService) PutAvatar(id int64, avatarUrl string) (*db.User, error) {
	userModel := &db.User{
		ID:     id,
		Avatar: avatarUrl,
	}
	return db.PutAvatar(s.ctx, userModel)
}
