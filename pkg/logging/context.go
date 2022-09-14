package logging

import "context"

type logCtx struct{}

// WithContext puts logger to the context.
func WithContext(ctx context.Context, logger Logger) context.Context {
	return context.WithValue(ctx, logCtx{}, logger)
}

// FromContext extracts Logger from context and returns itself
// otherwise, creates default one logger.
func FromContext(ctx context.Context) Logger {
	logger, ok := ctx.Value(logCtx{}).(Logger)
	if !ok {
		return GetLogger()
	}

	return logger
}
