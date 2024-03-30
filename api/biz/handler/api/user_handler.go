// Code generated by hertz generator.

package api

import (
	"bibi/api/biz/rpc"
	"bibi/kitex_gen/user"
	"bibi/pkg/errno"
	"bibi/pkg/pack"
	"context"
	"path/filepath"

	api "bibi/api/biz/model/api"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// Register .
// @Summary Register
// @Description userRegister
// @Accept json/form
// @Produce json
// @Param username query string true "用户名"
// @Param email query string true "邮箱"
// @Param password query string true "密码"
// @router /bibi/user/register/ [POST]
func Register(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.RegisterRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.RegisterResponse)

	rpcResp, err := rpc.UserRegister(ctx, &user.RegisterRequest{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		pack.SendRPCFailResp(c, err)
		return
	}

	resp.Base = pack.ConvertToAPIBaseResp(rpcResp.Base)
	if resp.Base.Code != errno.SuccessCode {
		c.JSON(consts.StatusOK, resp)
		return
	}
	resp.UserID = rpcResp.UserId
	c.JSON(consts.StatusOK, resp)
}

// Login .
// @Summary Login
// @Description login to get your auth token
// @Accept json/form
// @Produce json
// @Param username query string true "用户名"
// @Param password query string true "密码"
// @Param otp query string false "otp"
// @router /bibi/user/login/ [POST]
func Login(ctx context.Context, c *app.RequestContext) {
	resp := new(api.LoginResponse)

	resp.Base = pack.ConvertToAPIBaseResp(pack.BuildBaseResp(nil))
	//hertz jwt(mw)
	v1, _ := c.Get("user")
	resp.User = pack.ToUserResp(v1)
	//hertz jwt(mw)
	v2, _ := c.Get("access-token")
	v3, _ := c.Get("refresh-token")
	at := v2.(string)
	rt := v3.(string)
	resp.AccessToken = &at
	resp.RefreshToken = &rt

	c.JSON(consts.StatusOK, resp)
}

// Info .
// @Summary Info
// @Description get user's info
// @Accept json/form
// @Produce json
// @Param user_id query string true "用户id"
// @Param access-token header string false "access-token"
// @Param refresh-token header string false "refresh-token"
// @router /bibi/user/info [GET]
func Info(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.InfoRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.InfoResponse)

	rpcResp, err := rpc.UserInfo(ctx, &user.InfoRequest{})
	if err != nil {
		pack.SendRPCFailResp(c, err)
		return
	}

	resp.Base = pack.ConvertToAPIBaseResp(rpcResp.Base)
	if resp.Base.Code != errno.SuccessCode {
		c.JSON(consts.StatusOK, resp)
		return
	}

	resp.User = pack.ConvertToAPIUser(rpcResp.User)
	c.JSON(consts.StatusOK, resp)
}

// Avatar .
// @Summary PutAvatar
// @Description revise user's avatar
// @Accept json/form
// @Produce json
// @Param avatar_file formData file true "头像"
// @Param access-token header string false "access-token"
// @Param refresh-token header string false "refresh-token"
// @router /bibi/user/avatar/upload [PUT]
func Avatar(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.AvatarRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	file, err := c.FormFile("avatar_file")

	resp := new(api.AvatarResponse)

	//判断文件格式
	fileExt := filepath.Ext(file.Filename)
	allowExtMap := map[string]bool{
		".jpg":  true,
		".png":  true,
		".jpeg": true,
	}
	if !pack.IsAllowExt(fileExt, allowExtMap) {
		resp.Base = pack.ConvertToAPIBaseResp(pack.BuildBaseResp(errno.ParamError))
		c.JSON(consts.StatusOK, resp)
		return
	}

	v, ok := c.Get("current_user_id")
	if !ok {
		err = errno.ParamError
	}
	id, _ := v.(int64)

	fileBinary, err := pack.FileToByte(file)
	if err != nil {
		resp.Base = pack.ConvertToAPIBaseResp(pack.BuildBaseResp(errno.ReadFileError))
		c.JSON(consts.StatusOK, resp)
		return
	}

	rpcResp, err := rpc.UserAvatar(ctx, &user.AvatarRequest{
		UserId:     id,
		AvatarFile: fileBinary,
	})

	if err != nil {
		pack.SendRPCFailResp(c, err)
		return
	}

	resp.Base = pack.ConvertToAPIBaseResp(rpcResp.Base)
	if resp.Base.Code != errno.SuccessCode {
		c.JSON(consts.StatusOK, resp)
		return
	}

	resp.User = pack.ConvertToAPIUser(rpcResp.User)

	c.JSON(consts.StatusOK, resp)
}

// Switch2FA .
// @Summary switch_2fa
// @Description switch on/off 2fa mode
// @Accept json/form
// @Produce json
// @Param action_type query int true "关闭:0;开启:1"
// @Param totp query string false "totp"
// @Param access-token header string false "access-token"
// @Param refresh-token header string false "refresh-token"
// @router /bibi/user/switch2fa [POST]
func Switch2FA(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.Switch2FARequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.Switch2FAResponse)

	c.JSON(consts.StatusOK, resp)
}

// GetAccessToken .
// @Summary get_access-token
// @Description get available access-token by refresh-token
// @Accept json/form
// @Produce json
// @Param refresh-token header string true "refresh-token"
// @router /bibi/access_token/get [GET]
func GetAccessToken(ctx context.Context, c *app.RequestContext) {
	resp := new(api.GetAccessTokenResponse)

	resp.Base = pack.ConvertToAPIBaseResp(pack.BuildBaseResp(nil))

	//hertz jwt(mw)
	v2, _ := c.Get("access-token")
	at := v2.(string)
	resp.AccessToken = &at

	c.JSON(consts.StatusOK, resp)
}
