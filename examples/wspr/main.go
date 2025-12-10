// WSPR example
//
// tinygo flash -size short -tags=si5351 -target=pico -monitor ./examples/wspr
// tinygo flash -size short -tags=featherwing -target=pybadge -monitor ./examples/wspr
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

	// Example WSPR packet data
	// K1ABC FN42 37
	msg, err := wspr.WsprMessage("K1ABC", "FN42", 37)
	if err != nil {
		println("Error creating WSPR message:", err.Error())
		return
	}

	// transmit some data
	for range 5 {
		println("Transmitting WSPR message with", len(msg), "symbols")
		radio.WriteSymbols(msg)

		time.Sleep(3 * time.Second)
	}

	time.Sleep(2 * time.Second)

	// put the radio in standby
	println("Putting radio in standby mode...")
	radio.Standby()
	time.Sleep(1 * time.Second)

	println("WSPR modem example completed.")
	radio.Close()
}
