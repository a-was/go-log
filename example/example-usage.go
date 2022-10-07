package example

import (
	"os"

	"github.com/a-was/go-log"
)

type myStruct struct {
	Name            string
	priv            string
	NotStandardType chan int `json:"-"` // this needs to be disabled with json tag
}

func ExampleU1() {
	x := myStruct{
		Name:            "example",
		priv:            "private",
		NotStandardType: make(chan int),
	}

	log.UseDefaultConsoleHandler()

	log.DevelopmentMode() // makes DPanic level panics

	// default format for args is %+v
	log.Info("Some info message with arg:", 123, "and then struct:", x)
	// 2022-01-30 07:38:24.521 : INFO : Some info message with arg: 123 and then struct: {Name:example priv:private NotStandardType:0xc0000960c0}

	// you can use fmt formats
	log.Infof("Some info message with arg: %d and then struct %#v", 123, x)
	// 2022-01-30 07:38:24.521 : INFO : Some info message with arg: 123 and then struct main.myStruct{Name:"example", priv:"private", NotStandardType:(chan int)(0xc0000960c0)}

	// log with custom vars. Custom struct types needs to be disabled with json tag
	log.Infow("Some info message with vars", log.V{"int": 123, "struct": x})
	// 2022-01-30 07:38:24.521 : INFO : Some info message with vars : {"int": 123, "struct": {"Name":"example"}}

	// you can disable default handlers
	log.UnregisterHandler(log.DefaultConsoleStdoutHandlerName)
	log.UnregisterHandler(log.DefaultConsoleStderrHandlerName)

	log.RegisterHandler("json-handler", log.NewHandler(log.HandlerConfig{
		Type:   log.HandlerTypeJSON,
		Writer: os.Stdout,
	}))

	log.Info("Some info message with arg:", 123, "and then struct:", x)
	// {"level":"INFO","time":"2022-01-30T07:38:24.521+0200","message":"Some info message with arg: 123 and then struct: {Name:example priv:private NotStandardType:0xc0000960c0}"}
	log.Infof("Some info message with arg: %d and then struct %#v", 123, x)
	// {"level":"INFO","time":"2022-01-30T07:38:24.521+0200","message":"Some info message with arg: 123 and then struct example.myStruct{Name:\"example\", priv:\"private\", NotStandardType:(chan int)(0xc0000960c0)}"}
	log.Infow("Some info message with vars", log.V{"int": 123, "struct": x})
	// {"level":"INFO","time":"2022-01-30T07:38:24.521+0200","message":"Some info message with vars","int":123,"struct":{"Name":"example"}}
}

func ExampleU2() {
	log.RegisterHandler("console", log.NewHandler(log.HandlerConfig{
		Writer: os.Stdout,
		Keys: log.HandlerKeys{
			Function: "func",
		},
	}))
	log.RegisterHandler("console-json", log.NewHandler(log.HandlerConfig{
		Type:   log.HandlerTypeJSON,
		Writer: os.Stdout,
		Keys: log.HandlerKeys{
			Function: "func",
		},
	}))

	log.WithCaller()
	log.WithStacktrace()

	x := myStruct{
		Name:            "example",
		priv:            "private",
		NotStandardType: make(chan int),
	}

	log.Infow("Some info message with vars", log.V{"int": 123, "struct": x})
	// {"level":"INFO","time":"2022-01-30T07:38:24.521+0200","caller":"example/example-usage.go:73","func":"github.com/a-was/go-log/example.ExampleU2","message":"Some info message with vars","int":123,"struct":{"Name":"example"}}
	// 2022-01-30T07:38:24.521+0200 : INFO : example/example-usage.go:73 : Some info message with vars : {"int": 123, "struct": {"Name":"example"}}
}
