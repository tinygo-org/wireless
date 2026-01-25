//go:build featherwing

package main

import (
	"machine"
	"time"

	"tinygo.org/x/drivers/sx127x"
	"tinygo.org/x/wireless/fsk4"
)

var (
	rstPin  = machine.D11
	csPin   = machine.D10
	dio0Pin = machine.D6
	dio1Pin = machine.D9
	spi     = machine.SPI0
)

func initRadio() *fsk4.FSK4 {
	rstPin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	spi.Configure(machine.SPIConfig{Frequency: 500000, Mode: 0})

	time.Sleep(100 * time.Millisecond)

	println("Initializing SX127x FeatherWing...")

	dev := sx127x.New(spi, rstPin)

	println("resetting device")
	dev.Reset()

	println("going to sleep")
	dev.SetOpMode(sx127x.SX127X_OPMODE_SLEEP)

	println("setting FSK mode")
	dev.SetOpModeFsk()

	println("setting OOK modulation")
	dev.SetModulationType(sx127x.SX127X_OPMODE_MODULATION_OOK)

	fsk := fsk4.NewFSK4(&sx127xRadio{device: dev}, 14_097_060, 270, 100*time.Millisecond)
	fsk.Configure()

	return fsk
}

type sx127xRadio struct {
	device *sx127x.Device
}

func (r *sx127xRadio) Transmit(freq uint64) error {
	r.device.SetFrequency(uint32(freq))
	r.device.SetOpMode(sx127x.SX127X_OPMODE_TX)

	return nil
}

func (r *sx127xRadio) Standby() error {
	r.device.SetOpMode(sx127x.SX127X_OPMODE_STANDBY)

	return nil
}

func (r *sx127xRadio) Close() error {
	return nil
}
