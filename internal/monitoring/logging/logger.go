package logging

import (
	"context"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	*zap.Logger
}

type LoggerOption func(*zap.Config)

func NewLogger(opts ...LoggerOption) (*Logger, error) {
	config := zap.NewProductionConfig()
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	// Apply options
	for _, opt := range opts {
		opt(&config)
	}

	logger, err := config.Build(
		zap.AddCallerSkip(1),
		zap.AddStacktrace(zapcore.ErrorLevel),
	)
	if err != nil {
		return nil, err
	}

	return &Logger{logger}, nil
}

// WithContext adds context fields to the logger
func (l *Logger) WithContext(ctx context.Context) *zap.Logger {
    // Add trace ID if available
    if traceID := ctx.Value("traceID"); traceID != nil {
        return l.With(zap.String("traceID", traceID.(string)))
    }
    return l.Logger
}
