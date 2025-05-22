package service

import (
	"context"
	"e-commence/rpc_gen/kitex_gen/product"
	common "gomall/hertz_gen/frontend/common"
	"gomall/infra/rpc"

	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/hertz-contrib/sessions"
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

	p, err := rpc.ProductClient.ListProduct(h.Context, &product.ListProductReq{})
	if err != nil {
		return nil, err
	}
	var username string
	username = "NULL"
	// 获取用户名
	session := sessions.Default(h.RequestContext)
	if session.Get("user_id") != "" {
		username, _ = session.Get("user_name").(string)
	}
	for _, v := range p.Products {
		v.Picture = "https://api.paugram.com/wallpaper/"
	}

	return utils.H{
		"Icon":       "https://api.paugram.com/wallpaper/",
		"Title":      "首页",
		"Name":       username,
		"Items":      p.Products,
		"footerTime": time.Now().Format("2006"),
	}, nil
}
