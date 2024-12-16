package main

import (
	"machine"
	"time"
)

func activity_one() {
	// Configure pins for output
	red := machine.Pin(18)
	amber := machine.Pin(19)
	green := machine.Pin(20)

	red.Configure(machine.PinConfig{Mode: machine.PinOutput})
	amber.Configure(machine.PinConfig{Mode: machine.PinOutput})
	green.Configure(machine.PinConfig{Mode: machine.PinOutput})

	// Turn on LEDs
	red.High()
	amber.High()
	green.High()

	// Wait for 5 seconds
	time.Sleep(5 * time.Second)

	// Turn off LEDs
	red.Low()
	amber.Low()
	green.Low()
}

func activity_two() {
	// Configure pins for output
	red := machine.Pin(18)
	amber := machine.Pin(19)
	green := machine.Pin(20)

	red.Configure(machine.PinConfig{Mode: machine.PinOutput})
	amber.Configure(machine.PinConfig{Mode: machine.PinOutput})
	green.Configure(machine.PinConfig{Mode: machine.PinOutput})

	counter := 1 // Set the counter to start at 1

	for counter < 11 { // While counter is less than 11
		println(counter) // Print the current counter

		// LEDs all on
		red.High()
		amber.High()
		green.High()

		time.Sleep(500 * time.Millisecond) // Wait half a second

		// LEDs all off
		red.Low()
		amber.Low()
		green.Low()

		time.Sleep(500 * time.Millisecond) // Wait half a second

		counter++ // Add 1 to the counter
	}
}

func activity_three() {
	// Configure pins for output
	red := machine.Pin(18)
	amber := machine.Pin(19)
	green := machine.Pin(20)

	red.Configure(machine.PinConfig{Mode: machine.PinOutput})
	amber.Configure(machine.PinConfig{Mode: machine.PinOutput})
	green.Configure(machine.PinConfig{Mode: machine.PinOutput})

	counter := 1 // Set the counter to 1

	for counter < 11 { // While counter is less than 11
		println(counter) // Print the current counter

		// Red ON
		red.High()   // ON
		amber.Low()  // OFF
		green.Low()  // OFF
		time.Sleep(500 * time.Millisecond) // Wait half a second

		// Amber ON
		red.Low()    // OFF
		amber.High() // ON
		green.Low()  // OFF
		time.Sleep(500 * time.Millisecond) // Wait half a second

		// Green ON
		red.Low()    // OFF
		amber.Low()  // OFF
		green.High() // ON
		time.Sleep(500 * time.Millisecond) // Wait half a second

		counter++ // Increment the counter
	}
}

func main() {
	activity_three()
}