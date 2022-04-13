package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log Logger

type Logger interface {
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Fatal(args ...interface{})
	Infof(format string, args ...interface{})
	Info(args ...interface{})
	Warnf(format string, args ...interface{})
	Debugf(format string, args ...interface{})
	Debug(args ...interface{})
	Printf(format string, args ...interface{})
	Infow(msg string, keysAndValues ...interface{})
}

type loggerWrapper struct {
	lw *zap.SugaredLogger
}

func (logger *loggerWrapper) Info(args ...interface{}) {
	logger.lw.Info(args...)
}

func (logger *loggerWrapper) Infow(msg string, keysAndValues ...interface{}) {
	logger.lw.Infow(msg, keysAndValues...)
}

func (logger *loggerWrapper) Debug(args ...interface{}) {
	logger.lw.Debug(args...)
}

func (logger *loggerWrapper) Errorf(format string, args ...interface{}) {
	logger.lw.Errorf(format, args...)
}
func (logger *loggerWrapper) Fatalf(format string, args ...interface{}) {
	logger.lw.Fatalf(format, args...)
}
func (logger *loggerWrapper) Fatal(args ...interface{}) {
	logger.lw.Fatal(args...)
}
func (logger *loggerWrapper) Infof(format string, args ...interface{}) {
	logger.lw.Infof(format, args...)
}
func (logger *loggerWrapper) Warnf(format string, args ...interface{}) {
	logger.lw.Warnf(format, args...)
}
func (logger *loggerWrapper) Debugf(format string, args ...interface{}) {
	logger.lw.Debugf(format, args...)
}
func (logger *loggerWrapper) Printf(format string, args ...interface{}) {
	logger.lw.Infof(format, args...)
}

func NewZapLogger() Logger {
	cfg := zap.NewProductionConfig()
	cfg.Level.SetLevel(zapcore.InfoLevel)
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.EncoderConfig.TimeKey = "ts"
	cfg.DisableCaller = true
	logger, _ := cfg.Build()
	return &loggerWrapper{
		lw: logger.Sugar(),
	}

}
