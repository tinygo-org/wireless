# wireless

Wireless communication protocol implementations to be used by TinyGo supported radios.

For example, this program sends WSPR packet data:

```go
package main

import (
	"time"

	"tinygo.org/x/wireless/wspr"
)

func main() {
	println("Starting WSPR communication example...")
	time.Sleep(2 * time.Second)

	println("WSPR modem initialized.")
	radio := initRadio()

	frequency := radio.GetBaseFrequency()
	println("Transmitting on frequency", frequency, "Hz")

	data := make([]byte, 162)

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
	data = data[:n]

	// transmit some data
	for range 5 {
		println("Transmitting WSPR message with", len(data), "symbols")
		radio.WriteSymbols(data)

		time.Sleep(3 * time.Second)
	}

	time.Sleep(2 * time.Second)

	println("Putting radio in standby mode...")
	radio.Standby()
	time.Sleep(1 * time.Second)

	println("WSPR modem example completed.")
	radio.Close()
}
```

## Supported Protocols

### AFSK

Audio Frequency-Shift Keying 

https://notblackmagic.com/bitsnpieces/afsk/

### FSK4

Frequency-shift keying (FSK4)

https://en.wikipedia.org/wiki/Frequency-shift_keying

### WSPR

Weak Signal Propagation Reporter (WSPR)

https://en.wikipedia.org/wiki/WSPR_(amateur_radio_software)

## Supported Radios

- si5351
- SX126X
