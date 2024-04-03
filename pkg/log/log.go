package log

import (
	"context"
	"github.com/sirupsen/logrus"
	"io"
)

type Fields struct {
	FieldName string
	Val       string
}

type ILog interface {
	Info(ctx context.Context, message string, data ...Fields)
	Error(ctx context.Context, message string, data ...Fields)
	Warn(ctx context.Context, message string, data ...Fields)
	Debug(ctx context.Context, message string, data ...Fields)
	Fatal(ctx context.Context, message string, data ...Fields)
}

type Logger struct {
	logHandler *logrus.Logger
}

var _ ILog = (*Logger)(nil)

func New(output io.Writer) *Logger {
	logHandler := logrus.New()
	logHandler.SetOutput(output)
	logHandler.SetFormatter(&logrus.JSONFormatter{})
	return &Logger{
		logHandler: logHandler,
	}
}

func (l *Logger) getLogIns(fields []Fields) *logrus.Entry {
	f := make(logrus.Fields, 0)
	for k := range fields {
		f[fields[k].FieldName] = fields[k].Val
	}
	return l.logHandler.WithFields(f)
}

func (l *Logger) Debug(ctx context.Context, message string, fields ...Fields) {
	l.getLogIns(fields).Log(logrus.DebugLevel, message)
}

func (l *Logger) Info(ctx context.Context, message string, fields ...Fields) {
	l.getLogIns(fields).Log(logrus.InfoLevel, message)
}

func (l *Logger) Warn(ctx context.Context, message string, fields ...Fields) {
	l.getLogIns(fields).Log(logrus.WarnLevel, message)
}

func (l *Logger) Error(ctx context.Context, message string, fields ...Fields) {
	l.getLogIns(fields).Log(logrus.ErrorLevel, message)
}

func (l *Logger) Fatal(ctx context.Context, message string, fields ...Fields) {
	l.getLogIns(fields).Log(logrus.FatalLevel, message)
}
