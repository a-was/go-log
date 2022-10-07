package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *logger

type logger struct {
	*zap.SugaredLogger
	handlers map[string]*Handler
	options  []zap.Option
}

func (log *logger) build() {
	var cores []zapcore.Core
	for _, v := range log.handlers {
		cores = append(cores, v.core)
	}
	log.SugaredLogger = zap.New(zapcore.NewTee(cores...), log.options...).Sugar()
}

// RegisterHandler allows to register new logging handler.
func RegisterHandler(name string, h *Handler) {
	if h.core == nil {
		return
	}
	if log == nil {
		log = &logger{}
	}
	if log.handlers == nil {
		log.handlers = map[string]*Handler{}
	}
	log.handlers[name] = h
	log.build()
}

// UnregisterHandler allows to unregister existing logging handler.
func UnregisterHandler(name string) {
	if log == nil {
		return
	}
	delete(log.handlers, name)
	log.build()
}

func UseOptions(opts ...zap.Option) {
	log.options = append(log.options, opts...)
	log.SugaredLogger = log.WithOptions(opts...)
}

// WithCaller configures the logger to annotate each message with the filename,
// line number, and function name (it needs to be enabled by setting function key).
func WithCaller() {
	UseOptions(zap.AddCaller(), zap.AddCallerSkip(2))
}

// WithStacktrace configures the logger to record a stack trace for all messages at
// or above Error level.
func WithStacktrace() {
	UseOptions(zap.AddStacktrace(zap.ErrorLevel))
}

// DevelopmentMode puts the logger in development mode, which makes DPanic-level
// logs panic instead of simply logging an error.
func DevelopmentMode() {
	UseOptions(zap.Development())
}
