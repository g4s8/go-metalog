// The MIT License (MIT)
// Copyright (c) 2022 Kirill Ch. <g4s8.public@gmail.com>
// https://github.com/g4s8/go-metalog/LICENSE

// Package logrus implements metalog interface for
// logrus logger.
package logrus

import (
	"github.com/g4s8/go-metalog"
	log "github.com/sirupsen/logrus"
)

type adapter struct {
	log *log.Logger
}


// WrapLogger creates a new metalog adapter
// for logrus `Logger` instance. To use default
// logrus logger call it with `WrapLogger(logrus.New())`
func WrapLogger(log *log.Logger) metalog.Logger {
	return &adapter{log}
}

func (a *adapter) Log(lvl metalog.Level, msg string,
	fields... metalog.Field) {
	lf := a.log.WithFields(convertFields(fields))
	lf.Log(convertLevel(lvl), msg)
}

func convertLevel(lvl metalog.Level) log.Level {
	switch lvl {
	case metalog.DebugLevel:
		return log.DebugLevel
	case metalog.InfoLevel:
		return log.InfoLevel
	case metalog.WarnLevel:
		return log.WarnLevel
	case metalog.ErrorLevel:
		return log.ErrorLevel
	case metalog.PanicLevel:
		return log.PanicLevel
	default:
		return log.InfoLevel
	}
}

func convertFields(fields []metalog.Field) log.Fields {
	res := make(map[string]interface{}, len(fields))
	for _, f := range fields {
		res[f.Name] = f.Value
	}
	return log.Fields(res)
}
