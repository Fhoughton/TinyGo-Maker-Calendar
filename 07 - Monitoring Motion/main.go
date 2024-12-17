package main

import (
	"machine"
	"time"
)

func activity_one() {
	// Set up PIR pin as an input with a pull-down resistor
	pir := machine.Pin(26)
	pir.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})

	println("Warming up...")

	// Delay to allow the sensor to settle
	time.Sleep(10 * time.Second)

	println("Sensor ready!")

	for {
		// Short delay to reduce unnecessary processing
		time.Sleep(10 * time.Millisecond)

		// Check if PIR sensor detects movement
		if pir.Get() {
			println("I SEE YOU!")

			// Wait 5 seconds before detecting again
			time.Sleep(5 * time.Second)

			println("Sensor active")
		}
	}
}

func activity_two() {
	// Configure LED pins
	red := machine.Pin(18)
	amber := machine.Pin(19)
	green := machine.Pin(20)
	red.Configure(machine.PinConfig{Mode: machine.PinOutput})
	amber.Configure(machine.PinConfig{Mode: machine.PinOutput})
	green.Configure(machine.PinConfig{Mode: machine.PinOutput})

	// Configure PIR pin
	pir := machine.Pin(26)
	pir.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})

	// Configure Buzzer pin as PWM
	buzzer := machine.Pin(13)
	pwm := machine.PWM6

	// Initialize the PWM peripheral
	pwm.Configure(machine.PWMConfig{})

	// Get a channel for the buzzer pin
	channel, err := pwm.Channel(buzzer)
	if err != nil {
		println("Failed to get PWM channel:", err.Error())
		return
	}

	// Warm-up phase
	println("Warming up...")
	time.Sleep(10 * time.Second)
	println("Sensor ready!")

	for {
		time.Sleep(10 * time.Millisecond) // Short delay

		// Check PIR sensor
		if pir.Get() {
			println("I SEE YOU!")

			// Trigger alarm
			// Set buzzer duty cycle (volume on)
			pwm.Set(channel, 10000)

			for i := 0; i < 5; i++ {
				// Higher pitch
				pwm.SetPeriod(1e9 / 5000) // 5000 Hz

				// Turn on LEDs
				red.High()
				amber.High()
				green.High()

				time.Sleep(1 * time.Second)

				// Lower pitch
				pwm.SetPeriod(1e9 / 500) // 500 Hz

				// Turn off LEDs
				red.Low()
				amber.Low()
				green.Low()

				time.Sleep(1 * time.Second)
			}

			// Set buzzer duty cycle (volume off)
			pwm.Set(channel, 0)

			println("Sensor active")
		}
	}
}

func main() {
	activity_two()
}
