package main

import (
	"machine"
	"time"
)

func activity_one() {
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

	// Set the duty cycle to ~15% (10000 out of 65535)
	pwm.Set(channel, 10000)

	// Wait for 1 second
	time.Sleep(1 * time.Second)

	// Turn off the buzzer by setting the duty cycle to 0
	pwm.Set(channel, 0)
}

func activity_two() {
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

	// Set the duty cycle to ~15% (10000 out of 65535)
	pwm.Set(channel, 10000)

	// Set the PWM frequency to 1000 Hz (1 kHz)
	err = pwm.SetPeriod(1e9 / 1000) // 1 second divided by 1000 Hz
	if err != nil {
		println("Failed to set PWM period:", err.Error())
		return
	}
	// Wait for 1 second
	time.Sleep(1 * time.Second)

	// Set the PWM frequency to 500 Hz (0.5 kHz)
	err = pwm.SetPeriod(1e9 / 500) // 1 second divided by 1000 Hz
	if err != nil {
		println("Failed to set PWM period:", err.Error())
		return
	}
	// Wait for 1 second
	time.Sleep(1 * time.Second)

	// Turn off the buzzer by setting the duty cycle to 0
	pwm.Set(channel, 0)
}

// Note frequencies for "Jingle Bells"
const (
	C = 523 // C note
	D = 587 // D note
	E = 659 // E note
	G = 784 // G note
)

func playTone(frequency int, duration time.Duration, volume uint32) {
	// Configure the buzzer pin for PWM
	buzzerPin := machine.Pin(13)
	pwm := machine.PWM6
	pwm.Configure(machine.PWMConfig{})

	// Get the PWM channel for the buzzer
	channel, err := pwm.Channel(buzzerPin)
	if err != nil {
		println("Failed to get PWM channel:", err.Error())
		return
	}

	// Set the PWM frequency
	err = pwm.SetPeriod(1e9 / uint64(frequency))
	if err != nil {
		println("Failed to set frequency:", err.Error())
		return
	}

	// Set duty cycle for volume
	pwm.Set(channel, volume)

	// Play tone for the duration
	time.Sleep(duration)

	// Turn off the tone
	pwm.Set(channel, 0)
	time.Sleep(200 * time.Millisecond) // Add a delay between notes
}

func activity_five() {
	// Volume (Duty cycle, scaled to 16-bit range)
	volume := uint32(10000)

	// Play the tune
	playTone(E, 100*time.Millisecond, volume) // "Jin..."
	playTone(E, 100*time.Millisecond, volume) // "...gle"
	playTone(E, 100*time.Millisecond, volume) // "Bells"

	time.Sleep(500 * time.Millisecond) // Longer delay

	playTone(E, 100*time.Millisecond, volume) // "Jin..."
	playTone(E, 100*time.Millisecond, volume) // "...gle"
	playTone(E, 100*time.Millisecond, volume) // "Bells"

	time.Sleep(500 * time.Millisecond) // Longer delay

	playTone(E, 100*time.Millisecond, volume) // "Jin..."
	playTone(G, 100*time.Millisecond, volume) // "...gle"
	playTone(C, 100*time.Millisecond, volume) // "All"
	playTone(D, 100*time.Millisecond, volume) // "The"
	playTone(E, 100*time.Millisecond, volume) // "Way"

	// Ensure the buzzer is off at the end
	pwm := machine.PWM6
	pwm.Configure(machine.PWMConfig{})

	buzzerPin := machine.Pin(13)
	channel, err := pwm.Channel(buzzerPin)
	if err != nil {
		println("Failed to get PWM channel:", err.Error())
		return
	}
	pwm.Set(channel, 0)
}

func main() {
	activity_five()
}
