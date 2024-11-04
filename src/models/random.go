package models

import "gonum.org/v1/gonum/stat/distuv"

type Random struct {
}

func NewRandom() *Random {
	return &Random{}
}

func (pd *Random) Generate(rate float64) float64 {
	poisson := distuv.Poisson{Lambda: rate, Src: nil}

	return poisson.Rand()
}
