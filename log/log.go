package log

import (
	"fmt"
	"log"
)

func write(typ string, msg string, args ...interface{}) {
	line := fmt.Sprint(typ, msg)
	for i := 0; i < len(args)/2; i++ {
		line = fmt.Sprintf("%s %s=%v", line, args[2*i], args[2*i+1])
	}
	log.Println(line)
}

func Info(msg string, args ...interface{}) {
	write("INFO", msg, args)
}

func Warn(msg string, args ...interface{}) {
	write("WARN", msg, args)
}

func Debug(msg string, args ...interface{}) {
	write("DEBUG", msg, args)
}

func Error(msg string, args ...interface{}) {
	write("ERROR", msg, args)
}
