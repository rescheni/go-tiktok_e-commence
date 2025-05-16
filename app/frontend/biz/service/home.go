package service

import (
	"context"
	"time"

	home "gomall/hertz_gen/frontend/home"

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
func (h *HomeService) Run(req *home.Empty) (resp map[string]any, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code

	resp = make(map[string]any)

	items := []map[string]any{
		{"Name": "黄金", "Price": 1000, "Picture": "http://image.reschen.cn:88/random/pa"},
		{"Name": "石油", "Price": 100, "Picture": "http://image.reschen.cn:88/random/pb"},
		{"Name": "白银", "Price": 80, "Picture": "http://image.reschen.cn:88/random/pc"},
		{"Name": "铂金", "Price": 200, "Picture": "http://image.reschen.cn:88/random/pd"},
		{"Name": "奢侈品", "Price": 9999, "Picture": "http://image.reschen.cn:88/random/pe"},
		{"Name": "汽车", "Price": 9100000, "Picture": "http://image.reschen.cn:88/random/pf"},
		{"Name": "房子", "Price": 10000, "Picture": "http://image.reschen.cn:88/random/pg"},
		{"Name": "飞机", "Price": 100000, "Picture": "http://image.reschen.cn:88/random/ph"},
		{"Name": "游艇", "Price": 1000000, "Picture": "http://image.reschen.cn:88/random/pi"},
	}
	resp["Items"] = items
	resp["Title"] = "首页"
	resp["Icon"] = "http://image.reschen.cn:88/random/pa"
	resp["footerTime"] = time.Now().Format("2006")

	return
}
