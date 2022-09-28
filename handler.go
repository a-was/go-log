package log

import "go.uber.org/zap/zapcore"

type Handler struct {
	core zapcore.Core
}

func RegisterHandler(name string, h Handler) {
	if log == nil {
		log = &logger{}
	}
	if log.handlers == nil {
		log.handlers = map[string]Handler{}
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
