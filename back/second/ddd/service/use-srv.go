package service

import (
	"context"

	log2 "qianliout/leetcode/back/second/ddd/log"
	store2 "qianliout/leetcode/back/second/ddd/store"
)

type UserSrcInterface interface {
	Create(cxt context.Context)
}

type UserSrv struct {
	log log2.Logger
	dal store2.UserDal
}

func (s *UserSrv) Create(cxt context.Context) {
	panic("implement me")
}
func NewUserSrv(log log2.Logger, dal store2.UserDal) *UserSrv {
	return &UserSrv{
		log: log,
		dal: dal,
	}
}
