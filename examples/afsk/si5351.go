//go:build si5351

package main

import (
	"machine"

	"tinygo.org/x/drivers/si5351"
	"tinygo.org/x/wireless/afsk"
)

func initRadio() *afsk.AFSK {
	machine.I2C0.Configure(machine.I2CConfig{})
	dev := si5351.New(machine.I2C0)
	if err := dev.Configure(si5351.Config{}); err != nil {
		panic(err)
	}

	a := afsk.NewAFSK(&Si5351Radio{device: dev})
	a.Configure()

	return a
}

type Si5351Radio struct {
	device *si5351.Device
}

func (r *Si5351Radio) Transmit(freq uint64) error {
	if err := r.device.SetRawFrequency(si5351.Clock0, si5351.Frequency(freq)); err != nil {
		return err
	}

	r.device.EnableOutput(si5351.Clock0, true)

	return nil
}

func (r *Si5351Radio) Standby() error {
	r.device.EnableOutput(si5351.Clock0, false)

	return nil
}

func (r *Si5351Radio) Close() error {
	return nil
}
