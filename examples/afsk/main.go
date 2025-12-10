// AFSK modem example using Si5351 as the frequency generator.
//
// tinygo flash -size short -tags=si5351 -target=pico -monitor ./examples/afsk
// tinygo flash -size short -tags=featherwing -target=pybadge -monitor ./examples/afsk
package main

import (
	"time"
)

func main() {
	println("Starting AFSK modem example...")
	time.Sleep(2 * time.Second)

	// init the modem
	println("AFSK modem initialized.")
	radio := initRadio()

	// set tone frequency 1
	frequency := 1200 // Example frequency in Hz
	println("Setting tone frequency to", frequency, "Hz")
	radio.Tone(uint64(frequency))

	time.Sleep(2 * time.Second)

	// set tone frequency 2
	frequency2 := 2200 // Example frequency in Hz
	println("Setting tone frequency to", frequency2, "Hz")
	radio.Tone(uint64(frequency2))

	time.Sleep(2 * time.Second)

	// put the radio in standby
	println("Putting radio in standby mode...")
	radio.Standby()
	time.Sleep(1 * time.Second)

	println("AFSK modem example completed.")
	radio.Close()
}
