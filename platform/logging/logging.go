package logging

type LogLevel uint8

const (
	Trace LogLevel = iota
	Debug
	Info
	Warn
	Fatal
	None
)

// We define interfaces for all the features that the platform provides and use
// those interfaces to provide default implementations. This will allow the
// application to replace the default implementation if required and also make
// it possible to provide applications with features as services.
type Logger interface {
	Trace(string)
	Tracef(string, ...any)

	Debug(string)
	Debugf(string, ...any)

	Info(string)
	Infof(string, ...any)

	Warn(string)
	Warnf(string, ...any)
	Panic(string)
	Panicf(string, ...any)
}
