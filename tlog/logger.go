package tlog

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	zapLog   *zap.Logger
	LogMode  string = "dev"
	LogLevel string = "d"
)

func GetLogger() *zap.Logger {
	return zapLog
}

func customProdlogTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("Jan 2 15:04:05.000")) //Jan 2
}

func customDevlogTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("15:04:05.000")) //Jan 2
}

func customCallerEncoder(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("\033[38;5;241m" + caller.TrimmedPath() + ":" + "\033[0m")
}

func init() {
	if LogMode == "dev" {
		initDev()
		return
	}
	initProd()
}

func initProd() {
	c := zap.NewProductionConfig()
	c.OutputPaths = []string{"stdout" /* , "./logs" */}
	pe := zap.NewProductionEncoderConfig()
	pe.EncodeTime = customProdlogTimeEncoder
	pe.EncodeCaller = zapcore.FullCallerEncoder
	c.EncoderConfig = pe
	build(c)
}

func initDev() {
	c := zap.NewDevelopmentConfig()
	c.OutputPaths = []string{"stdout" /* , "./logs" */}
	de := zap.NewDevelopmentEncoderConfig()
	de.EncodeTime = customDevlogTimeEncoder
	// de.EncodeLevel = zapcore.CapitalColorLevelEncoder
	// de.EncodeCaller = customCallerEncoder
	de.ConsoleSeparator = " "
	c.EncoderConfig = de
	build(c)
}

func build(c zap.Config) {
	level := zap.InfoLevel

	if lvl := strings.ToLower(LogLevel); lvl != "" {
		switch lvl[0] {
		case 'e':
			level = zap.ErrorLevel
		case 'w':
			level = zap.WarnLevel
		case 'd':
			level = zap.DebugLevel
		case 'p':
			level = zap.PanicLevel
		case 'f':
			level = zap.FatalLevel
		}
	}

	c.Level.SetLevel(level)
	logger, err := c.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
	zapLog = logger
}

func addSpace(v []any) string {
	if len(v) == 0 {
		return " "
	}
	var out bytes.Buffer
	out.WriteString(fmt.Sprint(v[0]))
	for _, item := range v[1:] {
		out.WriteString(" ")
		out.WriteString(fmt.Sprint(item))
	}
	return out.String()
}

func Info(v ...any) {
	zapLog.Info(addSpace(v))
}

func Debug(v ...any) {
	zapLog.Debug(addSpace(v))
}

func Error(v ...any) {
	zapLog.Error(addSpace(v))
}

func Warning(v ...any) {
	zapLog.Warn(addSpace(v))
}

func Fatal(v ...any) {
	zapLog.Fatal(addSpace(v))
}

func PrintJSON(v ...any) {
	var out bytes.Buffer

	for _, element := range v {
		b, err := json.MarshalIndent(element, "", "  ")
		if err != nil {
			Error(err)
			continue
		}
		out.Write(b)
	}
	zapLog.Info("\n" + out.String())
}
