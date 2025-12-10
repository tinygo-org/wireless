//go:build featherwing

package main

import (
	"machine"

	"tinygo.org/x/drivers/sx127x"
	"tinygo.org/x/wireless/afsk"
)

var (
	spi    = machine.SPI1
	rstPin = machine.D9
)

func initRadio() *afsk.AFSK {
	dev := sx127x.New(spi, rstPin)
	dev.SetOpMode(sx127x.SX127X_OPMODE_SLEEP)
	dev.SetOpModeFsk()
	dev.SetModulationType(sx127x.SX127X_OPMODE_MODULATION_OOK)

	return afsk.NewAFSK(&sx127xRadio{device: dev})
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
