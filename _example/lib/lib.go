package lib

import "github.com/g4s8/go-metalog"
import "github.com/g4s8/go-metalog/sugar"

type Calculator struct {
	log metalog.Logger
}

func NewCalc(log metalog.Logger) *Calculator {
	return &Calculator{log}
}

func (calc *Calculator) Sum(args... int) int {
	var sum int
	for _, x := range args {
		sum += x
	}
	sugar.New(calc.log).WithField("args-count", len(args)).Info("calculated sum");
	return sum
}
