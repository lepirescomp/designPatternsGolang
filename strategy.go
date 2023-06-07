package main

import (
	"fmt"
)

type ABExoeriment interface {
	CalculateEarnings(amount []float64)
	CalculateCosts(amount []float64)
}

type BrazilABExperiment struct{}

func (b *BrazilABExperiment) CalculateEarnings(amount []float64) {
	var result float64
	for _, v := range amount {
		result += v
	}
	fmt.Println("Printing earnings for strategy of Brazil experiment", result)
}

func (b *BrazilABExperiment) CalculateCosts(amount []float64) {
	var result float64
	for _, v := range amount {
		result += v
	}
	fmt.Println("Printing costs for strategy of Brazil experiment", result)
}

type USABExperiment struct{}

func (b *USABExperiment) CalculateEarnings(amount []float64) {
	var result float64
	for _, v := range amount {
		result += v
	}
	fmt.Println("Printing earnings for strategy of US experiment", result)
}

func (b *USABExperiment) CalculateCosts(amount []float64) {
	var result float64
	for _, v := range amount {
		result += v
	}
	fmt.Println("Printing costs for strategy of US experiment", result)
}

type Experiment struct {
	country_code string
	exp          ABExoeriment
}

func (e *Experiment) SetABExperiment() {
	switch e.country_code {
	case "BR":
		e.exp = &BrazilABExperiment{}
	case "US":
		e.exp = &USABExperiment{}
	}
}

func main() {

	e := Experiment{
		country_code: "BR",
	}
	e.SetABExperiment()
	e.exp.CalculateEarnings([]float64{1.2, 2.4})

	e = Experiment{
		country_code: "US",
	}
	e.SetABExperiment()
	e.exp.CalculateEarnings([]float64{1.2, 2.4})
}
