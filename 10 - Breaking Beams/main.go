package main

import (
	"machine"
	"time"
)

func activity_one() {
	// Set up the beam sensor pin
	beam := machine.Pin(26)
	beam.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})

	for {
		time.Sleep(100 * time.Millisecond) // Short delay

		if !beam.Get() { // If sensor is HIGH
			println("Beam Broken!") // Print a string
		}
	}
}

func activity_two() {
	// Set up the beam sensor pin
	beam := machine.Pin(26)
	beam.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})

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

	// Game variables
	scoreCounter := 0
	state := 0
	timeCheck := int64(0)

	time.Sleep(1 * time.Second) // Short delay to enable the output
	println("Game starts after the beep")
	pwm.Set(channel, 2000)
	time.Sleep(2 * time.Second)
	pwm.Set(channel, 0)

	println("GO!")

	startTime := time.Now().UnixMilli()

	for {
		time.Sleep(100 * time.Millisecond) // Short delay

		// Calculate elapsed time in seconds
		timeCheck = (time.Now().UnixMilli() - startTime) / 1000

		// Check if the game is over
		if timeCheck >= 30 {
			println("GAME OVER!")

			// Beep to signal game end
			pwm.Set(channel, 200)
			time.Sleep(100 * time.Millisecond) // Short delay
			pwm.Set(channel, 0)

			// Print the player's score
			println("YOUR SCORE:", scoreCounter)

			// Exit the program
			break
		}

		// Game logic
		if state == 0 && !beam.Get() {
			scoreCounter++
			state = 1

			println("SCORE =", scoreCounter)
			println("Time remaining:", 30-timeCheck)
		} else if state == 1 && beam.Get() {
			state = 0
		}
	}
}

func main() {
	activity_two()
}
