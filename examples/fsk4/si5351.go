//go:build si5351

package main

import (
	"machine"

	"tinygo.org/x/drivers/si5351"
	"tinygo.org/x/wireless/fsk4"
)

func initRadio() *fsk4.FSK4 {
	dev := si5351.New(machine.I2C0)
	dev.Configure()

	f := fsk4.NewFSK4(&Si5351Radio{device: &dev}, 10140956, 270, 100)
	f.Configure()

	return f
}

type Si5351Radio struct {
	device *si5351.Device
}

func (r *Si5351Radio) Transmit(freq uint64) error {
	r.device.SetFrequency(freq, 0, si5351.PLL_A)

	return nil
}

func (r *Si5351Radio) Standby() error {
	r.device.OutputEnable(0, false)

	return nil
}

func (r *Si5351Radio) Close() error {
	return nil
}
