package main

import (
	"Herbal/shared/kitex_gen/user"
	"context"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	// TODO: Your code here...
	return
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
