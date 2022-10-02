package log

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	DefaultConsoleStdoutHandlerName = "default-console-stdout-handler"
	DefaultConsoleStderrHandlerName = "default-console-stderr-handler"
	DefaultFileHandlerName          = "default-file-handler"
)

func UseDefaultConsoleHandler() {
	RegisterHandler(DefaultConsoleStdoutHandlerName, NewHandler(HandlerConfig{
		Type:         HandlerTypeText,
		Writer:       os.Stdout,
		WriterSynced: false,
		Encoders: HandlerEncoders{
			Time: zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.999"),
		},
		Enabler: zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl < zapcore.ErrorLevel
		}),
	}))

	RegisterHandler(DefaultConsoleStderrHandlerName, NewHandler(HandlerConfig{
		Type:         HandlerTypeText,
		Writer:       os.Stderr,
		WriterSynced: false,
		Encoders: HandlerEncoders{
			Time: zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.999"),
		},
		Enabler: zapcore.ErrorLevel,
	}))
}

func UseDefaultFileHandler(path string) {
	file, _, err := zap.Open(path)
	if err != nil {
		panic(err)
	}

	RegisterHandler(DefaultFileHandlerName, NewHandler(HandlerConfig{
		Type:         HandlerTypeText,
		Writer:       file,
		WriterSynced: true,
		Enabler:      zap.DebugLevel,
	}))
}
