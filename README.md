# Another Go logging library

# Usage

## Simple

You can use default console handler and default log file handler

```go
package main

import "github.com/a-was/log"

func main() {
	log.UseDefaultConsoleHandler()
	log.UseDefaultFileHandler("app.log")

	log.Debug("Some debug message", 123)
	log.Debugf("Some debug message %d", 123)
	log.Info("Some info message", 123)
	log.Infof("Some info message %d", 123)
	log.Warn("Some warn message", 123)
	log.Warnf("Some warn message %d", 123)
	log.Error("Some error message", 123)
	log.Errorf("Some error message %d", 123)
	log.Panic("Some panic message", 123)
	log.Panicf("Some panic message %d", 123)
	log.Fatal("Some fatal message", 123)
	log.Fatalf("Some fatal message %d", 123)
}
```
