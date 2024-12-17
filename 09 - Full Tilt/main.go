package main

import (
	"machine"
	"time"
)

func activity_one() {
	// Set up the tilt sensor pin
	tilt := machine.Pin(26)
	tilt.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})

	for {
		time.Sleep(10 * time.Millisecond) // Short delay

		if tilt.Get() { // If sensor is HIGH
			println("I tilted!") // Print a string
		}
	}
}

func activity_two() {
	// Set up the tilt sensor pin
	tilt := machine.Pin(26)
	tilt.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})

	// Configure the buzzer pin for PWM
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

	// Set the PWM frequency to 1000 Hz (1 kHz)
	err = pwm.SetPeriod(1e9 / 1000) // 1 second divided by 1000 Hz
	if err != nil {
		println("Failed to set PWM period:", err.Error())
		return
	}

	for {
		time.Sleep(10 * time.Millisecond) // Short delay

		if tilt.Get() { // If sensor is HIGH
			println("***TILT DETECTED***") // Print a string

			// Set duty cycle to 10,000 (volume up)
			pwm.Set(channel, 10000)
			time.Sleep(200 * time.Millisecond) // Short delay

			// Set duty cycle to 0 (volume off)
			pwm.Set(channel, 0)
		}
	}
}

func activity_three() {
	// Set up the tilt sensor pin
	tilt := machine.Pin(26)
	tilt.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})

	tiltcount := 0
	state := 0

	for {
		time.Sleep(100 * time.Millisecond) // Short delay

		if state == 0 && tilt.Get() {
			tiltcount++
			state = 1
			println("tilts =", tiltcount)
		}
		if state == 1 && !tilt.Get() {
			state = 0
		}
	}
}

func main() {
	activity_three()
}
