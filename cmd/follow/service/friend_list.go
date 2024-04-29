package service

import (
	"bibi/cmd/follow/dal/db"
	"bibi/cmd/follow/rpc"
	"bibi/kitex_gen/base"
	"bibi/kitex_gen/follow"
	"bibi/kitex_gen/user"
	"bibi/pkg/errno"
)

func (s *FollowService) FriendList(req *follow.FriendListRequest) ([]*base.User, int64, error) {
	_, friendList, _, err := db.FollowerAndFriendList(s.ctx, req.UserId)
	if err != nil {
		return nil, 0, err
	}
	count := int64(len(*friendList))
	friendIdList := make([]int64, len(*friendList))
	for i, v := range *friendList {
		friendIdList[i] = v.FollowedId
	}

	rpcResp, err := rpc.UserGetUserList(s.ctx, &user.GetUsersRequest{
		UserIdList: friendIdList,
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

func FriendCount(uid int64) {

}
