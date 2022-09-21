package logging

import (
	"fmt"
	"os"
	"strings"

	"github.com/pzabolotniy/logging/pkg/hooks"
	"github.com/rs/zerolog"
)

type ZerologWrapper struct {
	zerolog.Context
}

func (l *ZerologWrapper) Trace(messages ...any) {
	logger := l.Context.Logger()
	zerologMessages(logger.Trace(), messages)
}

func (l *ZerologWrapper) Debug(messages ...any) {
	logger := l.Context.Logger()
	zerologMessages(logger.Debug(), messages)
}

func (l *ZerologWrapper) Info(messages ...any) {
	logger := l.Context.Logger()
	zerologMessages(logger.Info(), messages)
}

func (l *ZerologWrapper) Warn(messages ...any) {
	logger := l.Context.Logger()
	zerologMessages(logger.Warn(), messages)
}

func (l *ZerologWrapper) Error(messages ...any) {
	logger := l.Context.Logger()
	zerologMessages(logger.Error(), messages)
}

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

func (l *ZerologWrapper) WithField(key string, value interface{}) Logger {
	return &ZerologWrapper{l.Context.Fields(map[string]interface{}{key: value})}
}

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
	logger := zerolog.New(os.Stdout).Level(zerolog.TraceLevel).Hook(hooks.GetFileLineHook())

	l := &ZerologWrapper{logger.With()}

	return l
}

func (l *ZerologWrapper) WithError(err error) Logger {
	return &ZerologWrapper{l.Err(err)}
}
