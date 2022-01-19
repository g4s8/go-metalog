// The MIT License (MIT)
// Copyright (c) 2022 Kirill Ch. <g4s8.public@gmail.com>
// https://github.com/g4s8/go-metalog/LICENSE

// Package zap implements metalog interfaces for
// zap logger.
package zap

import (
	"fmt"
	"time"

	"github.com/g4s8/go-metalog"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type adapter struct {
	zlog *zap.Logger
}

// WrapLogger creates a new metalog adapter
// for zap `Logger` instance.
func WrapLogger(zlog *zap.Logger) metalog.Logger {
	return &adapter{zlog}
}

func (a *adapter) Log(lvl metalog.Level, msg string,
	fields... metalog.Field) {
	zfields := make([]zap.Field, len(fields))
	for i, f := range fields {
		var zf zap.Field
		zapField(&f, &zf)
		zfields[i] = zf
	}
	switch lvl {
	case metalog.DebugLevel:
		a.zlog.Debug(msg, zfields...)
	case metalog.InfoLevel:
		a.zlog.Info(msg, zfields...)
	case metalog.WarnLevel:
		a.zlog.Warn(msg, zfields...)
	case metalog.ErrorLevel:
		a.zlog.Error(msg, zfields...)
	case metalog.PanicLevel:
		a.zlog.Panic(msg, zfields...)
	default:
		a.zlog.Info(msg, zfields...)
	}
}

func zapFieldArray(val interface{}, out *zap.Field) {
	out.Type = zapcore.ArrayMarshalerType
	out.Interface = val
}

func zapField(field *metalog.Field, out *zap.Field) {
	out.Key = field.Name
	switch val := field.Value.(type) {
	case []byte:
		out.Type = zapcore.BinaryType
		out.Interface = val
	case bool:
		out.Type = zapcore.BoolType
		var ival int64
		if val {
			ival = 1
		}
		out.Integer = ival
	case []bool:
		zapFieldArray(val, out)
	case complex64:
		out.Type = zapcore.Complex64Type
		out.Interface = val
	case []complex64:
		zapFieldArray(val, out)
	case complex128:
		out.Type = zapcore.Complex128Type
		out.Interface = val
	case []complex128:
		zapFieldArray(val, out)
	case float64:
		out.Type = zapcore.Float64Type
		out.Interface = val
	case []float64:
		zapFieldArray(val, out)
	case float32:
		out.Type = zapcore.Float32Type
		out.Interface = val
	case []float32:
		zapFieldArray(val, out)
	case int:
		out.Type = zapcore.Int64Type
		out.Integer = int64(val)
	case []int:
		zapFieldArray(val, out)
	case int8:
		out.Type = zapcore.Int8Type
		out.Integer = int64(val)
	case []int8:
		zapFieldArray(val, out)
	case int16:
		out.Type = zapcore.Int16Type
		out.Integer = int64(val)
	case []int16:
		zapFieldArray(val, out)
	case int32:
		out.Type = zapcore.Int32Type
		out.Integer = int64(val)
	case []int32:
		zapFieldArray(val, out)
	case int64:
		out.Type = zapcore.Int64Type
		out.Integer = val
	case []int64:
		zapFieldArray(val, out)
	case uint:
		out.Type = zapcore.Uint64Type
		out.Integer = int64(val)
	case []uint:
		zapFieldArray(val, out)
	case uint8:
		out.Type = zapcore.Uint8Type
		out.Integer = int64(val)
	case uint16:
		out.Type = zapcore.Uint16Type
		out.Integer = int64(val)
	case []uint16:
		zapFieldArray(val, out)
	case uint32:
		out.Type = zapcore.Uint32Type
		out.Integer = int64(val)
	case []uint32:
		zapFieldArray(val, out)
	case uint64:
		out.Type = zapcore.Uint64Type
		out.Integer = int64(val)
	case []uint64:
		zapFieldArray(val, out)
	case string:
		out.Type = zapcore.StringType
		out.String = val
	case []string:
		zapFieldArray(val, out)
	case fmt.Stringer:
		out.Type = zapcore.StringerType
		out.Interface = val
	case time.Time:
		out.Type = zapcore.TimeType
		out.Interface = val
	case time.Duration:
		out.Type = zapcore.DurationType
		out.Interface = val
	case error:
		out.Type = zapcore.ErrorType
		out.Interface = val
	default:
		out.Type = zapcore.ReflectType
		out.Interface = val
	}
}
