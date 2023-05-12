package main

import (
	"context"

	"github.com/CloudWeOps/psychic/cmd/user/pkg/mysql"
	"github.com/CloudWeOps/psychic/share/errno"
	user "github.com/CloudWeOps/psychic/share/kitex_gen/user"
	"github.com/CloudWeOps/psychic/share/tools"
	"github.com/cloudwego/kitex/pkg/klog"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct {
	UserMysqlManager
}

type UserMysqlManager interface {
	CreateUser(user *mysql.User) (*mysql.User, error)
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.LoginRequest) (resp *user.LoginResponse, err error) {
	// TODO: Your code here...
	return
}

// AdminLogin implements the UserServiceImpl interface.
func (s *UserServiceImpl) AdminLogin(ctx context.Context, req *user.AdminLoginRequest) (resp *user.AdminLoginResponse, err error) {
	// TODO: Your code here...
	return
}

// ChangeAdminPassword implements the UserServiceImpl interface.
func (s *UserServiceImpl) ChangeAdminPassword(ctx context.Context, req *user.ChangeAdminPasswordRequest) (resp *user.ChangeAdminPasswordResponse, err error) {
	// TODO: Your code here...
	return
}

// UploadAvatar implements the UserServiceImpl interface.
func (s *UserServiceImpl) UploadAvatar(ctx context.Context, req *user.UploadAvatarRequset) (resp *user.UploadAvatarResponse, err error) {
	// TODO: Your code here...
	return
}

// GetUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUser(ctx context.Context, req *user.GetUserRequest) (resp *user.GetUserInfoResponse, err error) {
	// TODO: Your code here...
	return
}

// AddUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) AddUser(ctx context.Context, req *user.AddUserRequest) (resp *user.AddUserResponse, err error) {
	resp = new(user.AddUserResponse)
	_, err = s.UserMysqlManager.CreateUser(&mysql.User{
		ID:           req.AccountId,
		PhoneNumber:  req.PhoneNumber,
		AvatarBlobId: req.AvatarBlobId,
		Username:     req.Username,
		OpenID:       req.OpenId,
	})
	if err != nil {
		if err == errno.RecordAlreadyExist {
			klog.Error("add user error", err)
			resp.BaseResp = tools.BuildBaseResp(errno.RecordAlreadyExist)
			return resp, nil
		}
		klog.Error("add user error", err)
		resp.BaseResp = tools.BuildBaseResp(errno.UserSrvErr)
		return resp, nil
	}
	resp.BaseResp = tools.BuildBaseResp(nil)
	return resp, nil

}

// DeleteUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) DeleteUser(ctx context.Context, req *user.DeleteUserRequest) (resp *user.DeleteUserResponse, err error) {
	// TODO: Your code here...
	return
}

// UpdateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdateUser(ctx context.Context, req *user.UpdateUserRequest) (resp *user.UpdateUserResponse, err error) {
	// TODO: Your code here...
	return
}

// Pay implements the UserServiceImpl interface.
func (s *UserServiceImpl) Pay(ctx context.Context, req *user.PayRequest) (resp *user.PayResponse, err error) {
	// TODO: Your code here...
	return
}

// GetSomeUsers implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetSomeUsers(ctx context.Context, req *user.GetSomeUsersRequest) (resp *user.GetSomeUsersResponse, err error) {
	// TODO: Your code here...
	return
}

// GetAllUsers implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetAllUsers(ctx context.Context, req *user.GetAllUsersRequest) (resp *user.GetAllUsersResponse, err error) {
	// TODO: Your code here...
	return
}
