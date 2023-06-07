package main

import "fmt"

type currency struct {
	symbol      string
	punctuation string
}

type earning struct {
	currency currency
	amount   float64
}

type MailWalkingMan struct {
	name            string
	picking_address string
	e               earning
}

type MailDriverMan struct {
	name             string
	delivery_address string
	e                earning
}

type format interface {
	format_earnings()
}

func (p MailWalkingMan) format_earnings() {
	fmt.Print("Earnings by delivery walking: ", p.e.currency.symbol, p.e.amount)
}

func (d MailDriverMan) format_earnings() {
	fmt.Print("Earnings by delivery driving: ", d.e.currency.symbol, d.e.amount)
}

func polimorfism() {

	fmt.Println("Polimorfism")
	earning1 := earning{currency{"$", "."}, 1.0}
	earning2 := earning{currency{"R$", ","}, 10.0}

	mWalkingMan := MailWalkingMan{name: "MainWalkingMan1", picking_address: "p.A", e: earning1}
	mDriverMan := MailDriverMan{name: "MailDriverMan1", delivery_address: "p.B", e: earning2}

	personas := []format{mWalkingMan, mDriverMan}

	for _, v := range personas {
		v.format_earnings()
		fmt.Println()
	}
}
