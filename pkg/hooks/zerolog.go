package hooks

import (
	"fmt"

	"github.com/rs/zerolog"
)

const (
	// ZeroLogDepth is a depth of stack trace
	// this depth is required to find caller correctly.
	ZeroLogDepth int = 6
)

// Run implements zerolog.Hook interface.
func (hook *FileLineHook) Run(zerologEvent *zerolog.Event, _ zerolog.Level, _ string) {
	var (
		file string
		line int
	)
	file, line = getCaller(ZeroLogDepth)

	zerologEvent.Str(hook.LogKeyName, fmt.Sprintf("%s:%d", file, line))
}
