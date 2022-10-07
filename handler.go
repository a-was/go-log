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
	// Type allows to choose how messages should be encoded.
	//
	// Default: HandlerTypeText
	Type HandlerType
	// Writer defines where to write messages.
	Writer io.Writer
	// WriterSynced defines if writer is already synced.
	// If false, it is wrapped with mutex to make it safe for concurrent use.
	// *os.Files must have this set to false.
	WriterSynced bool

	// Define how to name logger keys.
	// Empty string is treated as default value.
	// If You want to diable some key, set its value to "-".
	Keys HandlerKeys
	// Define how to encode logger fields.
	// nil value falls back to default value.
	// if You want to disable some encoder, You have to disable its key.
	Encoders HandlerEncoders
	// ConsoleSeparator between line elements. Default " : " (TIME : LEVEL : MESSAGE).
	ConsoleSeparator string

	// When to log message. Most of the times it simply will be log level. Default to info level.
	Enabler zapcore.LevelEnabler
}

type HandlerKeys struct {
	Time       string // Default: time
	Level      string // Default: level
	Name       string // Default: logger
	Caller     string // Default: caller
	Function   string // Default: - (disabled)
	Message    string // Default: message
	Stacktrace string // Default: stacktrace
}

type HandlerEncoders struct {
	Level    zapcore.LevelEncoder    // Default: zapcore.CapitalLevelEncoder
	Time     zapcore.TimeEncoder     // Default: zapcore.ISO8601TimeEncoder
	Duration zapcore.DurationEncoder // Default: zapcore.StringDurationEncoder
	Caller   zapcore.CallerEncoder   // Default: zapcore.ShortCallerEncoder
	Name     zapcore.NameEncoder     // Optional
}

func NewHandler(c HandlerConfig) *Handler {
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

	return &Handler{
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
