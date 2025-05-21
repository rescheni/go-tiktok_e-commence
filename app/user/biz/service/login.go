package service

import (
	"context"
	"e-commence/app/user/biz/dal/mysql"
	"e-commence/app/user/model"
	user "e-commence/rpc_gen/kitex_gen/user"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type LoginService struct {
	ctx context.Context
} // NewLoginService new LoginService
func NewLoginService(ctx context.Context) *LoginService {
	return &LoginService{ctx: ctx}
}

// Run create note info
func (s *LoginService) Run(req *user.LoginReq) (resp *user.LoginResp, err error) {
	// Finish your business logic.

	if req.Username == "" || req.Password == "" {
		return nil, errors.New("username or password is empty")
	}
	logininfo := req.Username
	u, err := checkUser(logininfo)
	if err != nil {
		return nil, errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.PasswordHashed), []byte(req.Password))
	if err != nil {
		return nil, errors.New("password not match")
	}
	resp = &user.LoginResp{
		UserId:   int64(u.ID),
		Username: u.Username,
		Token:    "",
	}

	return
}

func checkUser(logininfo string) (*model.User, error) {

	u, err := model.GetUserByUsername(mysql.DB, logininfo)

	if u == nil {
		u, err = model.GetUserByEmail(mysql.DB, logininfo)
	}
	if err != nil {
		return nil, err
	}
	return u, nil
}
