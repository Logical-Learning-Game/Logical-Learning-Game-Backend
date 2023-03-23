package logger

import "go.uber.org/zap"

type ZapLogger struct {
	sugar *zap.SugaredLogger
}

func NewZapLogger(sugar *zap.SugaredLogger) Logger {
	return &ZapLogger{
		sugar: sugar,
	}
}

func (z ZapLogger) Debugf(template string, args ...interface{}) {
	z.sugar.Debugf(template, args...)
}

func (z ZapLogger) Errorf(template string, args ...interface{}) {
	z.sugar.Errorf(template, args...)
}

func (z ZapLogger) Infof(template string, args ...interface{}) {
	z.sugar.Infof(template, args...)
}

func (z ZapLogger) Warnf(template string, args ...interface{}) {
	z.sugar.Warnf(template, args...)
}

func (z ZapLogger) Fatalf(template string, args ...interface{}) {
	z.sugar.Fatalf(template, args...)
}

func (z ZapLogger) Debug(args ...interface{}) {
	z.sugar.Debug(args...)
}

func (z ZapLogger) Debugw(message string, keyAndValues ...interface{}) {
	z.sugar.Debugw(message, keyAndValues...)
}

func (z ZapLogger) Errorw(message string, keyAndValues ...interface{}) {
	z.sugar.Errorw(message, keyAndValues...)
}

func (z ZapLogger) Infow(message string, keyAndValues ...interface{}) {
	z.sugar.Infow(message, keyAndValues...)
}

func (z ZapLogger) Warnw(message string, keyAndValues ...interface{}) {
	z.sugar.Warnw(message, keyAndValues...)
}

func (z ZapLogger) Fatalw(message string, keyAndValues ...interface{}) {
	z.sugar.Fatalw(message, keyAndValues...)
}

func (z ZapLogger) Error(args ...interface{}) {
	z.sugar.Error(args...)
}

func (z ZapLogger) Info(args ...interface{}) {
	z.sugar.Info(args...)
}

func (z ZapLogger) Warn(args ...interface{}) {
	z.sugar.Warn(args...)
}

func (z ZapLogger) Fatal(args ...interface{}) {
	z.sugar.Fatal(args...)
}
