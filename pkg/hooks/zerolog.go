package hooks

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/rs/zerolog"
)

const (
	// ZeroLogStartDepth is a depth of stack trace
	// this depth is required to find caller correctly.
	ZeroLogStartDepth int = 6
)

// Run implements zerolog.Hook interface.
func (hook *FileLineHook) Run(zerologEvent *zerolog.Event, level zerolog.Level, message string) {
	var (
		file string
		line int
	)
	for i := 0; i < 10; i++ { //nolint:revive // allow magic number
		file, line = getCaller1(ZeroLogStartDepth + i)
		if !strings.HasPrefix(file, "zerolog") {
			break
		}
	}

	zerologEvent.Str(hook.LogKeyName, fmt.Sprintf("%s:%d", file, line))
}

func getCaller1(skip int) (string, int) {
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		return "", 0
	}

	n := 0
	for i := len(file) - 1; i > 0; i-- {
		if file[i] == '/' {
			n++
			if n >= PathLen {
				file = file[i+1:]

				break
			}
		}
	}

	return file, line
}
