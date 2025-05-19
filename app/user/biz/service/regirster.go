package service

import (
	"context"
	"e-commence/app/user/biz/dal/mysql"
	"e-commence/app/user/model"
	user "e-commence/rpc_gen/kitex_gen/user"
	"errors"

	"github.com/cloudwego/kitex/pkg/klog"
	"golang.org/x/crypto/bcrypt"
)

type RegirsterService struct {
	ctx context.Context
} // NewRegirsterService new RegirsterService
func NewRegirsterService(ctx context.Context) *RegirsterService {
	return &RegirsterService{ctx: ctx}
}

// Run create note info
func (s *RegirsterService) Run(req *user.RegirsterReq) (resp *user.RegirsterResp, err error) {

	if req.Email == "" || req.Password == "" || req.Username == "" {
		return nil, errors.New("email or passwd or username is empty")
	}

	if req.Password != req.Checkpassword {
		return nil, errors.New("not match password")
	}
	klog.Infof("req: %v", req)

	passwordHashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	newUser := &model.User{
		Username:       req.Username,
		PasswordHashed: string(passwordHashed),
		Email:          req.Email,
	}

	err = model.Create(mysql.DB, newUser)

	if err != nil {
		return nil, err
	}

	resp = &user.RegirsterResp{
		UserId: int64(newUser.ID),
		Token:  "",
	}

	return
}
