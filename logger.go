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

func WithCaller() {
	UseOptions(zap.AddCaller(), zap.AddCallerSkip(1))
}

func WithStacktrace() {
	UseOptions(zap.AddStacktrace(zap.ErrorLevel))
}

func DevelopmentMode() {
	UseOptions(zap.Development())
}
