package main

import "fmt"

// NxN problems: Tip/Commission to ElectricTechnician/BuildingTechnician,
// avoid TipElectricTechnician, TipBuildingTechnician, CommissionBuildingTechnician ...

type Calculator interface {
	calculateElectricTechnician(amount float64, number_of_lights_installed int)
	calculateBuildingTechnician(amount float64, number_of_walls int)
}

type TipCalculator struct {
}

type CommissionCalculator struct {
}

func (v *TipCalculator) calculateElectricTechnician(amount float64, number_of_lights_installed int) {
	r := amount * float64(number_of_lights_installed)
	fmt.Println("Calculating tip of electric technician", r)
}

func (v *TipCalculator) calculateBuildingTechnician(amount float64, number_of_walls int) {
	r := amount * float64(number_of_walls)
	fmt.Println("Calculating tip of building technician", r)
}

func (v *CommissionCalculator) calculateElectricTechnician(amount float64, number_of_lights_installed int) {
	r := amount * float64(number_of_lights_installed)
	fmt.Println("Calculating commission of electric technician", r)
}

func (v *CommissionCalculator) calculateBuildingTechnician(amount float64, number_of_walls int) {
	r := amount * float64(number_of_walls)
	fmt.Println("Calculating commission of building technician", r)
}

type ElectricTechnician struct {
	amount                  float64
	number_lights_installed int
	calculator              Calculator
}

type BuildingTechnician struct {
	amount          float64
	number_of_walls int
	calculator      Calculator
}

func bridge() {
	fmt.Println("Bridge")
	tipCalculator := &TipCalculator{}
	commissionCalculator := &CommissionCalculator{}

	calculators := []Calculator{tipCalculator, commissionCalculator}
	for _, c := range calculators {
		eTechnician := ElectricTechnician{
			amount:                  10,
			number_lights_installed: 2,
			calculator:              c,
		}
		eTechnician.calculator.calculateElectricTechnician(eTechnician.amount, eTechnician.number_lights_installed)

		bTechnician := BuildingTechnician{
			amount:          10,
			number_of_walls: 2,
			calculator:      c,
		}
		bTechnician.calculator.calculateBuildingTechnician(bTechnician.amount, bTechnician.number_of_walls)

	}

}
