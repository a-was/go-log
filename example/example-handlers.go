package example

import (
	"os"

	"github.com/a-was/go-log"
	"github.com/a-was/go-log/handlers"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// How to use zap:
// https://pkg.go.dev/go.uber.org/zap#example-package-AdvancedConfiguration

// this example shows how to use standard handlers
func ExampleH1() {
	log.RegisterHandler("console-stdout", log.NewHandler(log.HandlerConfig{
		Type:         log.HandlerTypeText,
		Writer:       os.Stdout,
		WriterSynced: false,
		Enabler: zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl < zapcore.ErrorLevel
		}),
	}))
	log.RegisterHandler("console-stderr", log.NewHandler(log.HandlerConfig{
		Type:         log.HandlerTypeText,
		Writer:       os.Stderr,
		WriterSynced: false,
		Enabler:      zap.ErrorLevel,
	}))

	file, close, err := zap.Open("app.log")
	if err != nil {
		panic(err)
	}
	defer close()

	log.RegisterHandler("file", log.NewHandler(log.HandlerConfig{
		Type:         log.HandlerTypeJSON,
		Writer:       file,
		WriterSynced: true,
		Enabler:      zap.InfoLevel,
	}))
}

// this example shows how to register a SMTP handler
func ExampleH2() {
	log.RegisterHandler("smtp", handlers.SMTPHandler(handlers.SMTPConfig{
		From:    "Error sender <error-sender@myserver.com>",
		To:      []string{"me@example.com", "you@example.com"},
		Subject: "My app error!",
		Server:  "localhost:25",
		Auth:    nil,
		Enabler: zap.WarnLevel,
	}))
}

type mywriter struct{}

func (mywriter) Write(p []byte) (int, error) {
	// ...
	return len(p), nil
}

// this example shows how to use custom writer
func Example3() {
	log.RegisterHandler("custom-handler-JSON", log.NewHandler(log.HandlerConfig{
		Type:         log.HandlerTypeJSON,
		Writer:       mywriter{},
		WriterSynced: true, // if unsure set this to false
		Enabler:      zapcore.DebugLevel,
	}))
}
