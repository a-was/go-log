package log

import (
	"fmt"
	"strings"
)

func formatArgs(args []any) string {
	sb := strings.Builder{}
	for i, v := range args {
		sb.WriteString(fmt.Sprintf("%v", v))
		if len(args)-1 > i {
			sb.WriteString(" ")
		}
	}
	return sb.String()
}

func Debug(v ...any) {
	if log == nil {
		return
	}
	log.Debug(formatArgs(v))
}

func Debugf(format string, v ...any) {
	if log == nil {
		return
	}
	log.Debug(fmt.Sprintf(format, v...))
}

func Info(v ...any) {
	if log == nil {
		return
	}
	log.Info(formatArgs(v))
}

func Infof(format string, v ...any) {
	if log == nil {
		return
	}
	log.Info(fmt.Sprintf(format, v...))
}

func Warn(v ...any) {
	if log == nil {
		return
	}
	log.Warn(formatArgs(v))
}

func Warnf(format string, v ...any) {
	if log == nil {
		return
	}
	log.Warn(fmt.Sprintf(format, v...))
}

func Error(v ...any) {
	if log == nil {
		return
	}
	log.Error(formatArgs(v))
}

func Errorf(format string, v ...any) {
	if log == nil {
		return
	}
	log.Error(fmt.Sprintf(format, v...))
}

func Panic(v ...any) {
	if log == nil {
		return
	}
	log.Panic(formatArgs(v))
}

func Panicf(format string, v ...any) {
	if log == nil {
		return
	}
	log.Panic(fmt.Sprintf(format, v...))
}

func Fatal(v ...any) {
	if log == nil {
		return
	}
	log.Fatal(formatArgs(v))
}

func Fatalf(format string, v ...any) {
	if log == nil {
		return
	}
	log.Fatal(fmt.Sprintf(format, v...))
}
