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

    log.Info("Some info message", 123)
}
```
