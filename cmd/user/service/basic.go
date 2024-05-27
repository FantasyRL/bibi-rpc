package service

import (
	"bibi/cmd/user/dal/db"
	"bibi/kitex_gen/user"
	"bibi/pkg/errno"
	"bibi/pkg/utils/otp2fa"
	"bibi/pkg/utils/pwd"
)

func (s *UserService) Register(req *user.RegisterRequest) (*db.User, error) {
	if len(req.Username) < 4 /*||len(req.Password)<8*/ {
		return nil, errno.ParamError
	}

	PwdDigest := pwd.SetPassword(req.Password)
	userModel := &db.User{
		UserName: req.Username,
		Email:    req.Email,
		Password: PwdDigest,
	}
	return db.Register(s.ctx, userModel)
}

func (s *UserService) Login(req *user.LoginRequest /*, stTracer opentracing.Tracer, parentSpan opentracing.Span*/) (*db.User, error) {
	userModel := &db.User{
		UserName: req.Username,
		Password: req.Password,
	}
	//childSpan1 := stTracer.StartSpan(
	//	"mysql",
	//	opentracing.ChildOf(parentSpan.Context()),
	//)
	userResp, err := db.Login(s.ctx, userModel)
	if err != nil {
		return nil, err
	}
	//childSpan1.Finish()
	//childSpan2 := stTracer.StartSpan(
	//	"otp",
	//	opentracing.ChildOf(childSpan1.Context()),
	//)
	if userResp.Type2fa == 1 && !otp2fa.CheckTotp(*req.Otp, userResp.Otp) {
		return nil, errno.Verify2FAError
	}
	//childSpan2.Finish()
	return userResp, nil
}

func (s *UserService) Info(req *user.InfoRequest) (*db.User, error) {
	userModel := &db.User{
		ID: req.UserId,
	}
	return db.QueryUserByID(s.ctx, userModel)
}
