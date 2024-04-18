package service

import "bibi/cmd/user/dal/db"

func (s *UserService) GetUserByIdList(uidList []int64) (*[]db.User, error) {
	return db.QueryUserByIDList(s.ctx, uidList)
}
