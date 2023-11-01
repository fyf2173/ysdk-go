package xlog

import (
	"context"
	"log/slog"

	"github.com/fyf2173/ysdk-go/xctx"
)

func Debug(ctx context.Context, msg string, args ...slog.Attr) {
	args = append(args, slog.String("trace_id", xctx.CtxId(ctx)))
	slog.LogAttrs(ctx, slog.LevelDebug, msg, args...)
}

func Info(ctx context.Context, msg string, args ...slog.Attr) {
	args = append(args, slog.String("trace_id", xctx.CtxId(ctx)))
	slog.LogAttrs(ctx, slog.LevelInfo, msg, args...)
}

func Warn(ctx context.Context, msg string, args ...slog.Attr) {
	args = append(args, slog.String("trace_id", xctx.CtxId(ctx)))
	slog.LogAttrs(ctx, slog.LevelWarn, msg, args...)
}

func Error(ctx context.Context, msg string, args ...slog.Attr) {
	args = append(args, slog.String("trace_id", xctx.CtxId(ctx)))
	slog.LogAttrs(ctx, slog.LevelError, msg, args...)
}
