// The MIT License (MIT)
// Copyright (c) 2022 Kirill Ch. <g4s8.public@gmail.com>
// https://github.com/g4s8/go-metalog/LICENSE

// Package metalog provides standard API for structured
// logging. API provides as `Logger` interface that
// could be used by libraries without requiring logger
// implementation as strong dependency.
package metalog

// Logger interface for structured loggers
type Logger interface {
	// Log message with fields
	Log(lvl Level, msg string, fiels... Field)
}

// Level of logging
type Level uint8

const (
	// DebugLevel is lowest level of logging,
	// is used for debugging message
	DebugLevel Level = iota
	// InfoLevel is standard log messages
	InfoLevel
	// WarnLevel has higher priority than info,
	// can be used for some important messages
	WarnLevel
	// ErrorLevel indicates that some error happens,
	// it doesn't stop the program
	ErrorLevel
	// PanicLevel is the highest level, it log error
	// message and exits program with panic
	PanicLevel
)

// Field of log entry
type Field struct {
	// Field name (key)
	Name string

	// Field value
	Value interface{}
}
