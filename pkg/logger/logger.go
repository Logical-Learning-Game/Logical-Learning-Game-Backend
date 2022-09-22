package logger

var GlobalLog Logger

type Logger interface {
	Debugw(message string, keyAndValues ...interface{})
	Errorw(message string, keyAndValues ...interface{})
	Infow(message string, keyAndValues ...interface{})
	Warnw(message string, keyAndValues ...interface{})
	Fatalw(message string, keyAndValues ...interface{})
	Error(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Fatal(args ...interface{})
}

func SetGlobalLogger(l Logger) {
	GlobalLog = l
}
