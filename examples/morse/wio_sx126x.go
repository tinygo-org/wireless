//go:build nicenano && wio_sx1262

// This is a Morse code modem example using Wio SX1262 module as the radio transmitter.
// Build this example for NiceNano with Wio SX1262 module with command like this:
// tinygo build -target nicenano -tags wio_sx1262
// For different targets, adjust the pin variables in this file accordingly.

package main

import (
	"machine"
	"time"

	"tinygo.org/x/drivers/sx126x"
	"tinygo.org/x/wireless/morse"
)

const freq = 432_300_000

var (
	spi     = machine.SPI0
	rstPin  = machine.D009
	misoPin = machine.D002
	mosiPin = machine.D115
	sckPin  = machine.D111
	nssPin  = machine.D113
	busyPin = machine.D029
	dio1Pin = machine.D010
	rfSw    = machine.D017
)

func initRadio() *morse.Morse {
	time.Sleep(3 * time.Second)

	rstPin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	rfSw.Configure(machine.PinConfig{Mode: machine.PinOutput})

	err := spi.Configure(machine.SPIConfig{
		Frequency: 500_000,
		Mode:      0,
		SCK:       sckPin,
		SDO:       mosiPin,
		SDI:       misoPin,
	})
	if err != nil {
		panic(err)
	}

	println("Initializing SX126x...")
	dev := sx126x.New(spi, rstPin)
	rc := sx126x.NewRadioControlWio(nssPin, busyPin, dio1Pin, rfSw)
	dev.SetRadioController(rc)

	println("Resetting device")
	dev.Reset()
	time.Sleep(100 * time.Millisecond)

	if !dev.DetectDevice() {
		for {
			println("SX126x device not found!")
			time.Sleep(5 * time.Second)
		}
	}
	println("SX126x device detected")

	// Wio specific configuration
	dev.SetDio3AsTcxoCtrl(sx126x.SX126X_DIO3_OUTPUT_1_8, 5*time.Millisecond)
	dev.SetRegulatorMode(sx126x.SX126X_REGULATOR_DC_DC)
	dev.SetDeviceType(sx126x.DEVICE_TYPE_SX1262)
	dev.Calibrate(sx126x.SX126X_CALIBRATE_ALL)

	dev.SetPacketType(sx126x.SX126X_PACKET_TYPE_GFSK)
	dev.SetTxParams(0, sx126x.SX126X_PA_RAMP_200U)
	dev.SetRfFrequency(freq)
	dev.CalibrateImage(freq)

	dev.SetStandby()

	m := morse.NewMorse(&sx126xRadio{device: dev}, freq, 20)
	_ = m.Configure()

	return m
}

type sx126xRadio struct {
	device       *sx126x.Device
	transmitting bool
}

func (r *sx126xRadio) Transmit(freq uint64) error {
	// Put the radio to XOSC mode before first transmission
	// Without this, the first tone after standby is not transmitted
	if !r.transmitting {
		r.device.ExecSetCommand(sx126x.SX126X_CMD_SET_STANDBY, []uint8{sx126x.SX126X_STANDBY_XOSC})
	}
	r.transmitting = true

	freqHz := freq / 100 // freq is in hundredths of Hz
	r.device.SetRfFrequency(uint32(freqHz))
	r.device.SetTxContinuousWave()

	return nil
}

func (r *sx126xRadio) Standby() error {
	r.transmitting = false
	r.device.SetStandby()
	return nil
}

func (r *sx126xRadio) Close() error {
	return r.Standby()
}
