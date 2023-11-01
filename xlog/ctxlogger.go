package xlog

import (
	"context"
	"log/slog"
	"os"

	"github.com/fyf2173/ysdk-go/xctx"
)

var levelMap = map[string]slog.Level{
	"debug": slog.LevelDebug,
	"info":  slog.LevelInfo,
	"warn":  slog.LevelWarn,
	"error": slog.LevelError,
}

func Init(level string, addSource bool) {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
		Level:     levelMap[level],
		AddSource: addSource,
	})))
}

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

func Error(ctx context.Context, err error, args ...slog.Attr) {
	if err == nil {
		return
	}
	args = append(args, slog.String("trace_id", xctx.CtxId(ctx)))
	slog.LogAttrs(ctx, slog.LevelError, err.Error(), args...)
}
