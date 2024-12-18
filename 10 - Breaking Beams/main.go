package main

import (
	"fmt"
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

func activity_three() {
	// Set up LED pins
	red := machine.Pin(18)
	amber := machine.Pin(19)
	green := machine.Pin(20)

	red.Configure(machine.PinConfig{Mode: machine.PinOutput})
	amber.Configure(machine.PinConfig{Mode: machine.PinOutput})
	green.Configure(machine.PinConfig{Mode: machine.PinOutput})

	// Set up the Break Beam pin
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
	err = pwm.SetPeriod(1e9 / 1000)
	if err != nil {
		println("Failed to set PWM period:", err.Error())
		return
	}

	// Game variables
	scoreCounter := 0
	state := 0
	timeCheck := int64(0)
	targetScore := 100

	// Start the game
	println("Game starts after the beep!")
	pwm.Set(channel, 10000)
	time.Sleep(2 * time.Second)
	pwm.Set(channel, 0)
	println("GO!")
	println("-------------------------------")

	startTime := time.Now().UnixMilli()

	for {
		// Calculate elapsed time in seconds
		timeCheck = (time.Now().UnixMilli() - startTime) / 1000

		// Check if the game is over (30-second limit)
		if timeCheck >= 30 {
			// Turn off LEDs
			red.Low()
			amber.Low()
			green.Low()

			// Beep to signal game end
			pwm.Set(channel, 10000)
			time.Sleep(200 * time.Millisecond)
			pwm.Set(channel, 0)

			// Print results
			println("-------------------------------")
			println("GAME OVER! YOU LOSE :(")
			fmt.Printf("The target was %d, you scored %d\n", targetScore, scoreCounter)
			println("-------------------------------")
			break
		}

		// Check if the player has reached the target score
		if scoreCounter >= targetScore {
			// Turn off LEDs
			red.Low()
			amber.Low()
			green.Low()

			// Beep to signal game end
			pwm.Set(channel, 10000)
			time.Sleep(200 * time.Millisecond)
			pwm.Set(channel, 0)

			// Print results
			println("-------------------------------")
			println("YOU WIN!")
			fmt.Printf("You took %d seconds!\n", timeCheck)
			println("-------------------------------")
			break
		}

		// Game logic
		if state == 0 && !beam.Get() { // Beam is broken
			scoreCounter++
			state = 1

			fmt.Printf("SCORE = %d / %d\n", scoreCounter, targetScore)
			fmt.Printf("Time remaining: %d seconds\n", 30-timeCheck)

			// Update LED status based on score
			if scoreCounter < targetScore*33/100 { // Score < 33%
				red.High()
				amber.Low()
				green.Low()
			} else if scoreCounter < targetScore*66/100 { // 33% <= Score < 66%
				red.High()
				amber.High()
				green.Low()
			} else { // Score >= 66%
				red.High()
				amber.High()
				green.High()
			}
		} else if state == 1 && beam.Get() { // Beam is unbroken
			state = 0
		}

		// Short delay to avoid busy-waiting
		time.Sleep(100 * time.Microsecond)
	}
}

func main() {
	activity_three()
}
