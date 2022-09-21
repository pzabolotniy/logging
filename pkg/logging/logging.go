package logging

// Logger interface.
type Logger interface {
	Trace(messages ...any)
	Debug(messages ...any)
	Info(messages ...any)
	Warn(messages ...any)
	Error(messages ...any)
	Fatal(messages ...any)
	WithError(err error) Logger
	WithField(key string, value interface{}) Logger
	WithFields(fields Fields) Logger
}

// GetLogger is a Logger getter with default settings.
func GetLogger() Logger {
	return GetZeroLog()
}
