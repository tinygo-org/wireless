//go:build si5351

package main

import (
	"machine"

	"tinygo.org/x/drivers/si5351"
	"tinygo.org/x/wireless/morse"
)

func initRadio() *morse.Morse {
	machine.I2C0.Configure(machine.I2CConfig{})

	dev := si5351.New(machine.I2C0)
	cnf := si5351.Config{}
	if err := dev.Configure(cnf); err != nil {
		panic(err)
	}

	m := morse.NewMorse(&Si5351Radio{device: dev}, 11_400, 20)
	m.Configure()

	return m
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
