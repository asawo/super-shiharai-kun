package logger

import (
	"strings"

	"github.com/pkg/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Log struct {
	zap   *zap.Logger
	sugar *zap.SugaredLogger
}

// New creates a new logger
func New(level string) (*Log, error) {
	l, err := logLevel(level)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse log level")
	}

	config := zap.NewProductionConfig()
	config.Level = zap.NewAtomicLevelAt(l)
	config.DisableStacktrace = true
	config.OutputPaths = []string{"stdout"}
	config.ErrorOutputPaths = []string{"stderr"}

	zapLogger, err := config.Build()
	if err != nil {
		return nil, errors.Wrap(err, "failed to create zap logger")
	}

	return &Log{
		zap:   zapLogger,
		sugar: zapLogger.Sugar(),
	}, nil
}

// logLevel converts string level into zap log level
func logLevel(level string) (zapcore.Level, error) {
	level = strings.ToUpper(level)
	var l zapcore.Level

	switch level {
	case "DEBUG":
		l = zapcore.DebugLevel
	case "INFO":
		l = zapcore.InfoLevel
	case "ERROR":
		l = zapcore.ErrorLevel
	default:
		return l, errors.Errorf("invalid loglevel: %s", level)
	}

	return l, nil
}

// ZapLogger creates a zap.Logger
func (l *Log) ZapLogger() *zap.Logger {
	return l.zap
}

// Info logs with INFO log level
func (l *Log) Info(msg string, fields ...zapcore.Field) {
	l.zap.Info(msg, fields...)
}

// Warn logs with WARN log level
func (l *Log) Warn(msg string, fields ...zapcore.Field) {
	l.zap.Warn(msg, fields...)
}

// Debug logs with DEBUG log level
func (l *Log) Debug(msg string, fields ...zapcore.Field) {
	l.zap.Debug(msg, fields...)
}

// Error logs with ERROR log level
func (l *Log) Error(msg string, fields ...zapcore.Field) {
	l.zap.Error(msg, fields...)
}

// Infof logs unstructured log with INFO level
func (l *Log) Infof(msg string, args ...interface{}) {
	l.sugar.Infof(msg, args...)
}

// Warnf logs unstructured log with WARN level
func (l *Log) Warnf(msg string, args ...interface{}) {
	l.sugar.Warnf(msg, args...)
}

// Debugf logs unstructured log with DEBUG level
func (l *Log) Debugf(msg string, args ...interface{}) {
	l.sugar.Debugf(msg, args...)
}

// Errorf logs unstructured log with ERROR level
func (l *Log) Errorf(msg string, args ...interface{}) {
	l.sugar.Errorf(msg, args...)
}
