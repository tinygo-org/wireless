// WSPR example
//
// This code must start exactly on an even minute boundary to conform with WSPR timing.
// tinygo flash -size short -tags=si5351 -target=pico2 -monitor .
// tinygo flash -size short -tags=featherwing -target=pybadge -monitor ./examples/wspr
// go run ./examples/wspr
package main

import (
	"time"

	"tinygo.org/x/wireless/wspr"
)

func main() {
	println("Starting WSPR communication example...")

	// init the modem
	println("WSPR modem initialized.")
	radio := initRadio()

	frequency := radio.GetBaseFrequency()
	println("Transmitting on frequency", frequency, "Hz")

	data := make([]byte, 256)

	// Example WSPR packet data
	// K1ABC FN42 37
	// See https://en.wikipedia.org/wiki/WSPR_(amateur_radio_software)
	msg, err := wspr.NewMessage("K1ABC", "FN42", 37)
	if err != nil {
		println("Error creating WSPR message:", err.Error())
		return
	}
	n, err := msg.WriteSymbols(data)
	if err != nil {
		println("error writing WSPR message")
		return
	}

	// transmit some data
	println("Transmitting WSPR message with", n, "symbols")
	if err := radio.WriteSymbols(data[:n]); err != nil {
		println("error transmitting WSPR message:", err.Error())
		return
	}

	// put the radio in standby
	println("Putting radio in standby mode...")
	radio.Standby()
	time.Sleep(1 * time.Second)

	println("WSPR modem example completed.")
	radio.Close()
}
