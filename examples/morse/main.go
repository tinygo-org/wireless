// FSK4 modem example
//
// tinygo flash -size short -tags=si5351 -target=pico -monitor ./examples/morse
// tinygo flash -size short -tags=featherwing -target=pybadge -monitor ./examples/morse
// go run ./examples/morse
package main

import (
	"time"
)

var message = "Hello TinyGo"

func main() {
	println("Starting Morse code modem example...")
	time.Sleep(2 * time.Second)

	// init the modem
	println("Morse code modem initialized.")
	radio := initRadio()

	frequency := radio.GetBaseFrequency()
	println("Transmitting on frequency", frequency, "Hz")

	// transmit some data
	for range 50 {
		println("Transmitting data:", message)
		radio.Write(message)

		time.Sleep(10 * time.Second)
	}

	time.Sleep(2 * time.Second)

	println("Morse code modem example completed.")
	radio.Close()
}
