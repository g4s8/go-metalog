module example.com/cli

replace example.com/lib => ../lib

replace github.com/g4s8/go-metalog => ../../../go-metalog

go 1.17

require example.com/lib v0.0.0-00010101000000-000000000000

require (
	github.com/g4s8/go-metalog v0.0.0-00010101000000-000000000000 // indirect
	github.com/rs/zerolog v1.26.1 // indirect
	github.com/sirupsen/logrus v1.8.1 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.7.0 // indirect
	go.uber.org/zap v1.20.0 // indirect
	golang.org/x/sys v0.0.0-20210809222454-d867a43fc93e // indirect
)

// require github.com/g4s8/go-metalog v0.0.0-00010101000000-000000000000 // indirect
