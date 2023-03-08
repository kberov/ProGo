package logging

import (
	. "fmt"
	"log"
)

type DefaultLogger struct {
	minLevel     LogLevel
	loggers      map[LogLevel]*log.Logger
	triggerPanic bool
}

func (l *DefaultLogger) MinLogLevel() LogLevel {
	return l.minLevel
}

func (l *DefaultLogger) write(level LogLevel, message string) {
	if l.minLevel <= level {
		_ = l.loggers[level].Output(2, message)
	}
}

func (l *DefaultLogger) Trace(msg string) {
	l.write(Trace, msg)
}

func (l *DefaultLogger) Tracef(tmpl string, vals ...any) {
	l.write(Trace, Sprintf(tmpl, vals...))
}

func (l *DefaultLogger) Debug(msg string) {
	l.write(Debug, msg)
}

func (l *DefaultLogger) Debugf(tmpl string, vals ...any) {
	l.write(Debug, Sprintf(tmpl, vals...))
}

func (l *DefaultLogger) Info(msg string) {
	l.write(Info, msg)
}

func (l *DefaultLogger) Infof(tmpl string, vals ...any) {
	l.write(Info, Sprintf(tmpl, vals...))
}

func (l *DefaultLogger) Warn(msg string) {
	l.write(Warn, msg)
}

func (l *DefaultLogger) Warnf(tmpl string, vals ...any) {
	l.write(Warn, Sprintf(tmpl, vals...))
}

func (l *DefaultLogger) Panic(msg string) {
	l.write(Fatal, msg)
	if l.triggerPanic {
		panic(msg)
	}
}

func (l *DefaultLogger) Panicf(tmpl string, vals ...any) {
	fmsg := Sprintf(tmpl, vals...)
	l.write(Fatal, fmsg)
	if l.triggerPanic {
		panic(fmsg)
	}
}
