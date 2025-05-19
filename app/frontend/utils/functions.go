package utils

import (
	"context"
)

func GetUserIdFromCtx(ctx context.Context) int64 {
	userid := ctx.Value(SessionUserId)
	if userid == nil {
		return 0
	}
	return userid.(int64)
}
