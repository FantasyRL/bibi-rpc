// Code generated by hertz generator.

package api

import (
	"bibi/cmd/api/biz/rpc"
	"bibi/kitex_gen/follow"
	"bibi/pkg/errno"
	"bibi/pkg/pack"
	"context"

	api "bibi/cmd/api/biz/model/api"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// FollowAction .
// @Summary follow_action
// @Description follow action
// @Accept json/form
// @Produce json
// @Param object_uid query int true "操作对象id"
// @Param action_type query int true "0：取消关注;1：关注"
// @Param access-token header string false "access-token"
// @Param refresh-token header string false "refresh-token"
// @router /bibi/follow/action [POST]
func FollowAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.FollowActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.FollowActionResponse)

	v, _ := c.Get("current_user_id")
	id := v.(int64)
	if req.ObjectUID == id {
		resp.Base = pack.BuildAPIBaseResp(errno.FollowMyselfError)
		c.JSON(consts.StatusOK, resp)
		return
	}
	rpcResp, err := rpc.FollowAction(ctx, &follow.FollowActionRequest{
		ObjectUid:  req.ObjectUID,
		ActionType: req.ActionType,
		UserId:     id,
	})
	if err != nil {
		pack.SendRPCFailResp(c, err)
	}
	resp.Base = pack.ConvertToAPIBaseResp(rpcResp.Base)
	c.JSON(consts.StatusOK, resp)
}

// FollowerList .
// @Summary follower_list
// @Description list your followers
// @Accept json/form
// @Produce json
// @Param page_num query int64 true "页码"
// @Param access-token header string false "access-token"
// @Param refresh-token header string false "refresh-token"
// @router /bibi/follow/follower [GET]
func FollowerList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.FollowerListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.FollowerListResponse)

	v, _ := c.Get("current_user_id")
	id := v.(int64)
	rpcResp, err := rpc.FollowerList(ctx, &follow.FollowerListRequest{
		PageNum: req.PageNum,
		UserId:  id,
	})
	if err != nil {
		pack.SendRPCFailResp(c, err)
	}
	resp.Base = pack.ConvertToAPIBaseResp(rpcResp.Base)
	if rpcResp.Base.Code != errno.SuccessCode {
		c.JSON(consts.StatusOK, resp)
	}
	resp.Count = rpcResp.Count
	resp.FollowerList = pack.ConvertToAPIUsers(rpcResp.FollowerList)
	c.JSON(consts.StatusOK, resp)
}

// FollowingList .
// @Summary following_list
// @Description list your followed
// @Accept json/form
// @Produce json
// @Param page_num query int64 true "页码"
// @Param access-token header string false "access-token"
// @Param refresh-token header string false "refresh-token"
// @router /bibi/follow/following [GET]
func FollowingList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.FollowingListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.FollowingListResponse)

	v, _ := c.Get("current_user_id")
	id := v.(int64)
	rpcResp, err := rpc.FollowingList(ctx, &follow.FollowingListRequest{
		PageNum: req.PageNum,
		UserId:  id,
	})
	if err != nil {
		pack.SendRPCFailResp(c, err)
	}
	resp.Base = pack.ConvertToAPIBaseResp(rpcResp.Base)
	if rpcResp.Base.Code != errno.SuccessCode {
		c.JSON(consts.StatusOK, resp)
	}
	resp.Count = rpcResp.Count
	resp.FollowingList = pack.ConvertToAPIUsers(rpcResp.FollowingList)
	c.JSON(consts.StatusOK, resp)
}

// FriendList .
// @Summary friend_list
// @Description list your friends
// @Accept json/form
// @Produce json
// @Param page_num query int64 true "页码"
// @Param access-token header string false "access-token"
// @Param refresh-token header string false "refresh-token"
// @router /bibi/follow/friend [GET]
func FriendList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.FriendListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.FriendListResponse)

	v, _ := c.Get("current_user_id")
	id := v.(int64)
	rpcResp, err := rpc.FriendList(ctx, &follow.FriendListRequest{
		PageNum: req.PageNum,
		UserId:  id,
	})
	if err != nil {
		pack.SendRPCFailResp(c, err)
	}
	resp.Base = pack.ConvertToAPIBaseResp(rpcResp.Base)
	if rpcResp.Base.Code != errno.SuccessCode {
		c.JSON(consts.StatusOK, resp)
	}
	resp.Count = rpcResp.Count
	resp.FriendList = pack.ConvertToAPIUsers(rpcResp.FriendList)
	c.JSON(consts.StatusOK, resp)
}
