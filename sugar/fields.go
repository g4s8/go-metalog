// The MIT License (MIT)
// Copyright (c) 2022 Kirill Ch. <g4s8.public@gmail.com>
// https://github.com/g4s8/go-metalog/LICENSE

package sugar

import "github.com/g4s8/go-metalog"

// Field creates new field structure from name and value
func Field(name string, val interface{}) metalog.Field {
	return metalog.Field{Name: name, Value: val}
}
