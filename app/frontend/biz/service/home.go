package service

import (
	"context"
	common "gomall/hertz_gen/frontend/common"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
)

type HomeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewHomeService(Context context.Context, RequestContext *app.RequestContext) *HomeService {
	return &HomeService{RequestContext: RequestContext, Context: Context}
}

// func (h *HomeService) Run(req *home.Empty) (resp *home.Empty, err error) {
func (h *HomeService) Run(req *common.Empty) (resp map[string]any, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code

	resp = make(map[string]any)

	items := []map[string]any{
		{"Name": "黄金", "Price": 1000, "Picture": "https://api.paugram.com/wallpaper/"},
		{"Name": "石油", "Price": 100, "Picture": "https://api.paugram.com/wallpaper/"},
		{"Name": "白银", "Price": 80, "Picture": "https://api.paugram.com/wallpaper/"},
		{"Name": "铂金", "Price": 200, "Picture": "https://api.paugram.com/wallpaper/"},
		{"Name": "奢侈品", "Price": 9999, "Picture": "https://api.paugram.com/wallpaper/"},
		{"Name": "汽车", "Price": 9100000, "Picture": "https://api.paugram.com/wallpaper/"},
		{"Name": "房子", "Price": 10000, "Picture": "https://api.paugram.com/wallpaper/"},
		{"Name": "飞机", "Price": 100000, "Picture": "https://api.paugram.com/wallpaper/"},
		{"Name": "游艇", "Price": 1000000, "Picture": "https://api.paugram.com/wallpaper/"},
	}
	resp["Items"] = items
	resp["Title"] = "首页"
	resp["Icon"] = "https://api.paugram.com/wallpaper/"
	resp["footerTime"] = time.Now().Format("2006")

	return
}
