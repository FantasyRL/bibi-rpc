package service

import (
	"bibi/cmd/user/dal/db"
	"bibi/cmd/user/rpc"
	"bibi/cmd/user/vector/vector"
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

func (s *AvatarService) ConvertAvatarToMilvus(req *user.AvatarRequest, uid int64) error {
	resp, err := rpc.UserGetVector(s.ctx, &vector.GetVectorRequest{
		Image: req.AvatarFile,
	})
	if err != nil {
		return err
	}
	v := make([]float32, 0)
	for _, d := range resp.Vector {
		v = append(v, float32(d))
	}

	return db.InsertData(s.ctx, v, uid)
}

func (s *AvatarService) SearchAvatar(req *user.SearchAvatarRequest) (*[]db.User, error) {
	resp, err := rpc.UserGetVector(s.ctx, &vector.GetVectorRequest{
		Image: req.Picture,
	})
	if err != nil {
		return nil, err
	}
	v := make([]float32, 0)
	for _, d := range resp.Vector {
		v = append(v, float32(d))
	}
	uidList, err := db.Search(s.ctx, v)
	if err != nil {
		return nil, err
	}
	return db.QueryUserByIDList(s.ctx, uidList)
}
