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
var loggerFileName = "./app.log" // 日志文件名

func init() {
	sugar()
}

func sugar() {
	syncWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:  loggerFileName,
		MaxSize:   1 << 30, //1G
		LocalTime: true,
		Compress:  true,
	})

	encoder := zap.NewProductionEncoderConfig()
	encoder.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(zapcore.NewJSONEncoder(encoder), syncWriter, zap.NewAtomicLevelAt(zapcore.DebugLevel))
	zapLogger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	logger = zapLogger.Sugar()
}

func New(filename string) {
	if filename != "" {
		loggerFileName = filename
	}
	sugar()
}

func Close() {
	logger.Sync()
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
