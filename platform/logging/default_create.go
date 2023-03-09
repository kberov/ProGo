package logging

import (
	"log"
	"os"
	"platform/config"
	"strings"
)

func NewDefaultLogger(cfg config.Configuration) Logger {
	var level LogLevel = Debug
	if configLevel, found := cfg.GetString("logging:level"); found {
		level = LogLevelFromString(configLevel)
	}
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

func LogLevelFromString(val string) (level LogLevel) {
	switch strings.ToLower(val) {
	case "debug":
		level = Debug
	case "info":
		level = Info
	case "warn":
		level = Warn
	case "fatal":
		level = Fatal
	case "none":
		// 868 Chapter 32 ■ Creating a Web Platform
		level = None
	default:
		level = Debug
	}
	return
}
