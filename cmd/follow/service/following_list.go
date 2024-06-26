package service

import (
	"bibi/cmd/follow/dal/db"
	"bibi/cmd/follow/rpc"
	"bibi/kitex_gen/base"
	"bibi/kitex_gen/follow"
	"bibi/kitex_gen/user"
	"bibi/pkg/errno"
)

func (s *FollowService) FollowingList(req *follow.FollowingListRequest) ([]*base.User, int64, error) {
	followingList, count, err := db.FollowingList(s.ctx, req.UserId)
	if err != nil {
		return nil, 0, err
	}
	//if err = cache.SetFollowingList(s.ctx, req.UserId, followingList); err != nil {
	//	return nil, 0, err
	//}
	//if err = cache.SetFollowingCount(s.ctx, req.UserId, count); err != nil {
	//	return nil, 0, err
	//}

	followerIdList := make([]int64, len(*followingList))
	for i, v := range *followingList {
		followerIdList[i] = v.Uid
	}

	rpcResp, err := rpc.UserGetUserList(s.ctx, &user.GetUsersRequest{
		UserIdList: followerIdList,
	})
	if rpcResp.Base.Code != errno.SuccessCode {
		return nil, 0, errno.NewErrNo(rpcResp.Base.Code, rpcResp.Base.Msg)
	}
	if err != nil {
		return nil, 0, err
	}
	for i := range rpcResp.UserList {
		rpcResp.UserList[i].IsFollow = true
	}
	return rpcResp.UserList, count, nil
}

func FollowedCount(uid int64) {

}
