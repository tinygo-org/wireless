// WSPR example
//
// tinygo flash -size short -tags=si5351 -target=pico -monitor ./examples/wspr
// tinygo flash -size short -tags=featherwing -target=pybadge -monitor ./examples/wspr
// go run ./examples/wspr
package main

import (
	"time"

	"tinygo.org/x/wireless/wspr"
)

func main() {
	println("Starting WSPR communication example...")
	time.Sleep(2 * time.Second)

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
	for range 50 {
		println("Transmitting WSPR message with", n, "symbols")
		if err := radio.WriteSymbols(data[:n]); err != nil {
			println("error transmitting WSPR message:", err.Error())
			return
		}

		println("Waiting for next transmission...")
		time.Sleep(15 * time.Second)
	}

	time.Sleep(2 * time.Second)

	// put the radio in standby
	println("Putting radio in standby mode...")
	radio.Standby()
	time.Sleep(1 * time.Second)

	println("WSPR modem example completed.")
	radio.Close()
}
