// The MIT License (MIT)
// Copyright (c) 2022 Kirill Ch. <g4s8.public@gmail.com>
// https://github.com/g4s8/go-metalog/LICENSE

// Package sugar provides additional syntax sugar for
// structured logging API. `New` method creates `Logger`
// instance with fluent methods for fields building and
// friendly method for levels.
package sugar

import "github.com/g4s8/go-metalog"

// Logger with syntax sugar
type Logger struct {
	logger metalog.Logger
	fields []metalog.Field
}

// New logger with sugar
func New(logger metalog.Logger) *Logger {
	return &Logger{
		logger: logger,
		fields: []metalog.Field{},
	}
}

// WithField creates a new logger with additional field entry
func (l *Logger) WithField(name string, val interface{}) *Logger {
	return l.WithFields(Field(name, val))
}

// WithFields creates a new logger with additiona fields entries
func (l *Logger) WithFields(fs... metalog.Field) *Logger {
	res := &Logger{logger: l.logger}
	res.fields = append(l.fields, fs...)
	return res
}

// Log current fields with given level and message
func (l *Logger) Log(lvl metalog.Level, msg string) {
	l.logger.Log(lvl, msg, l.fields...)
}

// Debug log a message
func (l *Logger) Debug(msg string) {
	l.Log(metalog.DebugLevel, msg)
}

// Info log a message
func (l *Logger) Info(msg string) {
	l.Log(metalog.InfoLevel, msg)
}

// Warn log a message
func (l *Logger) Warn(msg string) {
	l.Log(metalog.WarnLevel, msg)
}

// Err log a message
func (l *Logger) Err(msg string) {
	l.Log(metalog.ErrorLevel, msg)
}

// Panic log a message
func (l *Logger) Panic(msg string) {
	l.Log(metalog.PanicLevel, msg)
}
