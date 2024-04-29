package main

import (
	"bibi/cmd/follow/service"
	follow "bibi/kitex_gen/follow"
	"bibi/pkg/errno"
	"bibi/pkg/pack"
	"context"
)

// FollowHandlerImpl implements the last service interface defined in the IDL.
type FollowHandlerImpl struct{}

// FollowAction implements the FollowHandlerImpl interface.
func (s *FollowHandlerImpl) FollowAction(ctx context.Context, req *follow.FollowActionRequest) (resp *follow.FollowActionResponse, err error) {
	resp = new(follow.FollowActionResponse)
	if req.ObjectUid == req.UserId {
		resp.Base = pack.BuildBaseResp(errno.FollowMyselfError)
		return resp, nil
	}
	switch req.ActionType {
	case 1:
		err = service.NewFollowService(ctx).Follow(req)

	case 0:
		err = service.NewFollowService(ctx).UnFollow(req)
	default:
		err = errno.ParamError
	}
	resp.Base = pack.BuildBaseResp(err)
	return resp, nil
}

// FollowerList implements the FollowHandlerImpl interface.
func (s *FollowHandlerImpl) FollowerList(ctx context.Context, req *follow.FollowerListRequest) (resp *follow.FollowerListResponse, err error) {
	resp = new(follow.FollowerListResponse)
	followerResp, count, err := service.NewFollowService(ctx).FollowerList(req)
	resp.Base = pack.BuildBaseResp(err)
	if err != nil {
		return resp, nil
	}
	resp.Count = &count
	resp.FollowerList = followerResp
	return resp, nil
}

// FollowingList implements the FollowHandlerImpl interface.
func (s *FollowHandlerImpl) FollowingList(ctx context.Context, req *follow.FollowingListRequest) (resp *follow.FollowingListResponse, err error) {
	resp = new(follow.FollowingListResponse)
	followingResp, count, err := service.NewFollowService(ctx).FollowingList(req)
	resp.Base = pack.BuildBaseResp(err)
	if err != nil {
		return resp, nil
	}
	resp.Count = &count
	resp.FollowingList = followingResp
	return resp, nil
}

// FriendList implements the FollowHandlerImpl interface.
func (s *FollowHandlerImpl) FriendList(ctx context.Context, req *follow.FriendListRequest) (resp *follow.FriendListResponse, err error) {
	resp = new(follow.FriendListResponse)
	friendResp, count, err := service.NewFollowService(ctx).FriendList(req)
	resp.Base = pack.BuildBaseResp(err)
	if err != nil {
		return resp, nil
	}
	resp.Count = &count
	resp.FriendList = friendResp
	return resp, nil
}
