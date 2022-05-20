package backend

import (
	"context"
)

var (
	wailsContext *context.Context
)

func Start(ctx *context.Context) {
	wailsContext = ctx
}
