//go:build !si5351 && !featherwing

package main

import (
	"tinygo.org/x/wireless/afsk"
	"tinygo.org/x/wireless/examples/audio"
)

func initRadio() *afsk.AFSK {
	player := audio.NewPlayer()
	a := afsk.NewAFSK(player)
	a.Configure()

	return a
}
