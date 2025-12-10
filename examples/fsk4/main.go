// FSK4 modem example using Si5351 as the frequency generator.
//
// tinygo flash -size short -tags=si5351 -target=pico -monitor ./examples/fsk4
// tinygo flash -size short -tags=featherwing -target=pybadge -monitor ./examples/fsk4
package main

import (
	"encoding/hex"
	"time"
)

// An encoded Horus Binary telemetry packet.
// Refer here for packet format information:
// https://github.com/projecthorus/horusdemodlib/wiki/2---Modem-Details#horus-binary-v1-mode-4-fsk
// After demodulation, deinterleaving, and descrambling, this results in a packet:
// 00000001172D0000000000000000D20463010AFF2780
// This decodes to the Habitat-compatible telemetry string:
// $$4FSKTEST,0,01:23:45,0.00000,0.00000,1234,99,1,10,5.00*ABCD
var horusPacket = []byte{
	0x45, 0x24, 0x24, 0x48, 0x2F, 0x12, 0x16, 0x08, 0x15, 0xC1,
	0x49, 0xB2, 0x06, 0xFC, 0x92, 0xEB, 0x93, 0xD7, 0xEE, 0x5D,
	0x35, 0xA0, 0x91, 0xDA, 0x8D, 0x5F, 0x85, 0x6B, 0x63, 0x03,
	0x6B, 0x60, 0xEA, 0xFE, 0x55, 0x9D, 0xF1, 0xAB, 0xE5, 0x5E,
	0xDB, 0x7C, 0xDB, 0x21, 0x5A, 0x19,
}

func main() {
	println("Starting FSK4 modem example...")
	time.Sleep(2 * time.Second)

	// init the modem
	println("FSK4 modem initialized.")
	radio := initRadio()

	frequency := radio.GetBaseFrequency()
	println("Transmitting on frequency", frequency, "Hz")

	// transmit some data
	for range 5 {
		println("Transmitting data:", hex.EncodeToString(horusPacket))
		radio.Write(horusPacket)

		time.Sleep(3 * time.Second)
	}

	time.Sleep(2 * time.Second)

	// put the radio in standby
	println("Putting radio in standby mode...")
	radio.Standby()
	time.Sleep(1 * time.Second)

	println("FSK4 modem example completed.")
	radio.Close()
}
