package logging

import "context"

type logFieldsCtx struct{}

func WithContext(ctx context.Context, fields Fields) context.Context {
	return context.WithValue(ctx, logFieldsCtx{}, fields)
}

func ReplaceFieldsInContext(ctx context.Context, newFields Fields) context.Context {
	currentFields := FieldsFromContext(ctx)
	for k, v := range newFields {
		currentFields[k] = v
	}
	return WithContext(ctx, currentFields)
}

func FieldsFromContext(ctx context.Context) Fields {
	fields, ok := ctx.Value(logFieldsCtx{}).(Fields)
	if !ok {
		return make(Fields, 0)
	}
	return fields
}

func FromContext(ctx context.Context, logger Logger) Logger {
	return logger.WithFields(FieldsFromContext(ctx))
}
