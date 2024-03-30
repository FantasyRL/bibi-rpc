package main

import (
	user "bibi/kitex_gen/user"
	"bibi/kitex_gen/user/userhandler"
	"bibi/pkg/constants"
	"bibi/pkg/pack"
	"bibi/rpc/user/service"
	"context"
	"github.com/cloudwego/kitex/client"
)

// UserHandlerImpl implements the last service interface defined in the IDL.
type UserHandlerImpl struct {
	userCli userhandler.Client
}

func NewUserClient(addr string) (userhandler.Client, error) {
	return userhandler.NewClient(constants.UserServiceName, client.WithHostPorts(addr))
}

// Register implements the UserHandlerImpl interface.
func (s *UserHandlerImpl) Register(ctx context.Context, req *user.RegisterRequest) (resp *user.RegisterResponse, err error) {
	//不开辟空间就会死
	resp = new(user.RegisterResponse)

	userResp, err := service.NewUserService(ctx).Register(req)

	resp.Base = pack.BuildBaseResp(err)
	if err != nil {
		return resp, nil
	}

	resp.UserId = &userResp.ID

	return resp, nil
}

// Login implements the UserHandlerImpl interface.
func (s *UserHandlerImpl) Login(ctx context.Context, req *user.LoginRequest) (resp *user.LoginResponse, err error) {
	resp = new(user.LoginResponse)

	userResp, err := service.NewUserService(ctx).Login(req)

	resp.Base = pack.BuildBaseResp(err)
	if err != nil {
		return resp, nil
	}
	resp.User = service.BuildUserResp(userResp)
	return resp, nil
}

// Info implements the UserHandlerImpl interface.
func (s *UserHandlerImpl) Info(ctx context.Context, req *user.InfoRequest) (resp *user.InfoResponse, err error) {
	// TODO: Your code here...
	return
}

// Avatar implements the UserHandlerImpl interface.
func (s *UserHandlerImpl) Avatar(ctx context.Context, req *user.AvatarRequest) (resp *user.AvatarResponse, err error) {
	// TODO: Your code here...
	return
}

// Switch2FA implements the UserHandlerImpl interface.
func (s *UserHandlerImpl) Switch2FA(ctx context.Context, req *user.Switch2FARequest) (resp *user.Switch2FAResponse, err error) {
	// TODO: Your code here...
	return
}
