package main

import (
	"fmt"
	"machine"
	"time"
)

func main() {
	// Set up the LED pins
	red := machine.Pin(18)
	amber := machine.Pin(19)
	green := machine.Pin(20)
	red.Configure(machine.PinConfig{Mode: machine.PinOutput})
	amber.Configure(machine.PinConfig{Mode: machine.PinOutput})
	green.Configure(machine.PinConfig{Mode: machine.PinOutput})

	// Configure ADC pin 26 as ADC1 for the light sensor
	lightSensor := machine.ADC{Pin: machine.ADC1}
	lightSensor.Configure(machine.ADCConfig{})

	for {
		// Read sensor value
		light := lightSensor.Get()

		// Calculate light percentage (rounded to 1 decimal place)
		lightPercent := float64(light) / 65535.0 * 100.0
		fmt.Printf("%.1f%%\n", lightPercent)

		// 1-second delay between readings
		time.Sleep(1 * time.Second)

		// Control LEDs based on light percentage
		if lightPercent <= 30.0 {
			red.High()
			amber.Low()
			green.Low()
		} else if lightPercent > 30.0 && lightPercent < 60.0 {
			red.Low()
			amber.High()
			green.Low()
		} else {
			red.Low()
			amber.Low()
			green.High()
		}
	}
}
