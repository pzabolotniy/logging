package logging

import (
	"fmt"
	"os"
	"strings"

	"github.com/pzabolotniy/logging/pkg/hooks"
	"github.com/rs/zerolog"
)

// ZerologWrapper is a local wrapper around vanilla zerolog-module
// need to implement Logger interface.
type ZerologWrapper struct {
	zerolog.Context
}

// Trace implements Logger interface: creates log message with Trace-level.
func (l *ZerologWrapper) Trace(messages ...any) {
	logger := l.Context.Logger()
	zerologMessages(logger.Trace(), messages)
}

// Debug implements Logger interface: creates log message with Debug-level.
func (l *ZerologWrapper) Debug(messages ...any) {
	logger := l.Context.Logger()
	zerologMessages(logger.Debug(), messages)
}

// Info implements Logger interface: creates log message with Info-level.
func (l *ZerologWrapper) Info(messages ...any) {
	logger := l.Context.Logger()
	zerologMessages(logger.Info(), messages)
}

// Warn implements Logger interface: creates log message with Warn-level.
func (l *ZerologWrapper) Warn(messages ...any) {
	logger := l.Context.Logger()
	zerologMessages(logger.Warn(), messages)
}

// Error implements Logger interface: creates log message with Error-level.
func (l *ZerologWrapper) Error(messages ...any) {
	logger := l.Context.Logger()
	zerologMessages(logger.Error(), messages)
}

// Fatal implements Logger interface: creates log message with Fatal-level.
func (l *ZerologWrapper) Fatal(messages ...any) {
	logger := l.Context.Logger()
	zerologMessages(logger.Fatal(), messages)
}

func zerologMessages(zerologEvent *zerolog.Event, messages []any) {
	if len(messages) != 1 {
		return
	}
	realMessages, ok := messages[0].([]interface{})
	if !ok {
		return
	}
	msgs := make([]string, 0, len(realMessages))
	for _, msg := range messages {
		realMsgs, msgOK := msg.([]interface{})
		if !msgOK {
			continue
		}
		for i := range realMsgs {
			msgs = append(msgs, fmt.Sprintf("%s", realMsgs[i]))
		}
	}
	zerologEvent.Msg(strings.Join(msgs, ", "))
}

// WithField implements Logger interface: add context pair to the log message.
func (l *ZerologWrapper) WithField(key string, value interface{}) Logger {
	return &ZerologWrapper{l.Context.Fields(map[string]interface{}{key: value})}
}

// WithFields implements Logger interface:
// adds context fields to the log message.
func (l *ZerologWrapper) WithFields(fields Fields) Logger {
	zerologFields := make(map[string]interface{})
	for k, v := range fields {
		zerologFields[k] = v
	}

	return &ZerologWrapper{
		l.Context.Fields(zerologFields),
	}
}

// GetZeroLog is a Logger getter with default settings.
func GetZeroLog() Logger {
	logger := zerolog.New(os.Stdout).
		Level(zerolog.TraceLevel).
		Hook(hooks.GetFileLineHook())

	l := &ZerologWrapper{logger.With()}

	return l
}

// WithError implements Logger interface: adds error to the log-message.
func (l *ZerologWrapper) WithError(err error) Logger {
	return &ZerologWrapper{l.Err(err)}
}
