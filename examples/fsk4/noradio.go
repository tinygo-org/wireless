//go:build !si5351 && !featherwing

package main

import (
	"time"

	"tinygo.org/x/wireless/examples/audio"
	"tinygo.org/x/wireless/fsk4"
)

func initRadio() *fsk4.FSK4 {
	player := audio.NewPlayer()

	fsk := fsk4.NewFSK4(player, 440, 22000, 100*time.Millisecond)
	fsk.Configure()

	return fsk
}
