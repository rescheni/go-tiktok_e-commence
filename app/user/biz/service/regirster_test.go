package service

import (
	"context"
	"e-commence/app/user/biz/dal/mysql"
	user "e-commence/rpc_gen/kitex_gen/user"
	"testing"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/joho/godotenv"
)

func TestRegirster_Run(t *testing.T) {

	// 加载环境变量
	err := godotenv.Load("../../.env")
	if err != nil {
		klog.Fatal("Error loading .env file")
	}

	mysql.Init()
	if mysql.DB == nil {
		t.Fatal("mysql db is nil")
	}

	ctx := context.Background()
	s := NewRegirsterService(ctx)
	// init req and assert value

	req := &user.RegirsterReq{
		Username:      "testusaaer",
		Password:      "testpasasword",
		Checkpassword: "testpasasword",
		Email:         "1231aa231s@qq.com",
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}

// func TestRegirsterService_Run(t *testing.T) {
// 	type fields struct {
// 		ctx context.Context
// 	}
// 	type args struct {
// 		req *user.RegirsterReq
// 	}
// 	tests := []struct {
// 		name     string
// 		fields   fields
// 		args     args
// 		wantResp *user.RegirsterResp
// 		wantErr  bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			s := &RegirsterService{
// 				ctx: tt.fields.ctx,
// 			}
// 			gotResp, err := s.Run(tt.args.req)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("RegirsterService.Run() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(gotResp, tt.wantResp) {
// 				t.Errorf("RegirsterService.Run() = %v, want %v", gotResp, tt.wantResp)
// 			}
// 		})
// 	}
// }
