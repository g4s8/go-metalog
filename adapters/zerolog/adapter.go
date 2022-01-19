// The MIT License (MIT)
// Copyright (c) 2022 Kirill Ch. <g4s8.public@gmail.com>
// https://github.com/g4s8/go-metalog/LICENSE

// Package zerolog implements metalog interface
// for zerolog logger.
package zerolog

import (
	"fmt"

	"github.com/g4s8/go-metalog"
	zlog "github.com/rs/zerolog"
)

type adapter struct {
	log *zlog.Logger
}

// WrapLogger creates a new metalog adapter
// for zerolog `Logger` instance.
func WrapLogger(log *zlog.Logger) metalog.Logger {
	return &adapter{log}
}

func (a *adapter) Log(lvl metalog.Level, msg string,
	fields... metalog.Field) {
	evt := a.forLevel(lvl)
	for _, f := range fields {
		evt = addField(evt, &f)
	}
	evt.Msg(msg)
}

func addField(evt *zlog.Event, f *metalog.Field) *zlog.Event {
	switch val := f.Value.(type) {
	case string:
		return evt.Str(f.Name, val)
	case []string:
		return evt.Strs(f.Name, val)
	case error:
		return evt.AnErr(f.Name, val)
	case []error:
		return evt.Errs(f.Name, val)
	case int:
		return evt.Int(f.Name, val)
	case []int:
		return evt.Ints(f.Name, val)
	case int8:
		return evt.Int8(f.Name, val)
	case []int8:
		return evt.Ints8(f.Name, val)
	case int16:
		return evt.Int16(f.Name, val)
	case []int16:
		return evt.Ints16(f.Name, val)
	case int32:
		return evt.Int32(f.Name, val)
	case []int32:
		return evt.Ints32(f.Name, val)
	case int64:
		return evt.Int64(f.Name, val)
	case []int64:
		return evt.Ints64(f.Name, val)
	case uint:
		return evt.Uint(f.Name, val)
	case []uint:
		return evt.Uints(f.Name, val)
	case uint8:
		return evt.Uint8(f.Name, val)
	case uint16:
		return evt.Uint16(f.Name, val)
	case []uint16:
		return evt.Uints16(f.Name, val)
	case uint32:
		return evt.Uint32(f.Name, val)
	case []uint32:
		return evt.Uints32(f.Name, val)
	case uint64:
		return evt.Uint64(f.Name, val)
	case []uint64:
		return evt.Uints64(f.Name, val)
	case []byte:
		return evt.Bytes(f.Name, val)
	case bool:
		return evt.Bool(f.Name, val)
	case []bool:
		return evt.Bools(f.Name, val)
	case fmt.Stringer:
		return evt.Stringer(f.Name, val)
	default:
		return evt.Interface(f.Name, val)
	}
}

func (a *adapter) forLevel(lvl metalog.Level) *zlog.Event {
	switch lvl {
	case metalog.DebugLevel:
		return a.log.Debug()
	case metalog.InfoLevel:
		return a.log.Info()
	case metalog.WarnLevel:
		return a.log.Warn()
	case metalog.ErrorLevel:
		return a.log.Error()
	case metalog.PanicLevel:
		return a.log.Panic()
	default:
		return a.log.Info()
	}
}

