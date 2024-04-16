package service

import (
	"bibi/kitex_gen/user"
	aliyunoss "bibi/pkg/utils/oss"
	"bibi/rpc/user/dal/db"
	"context"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"log"
)

type UserService struct {
	ctx context.Context
}

type AvatarService struct {
	ctx    context.Context
	bucket *oss.Bucket
}

func NewUserService(ctx context.Context) *UserService {
	return &UserService{ctx: ctx}
}

func NewAvatarService(ctx context.Context) *AvatarService {
	bucket, err := aliyunoss.OSSBucketCreate()
	if err != nil {
		log.Fatal(err)
	}
	return &AvatarService{ctx: ctx, bucket: bucket}
}

func BuildUserResp(dbUser *db.User) *user.User {
	return &user.User{
		Id:     dbUser.ID,
		Name:   dbUser.UserName,
		Email:  dbUser.Email,
		Avatar: dbUser.Avatar,
	}
}

func BuildUsersResp(dbUsers *[]db.User) []*user.User {
	usersResp := make([]*user.User, len(*dbUsers))
	for i, v := range *dbUsers {
		usersResp[i] = BuildUserResp(&v)
	}
	return usersResp
}
