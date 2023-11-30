package xctx

import (
	"context"

	"github.com/google/uuid"
	"golang.org/x/text/language"
)

const (
	TraceId = "traceId"
	Locale  = "locale"
	Lang    = "lang"
)

func New() context.Context {
	return Wrap(context.Background())
}

func Wrap(ctx context.Context) context.Context {
	if CtxId(ctx) != "" {
		return ctx
	}
	ctx = context.WithValue(ctx, TraceId, uuid.New().String())
	ctx = context.WithValue(ctx, Locale, "zh")
	ctx = context.WithValue(ctx, Lang, language.SimplifiedChinese)
	return ctx
}

func CtxId(ctx context.Context) string {
	if traceId, ok := ctx.Value(TraceId).(string); ok {
		return traceId
	}
	return ""
}

func CtxLang(ctx context.Context) string {
	if lug, ok := ctx.Value(Lang).(language.Tag); ok {
		return lug.String()
	}
	return ""
}

func CtxLocale(ctx context.Context) string {
	if lc, ok := ctx.Value(Locale).(string); ok {
		return lc
	}
	return ""
}
