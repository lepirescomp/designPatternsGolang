package main

import "fmt"

type Currency struct {
	symbol      string
	punctuation string
}

type Earning struct {
	currency Currency
	amount   float64
}

type Spec interface {
	Format(earning Earning)
}

type MailWalkingMessage struct {
	mailWalkingMessage string
}

func (p MailWalkingMessage) Format(earning Earning) {
	fmt.Println(p.mailWalkingMessage, earning.currency.symbol, earning.amount)
}

type MailDrivingMessage struct {
	mailDrivingMessage string
}

func (d MailDrivingMessage) Format(earning Earning) {
	fmt.Println(d.mailDrivingMessage, earning.currency.symbol, earning.amount)
}

func showEarningsMessage(spec Spec, earning Earning) {
	spec.Format(earning)
}

type Message struct {
	spec    Spec
	earning Earning
}

func openClosePrinciple() {
	fmt.Println("Open Close Principle")
	earning1 := Earning{Currency{"$", "."}, 10.0}
	earning2 := Earning{Currency{"R$", ","}, 20.0}

	mail_driving_string := MailDrivingMessage{
		mailDrivingMessage: "MailDriverMan earnings:",
	}

	mail_walking_string := MailWalkingMessage{
		mailWalkingMessage: "MailWalking earnings:",
	}

	messages := []Message{
		Message{mail_driving_string, earning1},
		Message{mail_walking_string, earning2},
	}

	for _, v := range messages {
		showEarningsMessage(v.spec, v.earning)

	}
}
