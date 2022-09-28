package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *logger

type logger struct {
	*zap.Logger
	handlers map[string]Handler
}

func (log *logger) build() {
	var cores []zapcore.Core
	for _, v := range log.handlers {
		cores = append(cores, v.core)
	}
	log.Logger = zap.New(zapcore.NewTee(cores...))
}
