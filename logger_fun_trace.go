package logger

import "context"

func DebugTrace(ctx context.Context, args ...interface{}) {
	Trace(ctx).Debug(args...)
}

func DebugfTrace(ctx context.Context, template string, args ...interface{}) {
	Trace(ctx).Debugf(template, args...)
}

func InfoTrace(ctx context.Context, args ...interface{}) {
	Trace(ctx).Info(args...)
}

func InfofTrace(ctx context.Context, template string, args ...interface{}) {
	Trace(ctx).Infof(template, args...)
}

func WarnTrace(ctx context.Context, args ...interface{}) {
	Trace(ctx).Warn(args...)
}

func WarnfTrace(ctx context.Context, template string, args ...interface{}) {
	Trace(ctx).Warnf(template, args...)
}

func ErrorTrace(ctx context.Context, args ...interface{}) {
	Trace(ctx).Error(args...)
}

func ErrorfTrace(ctx context.Context, template string, args ...interface{}) {
	Trace(ctx).Errorf(template, args...)
}

func DPanicTrace(ctx context.Context, args ...interface{}) {
	Trace(ctx).DPanic(args...)
}

func DPanicfTrace(ctx context.Context, template string, args ...interface{}) {
	Trace(ctx).DPanicf(template, args...)
}

func PanicTrace(ctx context.Context, args ...interface{}) {
	Trace(ctx).Panic(args...)
}

func PanicfTrace(ctx context.Context, template string, args ...interface{}) {
	Trace(ctx).Panicf(template, args...)
}

func FatalTrace(ctx context.Context, args ...interface{}) {
	Trace(ctx).Fatal(args...)
}

func FatalfTrace(ctx context.Context, template string, args ...interface{}) {
	Trace(ctx).Fatalf(template, args...)
}
