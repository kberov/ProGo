package logging

import (
	"log"
	"os"
)

func NewDefaultLogger(level LogLevel) Logger {
	flags := log.Lmsgprefix | log.Ltime
	return &DefaultLogger{
		minLevel: level,
		loggers: map[LogLevel]*log.Logger{
			Trace: log.New(os.Stdout, "TRACE ", flags),
			Debug: log.New(os.Stdout, "DEBUG ", flags),
			Info:  log.New(os.Stdout, "INFO ", flags),
			Warn:  log.New(os.Stdout, "WARN ", flags),
			Fatal: log.New(os.Stdout, "FATAL ", flags),
		},
		triggerPanic: true,
	}
}
