package logger

import (
	"fmt"
	"github.com/fyf2173/ysdk-go/util"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rs/zerolog"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// RotateType 轮转类型
type RotateType time.Duration

const (
	// RotateWeek 按周轮转
	RotateWeek RotateType = RotateType(time.Hour * 24 * 7)
	// RotateDay 按日轮转
	RotateDay RotateType = RotateType(time.Hour * 24)
)

// NewRotateWriter 轮转日志写入
func NewRotateWriter(name string, rt RotateType) io.Writer {
	var pattern string = name + ".%Y%m%d%H.log"
	var duration time.Duration = time.Duration(rt)
	switch rt {
	case RotateWeek:
		pattern = name + ".%Y%W.log"
	case RotateDay:
		pattern = name + ".%Y%m%d.log"
	}
	fd, err := rotatelogs.New(
		pattern,
		rotatelogs.WithLinkName(name+".log"),
		rotatelogs.WithRotationTime(duration),
		rotatelogs.WithMaxAge(-1),
		rotatelogs.WithLocation(time.Local),
		rotatelogs.WithRotationCount(30),
		rotatelogs.WithLinkName(""),
	)
	util.Assert(err)
	return fd
}

// NewRotateWriterWithHostName 增加主机名称
func NewRotateWriterWithHostName(path string, name string, rt RotateType) io.Writer {
	h, _ := os.Hostname()
	return NewRotateWriter(filepath.Join(path, name+"-"+h), rt)
}

// NewSimpleRotateLogger 新建轮转日志
func NewSimpleRotateLogger(name string, rt RotateType) *log.Logger {
	return log.New(NewRotateWriter(name, rt), "", 0)
}

// NewSimpleRotateLoggerWithHostName 新建轮转日志
func NewSimpleRotateLoggerWithHostName(path string, name string, rt RotateType) *log.Logger {
	return log.New(NewRotateWriterWithHostName(path, name, rt), "", 0)
}

// NewZeroRotateLogger zerolog轮转日志
func NewZeroRotateLogger(name string, rt RotateType) zerolog.Logger {
	return zerolog.New(NewRotateWriter(name, rt)).With().Timestamp().Logger()
}

// NewCurrentLogger 新建本次启动日志
func NewCurrentLogger(name string) *log.Logger {
	fd, err := os.OpenFile(fmt.Sprintf("%s.%d", name, time.Now().Unix()), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	util.Assert(err)
	return log.New(fd, "", 0)
}

// NewRotateLogger 创建轮转日志
func NewRotateLogger(filename string, level zapcore.Level) *zap.Logger {
	hook := lumberjack.Logger{
		Filename:   filename, // 日志文件路径
		MaxBackups: 7,
		MaxAge:     28, //days
		LocalTime:  true,
		Compress:   true, // 是否压缩 disabled by default
	}
	w := zapcore.AddSync(&hook)

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	return zap.New(zapcore.NewCore(zapcore.NewConsoleEncoder(encoderConfig), w, level))
}
