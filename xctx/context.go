package xctx

import (
	"context"

	"github.com/google/uuid"
	"golang.org/x/text/language"
)

const (
	traceId = "traceId"
	locale  = "locale"
	lang    = "lang"
)

func New() context.Context {
	return Wrap(context.Background())
}

func Wrap(ctx context.Context) context.Context {
	if CtxId(ctx) != "" {
		return ctx
	}
	ctx = context.WithValue(ctx, traceId, uuid.New().String())
	ctx = context.WithValue(ctx, locale, "zh")
	ctx = context.WithValue(ctx, lang, language.SimplifiedChinese)
	return ctx
}

func CtxId(ctx context.Context) string {
	if traceId, ok := ctx.Value(traceId).(string); ok {
		return traceId
	}
	return ""
}

func CtxLang(ctx context.Context) string {
	if lug, ok := ctx.Value(lang).(language.Tag); ok {
		return lug.String()
	}
	return ""
}

func CtxLocale(ctx context.Context) string {
	if lc, ok := ctx.Value(locale).(string); ok {
		return lc
	}
	return ""
}
