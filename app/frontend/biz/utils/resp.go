package utils

import (
	"context"
	"gomall/middleware"

	"github.com/cloudwego/hertz/pkg/app"
)

// SendErrResponse  pack error response
func SendErrResponse(ctx context.Context, c *app.RequestContext, code int, err error) {
	// todo edit custom code
	c.String(code, err.Error())
}

// SendSuccessResponse  pack success response
func SendSuccessResponse(ctx context.Context, c *app.RequestContext, code int, data interface{}) {
	// todo edit custom code
	c.JSON(code, data)
}

func WarpRespose(ctx context.Context, c *app.RequestContext, content map[string]any) map[string]any {
	// todo edit custom code

	userId := ctx.Value(middleware.SessionUserId)
	content["user_id"] = userId
	return content
}
