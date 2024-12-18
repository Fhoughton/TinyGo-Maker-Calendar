package main

import (
	"machine"
	"time"
)

func activity_one() {
	// Configure the button as input with pull-down resistor
	button1 := machine.Pin(13)
	button1.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})

	for { // Infinite loop
		time.Sleep(200 * time.Millisecond) // Short delay

		if button1.Get() { // If button1 is pressed (value is 1)
			println("Button 1 pressed")
		}
	}
}

func activity_two() {
	// Configure the buttons as inputs with pull-down resistors
	button1 := machine.Pin(13)
	button2 := machine.Pin(8)
	button3 := machine.Pin(3)

	button1.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	button2.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	button3.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})

	for { // Infinite loop
		time.Sleep(200 * time.Millisecond) // Short delay

		if button1.Get() { // If button 1 is pressed
			println("Button 1 pressed")
		}

		if button2.Get() { // If button 2 is pressed
			println("Button 2 pressed")
		}

		if button3.Get() { // If button 3 is pressed
			println("Button 3 pressed")
		}
	}
}

func activity_three() {
	// Configure the buttons as inputs with pull-down resistors
	button1 := machine.Pin(13)
	button2 := machine.Pin(8)
	button3 := machine.Pin(3)

	button1.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	button2.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	button3.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})

	// Configure the LEDs as outputs
	red := machine.Pin(18)
	amber := machine.Pin(19)
	green := machine.Pin(20)

	red.Configure(machine.PinConfig{Mode: machine.PinOutput})
	amber.Configure(machine.PinConfig{Mode: machine.PinOutput})
	green.Configure(machine.PinConfig{Mode: machine.PinOutput})

	for { // Infinite loop
		time.Sleep(200 * time.Millisecond) // Short delay

		if button1.Get() && button2.Get() { // If button 1 and button 2 are pressed
			println("Buttons 1 and 2 pressed")
			green.High() // Green LED on
			red.Low()    // Red LED off

		} else if button1.Get() { // If button 1 is pressed
			println("Button 1 pressed")
			amber.High() // Amber LED on
			red.Low()    // Red LED off

		} else { // If no buttons are being pressed
			red.High()  // Red LED on
			amber.Low() // Amber LED off
			green.Low() // Green LED off
		}
	}
}

func activity_four() {
	// Configure the buttons as inputs with pull-down resistors
	button1 := machine.Pin(13)
	button2 := machine.Pin(8)

	button1.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	button2.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})

	// Configure the green LED as an output
	green := machine.Pin(20)
	green.Configure(machine.PinConfig{Mode: machine.PinOutput})

	for { // Infinite loop
		time.Sleep(200 * time.Millisecond) // Short delay

		if button1.Get() || button2.Get() { // If button 1 OR button 2 is pressed
			println("Button 1 or 2 pressed")

			green.High()                // Green LED on
			time.Sleep(2 * time.Second) // Wait for 2 seconds
			green.Low()                 // Green LED off
		}
	}
}

func activity_five() {
	// Configure the button as an input with a pull-down resistor
	button1 := machine.Pin(13)
	button1.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})

	// Configure the red LED as an output
	red := machine.Pin(18)
	red.Configure(machine.PinConfig{Mode: machine.PinOutput})

	for { // Infinite loop
		time.Sleep(500 * time.Millisecond) // Short delay

		if button1.Get() { // If button 1 is pressed
			println("Button 1 pressed")
			red.Set(!red.Get()) // Toggle Red LED on/off
		}
	}
}

func main() {
	activity_five()
}
