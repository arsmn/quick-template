package logger

import (
	"github.com/sirupsen/logrus"
)

// Logger defines structured app logger
// that should be injected into layers
type Logger struct {
	logrus *logrus.Logger
}

func New() *Logger {
	l := logrus.New()
	l.SetFormatter(&logrus.JSONFormatter{})

	return &Logger{
		logrus: l,
	}
}

func (l *Logger) WithField(key string, value any) *Logger {
	with := l.logrus.WithField(key, value)
	return &Logger{
		logrus: with.Logger,
	}
}

func (l *Logger) Debug(args ...any) {
	l.logrus.Debug(args...)
}

func (l *Logger) Info(args ...any) {
	l.logrus.Info(args...)
}

func (l *Logger) Warn(args ...any) {
	l.logrus.Warn(args...)
}

func (l *Logger) Error(args ...any) {
	l.logrus.Error(args...)
}

func (l *Logger) Fatal(args ...any) {
	l.logrus.Fatal(args...)
}

func (l *Logger) Panic(args ...any) {
	l.logrus.Panic(args...)
}
