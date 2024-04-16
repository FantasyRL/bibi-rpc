package service

import "bibi/rpc/user/dal/db"

func (s *UserService) GetUserByIdList(uidList []int64) (*[]db.User, error) {
	return db.QueryUserByIDList(s.ctx, uidList)
}
