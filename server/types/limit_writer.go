package types

import (
	"context"
	"io"

	"golang.org/x/time/rate"
)

type RateLimitWriter struct {
	writer  io.Writer
	limit   int
	limiter *rate.Limiter
	context context.Context
	sktInfo *SocketInfo
}
