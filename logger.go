package logger

import (
	"context"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logger *zap.SugaredLogger

var levelMap = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

func getLoggerLevel(lvl string) zapcore.Level {
	if level, ok := levelMap[lvl]; ok {
		return level
	}
	return zapcore.InfoLevel
}

func init() {
	fileName := "./app.log"
	level := getLoggerLevel("debug")
	syncWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:  fileName,
		MaxSize:   1 << 30, //1G
		LocalTime: true,
		Compress:  true,
	})

	encoder := zap.NewProductionEncoderConfig()
	encoder.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(zapcore.NewJSONEncoder(encoder), syncWriter, zap.NewAtomicLevelAt(level))
	zapLogger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	logger = zapLogger.Sugar()
}

func Debug(args ...interface{}) {
	logger.Debug(args...)
}

func DebugTrace(ctx context.Context, args ...interface{}) {
	Trace(ctx).Debug(args...)
}

func Debugf(template string, args ...interface{}) {
	logger.Debugf(template, args...)
}

func DebugfTrace(ctx context.Context, template string, args ...interface{}) {
	Trace(ctx).Debugf(template, args...)
}

func Info(args ...interface{}) {
	logger.Info(args...)
}

func InfoTrace(ctx context.Context, args ...interface{}) {
	Trace(ctx).Info(args...)
}

func Infof(template string, args ...interface{}) {
	logger.Infof(template, args...)
}

func InfofTrace(ctx context.Context, template string, args ...interface{}) {
	Trace(ctx).Infof(template, args...)
}

func Warn(args ...interface{}) {
	logger.Warn(args...)
}

func WarnTrace(ctx context.Context, args ...interface{}) {
	Trace(ctx).Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	logger.Warnf(template, args...)
}

func WarnfTrace(ctx context.Context, template string, args ...interface{}) {
	Trace(ctx).Warnf(template, args...)
}

func Error(args ...interface{}) {
	logger.Error(args...)
}

func ErrorTrace(ctx context.Context, args ...interface{}) {
	Trace(ctx).Error(args...)
}

func Errorf(template string, args ...interface{}) {
	logger.Errorf(template, args...)
}

func ErrorfTrace(ctx context.Context, template string, args ...interface{}) {
	Trace(ctx).Errorf(template, args...)
}

func DPanic(args ...interface{}) {
	logger.DPanic(args...)
}

func DPanicTrace(ctx context.Context, args ...interface{}) {
	Trace(ctx).DPanic(args...)
}

func DPanicf(template string, args ...interface{}) {
	logger.DPanicf(template, args...)
}

func DPanicfTrace(ctx context.Context, template string, args ...interface{}) {
	Trace(ctx).DPanicf(template, args...)
}

func Panic(args ...interface{}) {
	logger.Panic(args...)
}

func PanicTrace(ctx context.Context, args ...interface{}) {
	Trace(ctx).Panic(args...)
}

func Panicf(template string, args ...interface{}) {
	logger.Panicf(template, args...)
}

func PanicfTrace(ctx context.Context, template string, args ...interface{}) {
	Trace(ctx).Panicf(template, args...)
}

func Fatal(args ...interface{}) {
	logger.Fatal(args...)
}

func FatalTrace(ctx context.Context, args ...interface{}) {
	Trace(ctx).Fatal(args...)
}

func Fatalf(template string, args ...interface{}) {
	logger.Fatalf(template, args...)
}

func FatalfTrace(ctx context.Context, template string, args ...interface{}) {
	Trace(ctx).Fatalf(template, args...)
}

// logger带上TraceId
func Trace(ctx context.Context) *zap.SugaredLogger {
	var jTraceId jaeger.TraceID
	if parent := opentracing.SpanFromContext(ctx); parent != nil {
		parentCtx := parent.Context()
		if tracer := opentracing.GlobalTracer(); tracer != nil {
			mySpan := tracer.StartSpan("my info", opentracing.ChildOf(parentCtx))
			if sc, ok := mySpan.Context().(jaeger.SpanContext); ok {
				jTraceId = sc.TraceID()
			}
			defer mySpan.Finish()
		}
	}

	return logger.With(zap.String(jaeger.TraceContextHeaderName, fmt.Sprint(jTraceId)))
}
