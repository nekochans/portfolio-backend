package infrastructure

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Logger(l *zap.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			writer := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

			now := time.Now()
			defer func() {
				// TODO r.Bodyをログに出したほうが良さそう
				l.Info("message",
					zap.String("protocol", r.Proto),
					zap.String("path", r.URL.Path),
					zap.Duration("lat", time.Since(now)),
					zap.Int("HttpStatus", writer.Status()),
					zap.Int("size", writer.BytesWritten()),
					zap.String("RequestId", middleware.GetReqID(r.Context())))
			}()

			next.ServeHTTP(writer, r)
		}
		return http.HandlerFunc(fn)
	}
}

func CreateLogger() *zap.Logger {
	level := zap.NewAtomicLevel()
	level.SetLevel(zapcore.DebugLevel)

	zapConfig := zap.Config{
		Level:    level,
		Encoding: "json",
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "time",
			LevelKey:       "level",
			NameKey:        "name",
			CallerKey:      "caller",
			MessageKey:     "message",
			StacktraceKey:  "stackTrace",
			EncodeLevel:    zapcore.CapitalLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}
	logger, _ := zapConfig.Build(zap.AddCallerSkip(1))

	return logger
}
