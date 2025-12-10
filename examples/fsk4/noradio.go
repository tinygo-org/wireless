//go:build !si5351 && !featherwing

package main

import (
	"tinygo.org/x/wireless/fsk4"
)

func initRadio() *fsk4.FSK4 {
	return fsk4.NewFSK4(&NoRadio{}, 144800000, 5000, 1200)
}

type NoRadio struct{}

func (r *NoRadio) Transmit(freq uint64) error {
	println("NoRadio Transmit called with freq:", freq)
	return nil
}

func (r *NoRadio) Standby() error {
	println("NoRadio Standby called")
	return nil
}

func (r *NoRadio) Close() error {
	println("NoRadio Close called")
	return nil
}
