package log

import (
	"fmt"
	"log"
	"os"
)

type Logger struct {
	writer *log.Logger
	module string
}

func New(module string) Logger {
	return Logger{
		writer: log.New(os.Stdout, "", log.LstdFlags|log.Lmicroseconds),
		module: module,
	}
}

func (l *Logger) write(typ string, msg string, args ...interface{}) {
	line := fmt.Sprintf("%s %s", typ, msg)
	for i := 0; i < len(args)/2; i++ {
		line = fmt.Sprintf("%s modedule=%s %s=%v", line, l.module, args[2*i], args[2*i+1])
	}
	l.writer.Println(line)
}

func (l *Logger) Info(msg string, args ...interface{}) {
	l.write("INFO", msg, args...)
}

func (l *Logger) Error(msg string, args ...interface{}) {
	l.write("ERROR", msg, args...)
}
