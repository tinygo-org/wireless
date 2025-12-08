//go:build si5351

package main

import (
	"machine"

	"tinygo.org/x/drivers/si5351"
	"tinygo.org/x/wireless/afsk"
)

func initRadio() afsk.Radio {
	dev := si5351.New(machine.I2C0)
	dev.Configure()

	return &Si5351Radio{device: &dev}
}

type Si5351Radio struct {
	device *si5351.Device
}

func (r *Si5351Radio) Transmit(freq uint64) error {
	println("Si5351Radio Transmit called with freq:", freq)
	r.device.SetFrequency(freq, 0, si5351.PLL_A)

	return nil
}

func (r *Si5351Radio) Standby() error {
	println("Si5351Radio Standby called")
	r.device.OutputEnable(0, false)

	return nil
}

func (r *Si5351Radio) Close() error {
	println("Si5351Radio Close called")
	return nil
}
