package logger

var GlobalLog Logger

type Logger interface {
	Debugw(message string, keyAndValues ...interface{})
	Errorw(message string, keyAndValues ...interface{})
	Infow(message string, keyAndValues ...interface{})
	Warnw(message string, keyAndValues ...interface{})
	Fatalw(message string, keyAndValues ...interface{})
	Debugf(template string, args ...interface{})
	Errorf(template string, args ...interface{})
	Infof(template string, args ...interface{})
	Warnf(template string, args ...interface{})
	Fatalf(template string, args ...interface{})
	Debug(args ...interface{})
	Error(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Fatal(args ...interface{})
}

func SetGlobalLogger(l Logger) {
	GlobalLog = l
}
