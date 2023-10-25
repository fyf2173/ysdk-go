package xctx

import (
	"context"
	"fmt"
	"testing"
)

func TestNewContext(t *testing.T) {
	ctx := New()
	ctx = context.WithValue(ctx, "user", "meme")

	fmt.Println("dddd1--------", ctx.Value("user"), CtxId(ctx), CtxLang(ctx), CtxLocale(ctx))

	// output:
}
