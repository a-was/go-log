package log

import (
	"testing"
	"time"
)

func Test(t *testing.T) {
	UseDefaultConsoleHandler()
	UseDefaultFileHandler("app.log")

	var x any = []string{"a", "b"}
	Info("My info", time.Duration(1234*time.Millisecond), x)
	Error("My error")

	// UseOptions(zap.AddCaller(), zap.AddCallerSkip(1))
	WithCaller()
	WithStacktrace()

	Info("My info", time.Duration(1234*time.Millisecond), x)
	Error("My error")
}
