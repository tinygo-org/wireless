//go:build pwm

package main

import (
	"machine"
	"time"

	"github.com/chewxy/math32"
	"tinygo.org/x/wireless/morse"
)

var (
	pwm                = machine.PWM0
	pin                = machine.GPIO16
	sampleRate float32 = 11_025.0 // suitable for AM radio
)

func initRadio() *morse.Morse {
	err := pwm.Configure(machine.PWMConfig{
		Period: 1854, // 1854 ns, approx 540 kHz
	})

	transmitChannel, err := pwm.Channel(pin)
	if err != nil {
		println("failed to configure channel")
		return nil
	}

	samples := make([]uint32, 512)
	generateSineWave(samples, 440, sampleRate)

	m := morse.NewMorse(&PinRadio{transmitChannel: transmitChannel, samples: samples}, 540_000, 5)
	m.Configure()

	return m
}

// PinRadio implements a simple radio using PWM on a GPIO pin.
//
// It uses PWM to generate an audio tone that is transmitted using amplitude
// modulation (AM).
//
// NOTE: You should use a low-pass filter connected to the output pin to get
// a cleaner signal and avoid harmonics that will cause interference for both
// yourself and others.
//
// See https://github.com/tudbut/picoAM for the original inspiration.
type PinRadio struct {
	transmitChannel uint8
	samples         []uint32
	stopChan        chan struct{}
}

func (r *PinRadio) Transmit(freq uint64) error {
	r.stop()
	r.stopChan = make(chan struct{})

	go func(stop <-chan struct{}) {
		for {
			select {
			case <-stop:
				return
			default:
			}
			for _, v := range r.samples {
				pwm.Set(r.transmitChannel, v)
				time.Sleep(time.Duration(1e9/sampleRate) * time.Nanosecond)
			}
		}
	}(r.stopChan)

	return nil
}

func (r *PinRadio) Standby() error {
	r.stop()
	pwm.Set(r.transmitChannel, 0)

	return nil
}

func (r *PinRadio) Close() error {
	r.stop()
	pwm.Set(r.transmitChannel, 0)

	return nil
}

func (r *PinRadio) stop() {
	if r.stopChan != nil {
		close(r.stopChan)
		r.stopChan = nil
	}
}

func generateSineWave(samples []uint32, freq uint64, sampleRate float32) {
	top := pwm.Top()
	min := top / 4
	max := top

	for i := range samples {
		v := math32.Sin(2.0 * math32.Pi * float32(i) * float32(freq) / sampleRate)
		v = (v + 1.0) / 2.0 // shift to [0, 1]
		samples[i] = uint32(float32(min) + v*float32(max-min))
	}
}
