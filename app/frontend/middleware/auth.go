package middleware

import (
	"context"

	frontendUtil "e-commence/app/frontend/utils"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
)

// 认证中间件
func GlobalAuth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// 这里可以添加认证逻辑
		// 例如检查请求头中的 token 是否有效

		//  认证逻辑
		// 从session 中获取用户信息  放到context 中   后续需要用户数据直接在context 中获取
		session := sessions.Default(c)
		session.Get("user_id")
		ctx = context.WithValue(ctx, frontendUtil.SessionUserId, session.Get("user_id"))
		// fmt.Println("session user_id = ", session.Get("user_id"))
		// fmt.Println("🀄️间件")
		c.Next(ctx)

	}
}

// 鉴权中间件
func Auth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		session := sessions.Default(c)
		sessionId := session.Get("user_id")
		if sessionId == nil {
			// 没有获取到用户信息，表示没有登录
			// GOTO /sign-in
			// c.FullPath()
			c.Redirect(302, []byte("/sign-in?next="+c.FullPath()))
			c.Abort()
			return
		}
	}
}
