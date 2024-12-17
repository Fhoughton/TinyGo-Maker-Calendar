package main

import (
	"encoding/hex"
	"fmt"
	"machine"
	"time"

	"tinygo.org/x/drivers/ds18b20"
	"tinygo.org/x/drivers/onewire"
)

func activity_one() {
	// Define the pin connected to the DS18B20 sensor
	pin := machine.Pin(26)
	pin.Configure(machine.PinConfig{Mode: machine.PinInput})

	// Initialize the 1-Wire bus
	ow := onewire.New(pin)

	// Search for connected DS18B20 devices
	romIDs, err := ow.Search(onewire.SEARCH_ROM)
	if err != nil {
		println("Error during ROM search:", err.Error())
		return
	}

	if len(romIDs) == 0 {
		println("No DS18B20 sensors found on the bus.")
		return
	}

	// Initialize the DS18B20 sensor
	sensor := ds18b20.New(ow)

	// Print the ROM IDs of all detected sensors
	for i, romid := range romIDs {
		println("Found sensor", i, "with ROM ID:", hex.EncodeToString(romid))
	}

	// Main loop to read and print temperature values
	for {
		// Request temperature conversion from all connected sensors
		for _, romid := range romIDs {
			println("Requesting temperature for sensor:", hex.EncodeToString(romid))
			sensor.RequestTemperature(romid)
		}

		// Wait for conversion to complete (750ms for 12-bit resolution)
		time.Sleep(750 * time.Millisecond)

		// Read and print the temperature values
		for _, romid := range romIDs {
			println("Reading temperature for sensor:", hex.EncodeToString(romid))
			tempMilliC, err := sensor.ReadTemperature(romid)
			if err != nil {
				println("Error reading temperature:", err.Error())
				continue
			}

			// Convert milli-degrees Celsius to degrees Celsius
			tempC := float64(tempMilliC) / 1000.0
			fmt.Printf("Temperature: %f °C\n", tempC)
		}

		// Delay before the next reading
		time.Sleep(5 * time.Second)
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

	// Define the pin connected to the DS18B20 sensor
	pin := machine.Pin(26)
	pin.Configure(machine.PinConfig{Mode: machine.PinInput})

	// Initialize the 1-Wire bus
	ow := onewire.New(pin)

	// Search for connected DS18B20 devices
	romIDs, err := ow.Search(onewire.SEARCH_ROM)
	if err != nil {
		println("Error during ROM search:", err.Error())
		return
	}

	if len(romIDs) == 0 {
		println("No DS18B20 sensors found on the bus.")
		return
	}

	// Initialize the DS18B20 sensor
	sensor := ds18b20.New(ow)

	// Main loop to read and print temperature values
	for {
		// Request temperature conversion from all connected sensors
		for _, romid := range romIDs {
			sensor.RequestTemperature(romid)
		}

		// Wait for conversion to complete
		time.Sleep(time.Second)

		// Read and process the temperature values
		for _, romid := range romIDs {
			tempMilliC, err := sensor.ReadTemperature(romid)
			if err != nil {
				println("Error reading temperature:", err.Error())
				continue
			}

			// Convert milli-degrees Celsius to degrees Celsius
			tempC := float64(tempMilliC) / 1000.0

			if tempC <= 18 {
				red.High()
				amber.Low()
				green.Low()
			} else if tempC > 18 && tempC < 22 {
				red.Low()
				amber.High()
				green.Low()
			} else if tempC > 22 {
				red.Low()
				amber.Low()
				green.High()
			}
		}

		// Delay before the next reading
		time.Sleep(5 * time.Second)
	}
}

func alarm() {
	// Configure LED pins
	red := machine.Pin(18)
	amber := machine.Pin(19)
	green := machine.Pin(20)
	red.Configure(machine.PinConfig{Mode: machine.PinOutput})
	amber.Configure(machine.PinConfig{Mode: machine.PinOutput})
	green.Configure(machine.PinConfig{Mode: machine.PinOutput})

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
}

func activity_three() {
	// Define the pin connected to the DS18B20 sensor
	pin := machine.Pin(26)
	pin.Configure(machine.PinConfig{Mode: machine.PinInput})

	// Initialize the 1-Wire bus
	ow := onewire.New(pin)

	// Search for connected DS18B20 devices
	romIDs, err := ow.Search(onewire.SEARCH_ROM)
	if err != nil {
		println("Error during ROM search:", err.Error())
		return
	}

	if len(romIDs) == 0 {
		println("No DS18B20 sensors found on the bus.")
		return
	}

	// Initialize the DS18B20 sensor
	sensor := ds18b20.New(ow)

	// Print the ROM IDs of all detected sensors
	for i, romid := range romIDs {
		println("Found sensor", i, "with ROM ID:", hex.EncodeToString(romid))
	}

	// Main loop to read and print temperature values
	for {
		// Request temperature conversion from all connected sensors
		for _, romid := range romIDs {
			sensor.RequestTemperature(romid)
		}

		// Wait for conversion to complete (1 second)
		time.Sleep(1 * time.Second)

		// Read and print the temperature values
		for _, romid := range romIDs {
			tempMilliC, err := sensor.ReadTemperature(romid)
			if err != nil {
				println("Error reading temperature:", err.Error())
				continue
			}

			// Convert milli-degrees Celsius to degrees Celsius
			tempC := float64(tempMilliC) / 1000.0

			if tempC > 17 {
				alarm()
			}
		}

		// Delay before the next reading
		time.Sleep(5 * time.Second)
	}
}

func main() {
	activity_three()
}
