package core

import (
	"os"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	logger *logrus.Logger
}

func NewLogger() *Logger {
	return &Logger{
		logger: &logrus.Logger{
			Out: os.Stdout,
			Formatter: &logrus.TextFormatter{
				FullTimestamp:          true,
				DisableLevelTruncation: true,
			},
			Level: logrus.DebugLevel,
		},
	}
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	l.logger.Debugf(format, args...)
}

func (l *Logger) Infof(format string, args ...interface{}) {
	l.logger.Infof(format, args...)
}

func (l *Logger) Warnf(format string, args ...interface{}) {
	l.logger.Warnf(format, args...)
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	l.logger.Errorf(format, args...)
}

func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.logger.Fatalf(format, args...)
}

func (l *Logger) Panicf(format string, args ...interface{}) {
	l.logger.Fatalf(format, args...)
}

func (l *Logger) Debug(args ...interface{}) {
	l.logger.Debug(args...)
}

func (l *Logger) Info(args ...interface{}) {
	l.logger.Info(args...)
}

func (l *Logger) Warn(args ...interface{}) {
	l.logger.Warn(args...)
}

func (l *Logger) Error(args ...interface{}) {
	l.logger.Error(args...)
}

func (l *Logger) Fatal(args ...interface{}) {
	l.logger.Fatal(args...)
}

func (l *Logger) Panic(args ...interface{}) {
	l.logger.Fatal(args...)
}
