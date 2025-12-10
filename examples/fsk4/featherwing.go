//go:build featherwing

package main

import (
	"machine"

	"tinygo.org/x/drivers/sx127x"
	"tinygo.org/x/wireless/fsk4"
)

var (
	spi    = machine.SPI1
	rstPin = machine.D9
)

func initRadio() *fsk4.FSK4 {
	dev := sx127x.New(spi, rstPin)
	dev.SetOpMode(sx127x.SX127X_OPMODE_SLEEP)
	dev.SetOpModeFsk()
	dev.SetModulationType(sx127x.SX127X_OPMODE_MODULATION_OOK)

	return fsk4.NewFSK4(&sx127xRadio{device: dev}, 10140956, 270, 100)
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
