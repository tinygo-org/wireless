//go:build !si5351 && !featherwing && !pwm

package main

import (
	"tinygo.org/x/wireless/examples/audio"
	"tinygo.org/x/wireless/morse"
)

func initRadio() *morse.Morse {
	player := audio.NewPlayer()

	m := morse.NewMorse(player, 440, 5)
	m.Configure()

	return m
}
