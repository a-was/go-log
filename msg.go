package log

import (
	"fmt"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type V map[string]any

func formatArgs(args []any) string {
	sb := strings.Builder{}
	for i, v := range args {
		sb.WriteString(fmt.Sprintf("%+v", v))
		if len(args)-1 > i {
			sb.WriteByte(' ')
		}
	}
	return sb.String()
}

func Log(level zapcore.Level, v []any) {
	if log == nil {
		return
	}
	args := formatArgs(v)
	switch level {
	case zap.DebugLevel:
		log.Debug(args)
	case zap.InfoLevel:
		log.Info(args)
	case zap.WarnLevel:
		log.Warn(args)
	case zap.ErrorLevel:
		log.Error(args)
	case zap.DPanicLevel:
		log.DPanic(args)
	case zap.PanicLevel:
		log.Panic(args)
	case zap.FatalLevel:
		log.Fatal(args)
	}
}

func Logf(level zapcore.Level, format string, v []any) {
	if log == nil {
		return
	}
	switch level {
	case zap.DebugLevel:
		log.Debugf(format, v...)
	case zap.InfoLevel:
		log.Infof(format, v...)
	case zap.WarnLevel:
		log.Warnf(format, v...)
	case zap.ErrorLevel:
		log.Errorf(format, v...)
	case zap.DPanicLevel:
		log.DPanicf(format, v...)
	case zap.PanicLevel:
		log.Panicf(format, v...)
	case zap.FatalLevel:
		log.Fatalf(format, v...)
	}
}

func Logw(level zapcore.Level, msg string, vars V) {
	if log == nil {
		return
	}
	var args []any
	for k, v := range vars {
		args = append(args, k, v)
	}
	switch level {
	case zap.DebugLevel:
		log.Debugw(msg, args...)
	case zap.InfoLevel:
		log.Infow(msg, args...)
	case zap.WarnLevel:
		log.Warnw(msg, args...)
	case zap.ErrorLevel:
		log.Errorw(msg, args...)
	case zap.DPanicLevel:
		log.DPanicw(msg, args...)
	case zap.PanicLevel:
		log.Panicw(msg, args...)
	case zap.FatalLevel:
		log.Fatalw(msg, args...)
	}
}

// DEBUG

func Debug(v ...any) {
	Log(zap.DebugLevel, v)
}

func Debugf(format string, v ...any) {
	Logf(zap.DebugLevel, format, v)
}

func Debugw(msg string, vars V) {
	Logw(zap.DebugLevel, msg, vars)
}

// INFO

func Info(v ...any) {
	Log(zap.InfoLevel, v)
}

func Infof(format string, v ...any) {
	Logf(zap.InfoLevel, format, v)
}

func Infow(msg string, vars V) {
	Logw(zap.InfoLevel, msg, vars)
}

// WARN

func Warn(v ...any) {
	Log(zap.WarnLevel, v)
}

func Warnf(format string, v ...any) {
	Logf(zap.WarnLevel, format, v)
}

func Warnw(msg string, vars V) {
	Logw(zap.WarnLevel, msg, vars)
}

// ERROR

func Error(v ...any) {
	Log(zap.ErrorLevel, v)
}

func Errorf(format string, v ...any) {
	Logf(zap.ErrorLevel, format, v)
}

func Errorw(msg string, vars V) {
	Logw(zap.ErrorLevel, msg, vars)
}

// DPANIC

func DPanic(v ...any) {
	Log(zap.DPanicLevel, v)
}

func DPanicf(format string, v ...any) {
	Logf(zap.DPanicLevel, format, v)
}

func DPanicw(msg string, vars V) {
	Logw(zap.DPanicLevel, msg, vars)
}

// PANIC

func Panic(v ...any) {
	Log(zap.PanicLevel, v)
}

func Panicf(format string, v ...any) {
	Logf(zap.PanicLevel, format, v)
}

func Panicw(msg string, vars V) {
	Logw(zap.PanicLevel, msg, vars)
}

// FATAL

func Fatal(v ...any) {
	Log(zap.FatalLevel, v)
}

func Fatalf(format string, v ...any) {
	Logf(zap.FatalLevel, format, v)
}

func Fatalw(msg string, vars V) {
	Logw(zap.FatalLevel, msg, vars)
}
