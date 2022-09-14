// Package logging is a local wrapper around logrus package
// if, suddenly, logrus will become deprecated module
// it will be possible to move to another solution
package logging

import (
	"os"

	"github.com/pzabolotniy/logging/pkg/hooks"
	"github.com/sirupsen/logrus"
)

// Fields is a wrapper around logrus.Field.
type Fields logrus.Fields

// logWrapper is a wrapper around *logrus.Entry.
type logWrapper struct {
	*logrus.Entry
}

// Logger interface.
type Logger interface {
	Trace(...interface{})
	Debug(...interface{})
	Info(...interface{})
	Warn(...interface{})
	Error(...interface{})
	Fatal(...interface{})
	WithError(err error) Logger
	WithField(key string, value interface{}) Logger
	WithFields(fields Fields) Logger
}

// GetLogger is a Logger getter with default settings.
func GetLogger() Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{ //nolint:exhaustruct // set only modified fields
		FullTimestamp: true,
	})
	logger.SetLevel(logrus.TraceLevel)
	logger.SetOutput(os.Stdout)
	logger.AddHook(hooks.GetFileLineHook())
	l := &logWrapper{logger.WithFields(nil)}

	return l
}

func (lw *logWrapper) WithError(err error) Logger {
	return &logWrapper{lw.Entry.WithError(err)}
}

func (lw *logWrapper) WithField(key string, value interface{}) Logger {
	return &logWrapper{lw.Entry.WithField(key, value)}
}

func (lw *logWrapper) WithFields(fields Fields) Logger {
	return &logWrapper{lw.Entry.WithFields(logrus.Fields(fields))}
}
