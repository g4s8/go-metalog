Metalog is a standard API for structured logging and adapters
for its implementation.

# Motivation

Go has a lot of good structured logging libraries, the most popular
are [zap](https://github.com/uber-go/zap),
[logrus](https://github.com/Sirupsen/logrus),
[zerolog](https://github.com/rs/zerolog), and some others. The problem is
they don't have some standard interfaces, it means it's almost impossible
to use it in libraries (because it would be not really good if library
create some strong dependency of logging implementation and requires logger
dependency).

Existing `log` API from standard Go library is not suitable for structured
logging, and it doesn't provide interfaces too.

# Proposal

To solve this, I'm proposing some standard API for structured
logging as `metalog.Logger` interface:
```go
// Logger interface for structured loggers
type Logger interface {
	// Log message with fields
	Log(lvl Level, msg string, fields... Field)
}
```
It's minimal interface that could be used as API for any
structured logger. So any library can depend on this
interfaces and require applicaton code to provide implemenations
of any logger. Library can call `Logger` interface directly or
use `sugar` package for more friendly logging methods. Library
client can use any adapter from `adapters/` sub-package to use
adapters for zap, logrus and zerolog loggers or can implement
adapter by-self. So library doesn't create any strong dependency
on any logger but has an ability to use it.

## Usage

Check `_examples` dir for some small examples:
 - `lib` - a library uses `Logger`
 - `cli` - command-line interface which uses `lib` and some logger:
   - `cmd/logrus` uses logrus
   - `cmd/zap` uses zal
   - `cmd/zerolog` uses zerolog

### For library

Library needs dependency:
`go get -u github.com/g4s8/go-metalog`.

Public methods which requires logger can declare `Logger`
arguments or fields:
```
import "github.com/g4s8/go-metalog"

func NewService(logger metalog.Logger) *Service {...}
```

This `logger` could be used with `Log` method or accessed with
help of `sugar` package:
```
import "github.com/g4s8/go-metalog/sugar"

func (s *Service) Run() {
  s := sugar.New(s.logger)
  s.WithFields(sugar.Field("worker", s.id)).Debug("service started")
}
```

### For applications

Application decides which logger implementation it's using,
then it wraps logger instance with one of metalog adapters,
and passes it to library method:
```go
import (
   "example.com/lib"
   metalogrus "github.com/g4s8/go-metalog/adapters/logrus"
   "github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()
	wrap := metalogrus.WrapLogger(log)

        s := lib.New(wrap)
        s.Run()
}
```
