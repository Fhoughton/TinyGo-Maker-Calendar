package main

import (
	"machine"
	"time"
)

func activity_one() {
	// Configure the potentiometer
	potentiometer := machine.ADC{Pin: machine.ADC1}
	potentiometer.Configure(machine.ADCConfig{})

	for { // Infinite loop
		// Read the potentiometer value (10-bit resolution in TinyGo, ranging from 0 to 1023)
		value := potentiometer.Get()

		// Print the potentiometer value
		println(value)

		time.Sleep(1 * time.Second) // Wait a second
	}
}

func main() {
	activity_one()
}