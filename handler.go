package log

import (
	"io"

	"go.uber.org/zap/zapcore"
)

type Handler struct {
	core zapcore.Core
}

type HandlerType int

const (
	HandlerTypeText HandlerType = iota
	HandlerTypeJSON
)

type HandlerConfig struct {
	Type         HandlerType
	Writer       io.Writer
	WriterSynced bool

	Keys             HandlerKeys
	Encoders         HandlerEncoders
	ConsoleSeparator string

	Enabler zapcore.LevelEnabler
}

type HandlerKeys struct {
	Time       string
	Level      string
	Name       string
	Caller     string
	Function   string
	Message    string
	Stacktrace string
}

type HandlerEncoders struct {
	Level    zapcore.LevelEncoder
	Time     zapcore.TimeEncoder
	Duration zapcore.DurationEncoder
	Caller   zapcore.CallerEncoder
	Name     zapcore.NameEncoder // optional
}

func NewHandler(c HandlerConfig) Handler {
	if c.Writer == nil {
		panic("writer cannot be nil")
	}

	c.Keys.defaults()
	c.Encoders.defaults()

	if c.ConsoleSeparator == "" {
		c.ConsoleSeparator = " : "
	}

	if c.Enabler == nil {
		c.Enabler = zapcore.InfoLevel
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:       c.Keys.Time,
		LevelKey:      c.Keys.Level,
		NameKey:       c.Keys.Name,
		CallerKey:     c.Keys.Caller,
		FunctionKey:   c.Keys.Function,
		MessageKey:    c.Keys.Message,
		StacktraceKey: c.Keys.Stacktrace,

		LineEnding:       zapcore.DefaultLineEnding,
		ConsoleSeparator: c.ConsoleSeparator,

		EncodeLevel:    c.Encoders.Level,
		EncodeTime:     c.Encoders.Time,
		EncodeDuration: c.Encoders.Duration,
		EncodeCaller:   c.Encoders.Caller,
		EncodeName:     c.Encoders.Name,
	}

	var encoder zapcore.Encoder
	var writer zapcore.WriteSyncer

	switch c.Type {
	case HandlerTypeJSON:
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	default:
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	if ws, ok := c.Writer.(zapcore.WriteSyncer); ok {
		writer = ws
	} else if c.WriterSynced {
		writer = zapcore.AddSync(c.Writer)
	} else {
		writer = zapcore.Lock(zapcore.AddSync(c.Writer))
	}

	return Handler{
		core: zapcore.NewCore(encoder, writer, c.Enabler),
	}
}

func (k *HandlerKeys) defaults() {
	switch k.Time {
	case "":
		k.Time = "time"
	case "-":
		k.Time = ""
	}

	switch k.Level {
	case "":
		k.Level = "level"
	case "-":
		k.Level = ""
	}

	switch k.Name {
	case "":
		k.Name = "logger"
	case "-":
		k.Name = ""
	}

	switch k.Message {
	case "":
		k.Message = "message"
	case "-":
		k.Message = ""
	}

	switch k.Function {
	case "-":
		k.Function = ""
	}

	switch k.Caller {
	case "":
		k.Caller = "caller"
	case "-":
		k.Caller = ""
	}

	switch k.Stacktrace {
	case "":
		k.Stacktrace = "stacktrace"
	case "-":
		k.Stacktrace = ""
	}
}

func (e *HandlerEncoders) defaults() {
	if e.Level == nil {
		e.Level = zapcore.CapitalLevelEncoder
	}
	if e.Time == nil {
		e.Time = zapcore.ISO8601TimeEncoder
	}
	if e.Duration == nil {
		e.Duration = zapcore.StringDurationEncoder
	}
	if e.Caller == nil {
		e.Caller = zapcore.ShortCallerEncoder
	}
}
