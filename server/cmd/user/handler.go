package main

import (
	"Herbal/server/shared/kitex_gen/base"
	"Herbal/server/shared/kitex_gen/user"
	"context"
)

type MysqlManager interface {
	CreateUser(username, role string) error
}

type RedisManager interface {
}

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct {
	MysqlManager
	RedisManager
}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	resp = &user.RegisterResp{
		Password:        "22",
		Role:            "33",
		ConfirmPassword: "22",
		BaseResp: &base.BaseResponse{
			StatusCode: 0,
			StatusMsg:  "sss",
		},
	}
	return resp, nil
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.LoginReq) (resp *user.LoginResp, err error) {
	// TODO: Your code here...
	return
}

// ChangePassword implements the UserServiceImpl interface.
func (s *UserServiceImpl) ChangePassword(ctx context.Context, req *user.ChangePasswordReq) (resp *user.ChangePasswordResp, err error) {
	// TODO: Your code here...
	return
}
