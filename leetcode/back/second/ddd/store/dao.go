package store

import (
	log2 "qianliout/leetcode/back/second/ddd/log"

	"gorm.io/gorm"
)

type UserDal interface {
	Create()
}

type UserDao struct {
	log log2.Logger
	db  *gorm.DB
}

func (s *UserDao) Create() {
	panic("implement me")
}

func NewUserDao(logger log2.Logger, db *gorm.DB) *UserDao {
	return &UserDao{log: logger, db: db}
}
