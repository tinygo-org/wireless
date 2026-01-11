package fsk4

import (
	"errors"
	"time"
)

var (
	baudRateTooHighError = errors.New("Baud rate too high, cannot keep up with transmission")
)

// FSK4 represents an FSK4 modem.
type FSK4 struct {
	radio Radio
	base  uint64
	shift uint32
	rate  time.Duration
	tones [4]uint32
}

// NewFSK4 creates a new FSK4 modem instance.
// radio: the Radio interface implementation
// base: the base frequency in Hz
// shift: the frequency shift in Hz*100 eg. 270 = 2.7 Hz
// rate: the send rate
func NewFSK4(radio Radio, base uint64, shift uint32, rate time.Duration) *FSK4 {
	return &FSK4{
		radio: radio,
		base:  base,
		shift: shift,
		rate:  rate,
	}
}

// Close releases resources associated with the FSK4 modem.
func (r *FSK4) Close() error {
	return r.radio.Close()
}

// Configure sets up the FSK4 modem parameters.
func (r *FSK4) Configure() error {
	for i := range 4 {
		r.tones[i] = r.shift * uint32(i)
	}

	return nil
}

// Write sends data using FSK4 modulation.
func (r *FSK4) Write(data []byte) (int, error) {
	err := r.write(data)
	if err != nil {
		return 0, err
	}
	return len(data), nil
}

// GetFrequency returns the current transmission frequency.
func (r *FSK4) GetBaseFrequency() uint64 {
	return r.base
}

// GetRate returns the current transfer rate.
func (r *FSK4) GetRate() time.Duration {
	return r.rate
}

// SetBaseFrequency sets the base transmission frequency.
func (r *FSK4) SetBaseFrequency(freq uint64) {
	r.base = freq
}

// SetSampleRate sets the transfer rate.
func (r *FSK4) SetRate(rate time.Duration) {
	r.rate = rate
}

// SetShift sets the frequency shift.
func (r *FSK4) SetShift(shift uint32) {
	r.shift = shift
}

// Standby puts the radio in standby mode.
func (r *FSK4) Standby() error {
	return r.radio.Standby()
}

// WriteSymbols sends FSK4 symbols directly (values 0-3).
// This is useful for protocols like WSPR that provide pre-encoded symbols.
func (r *FSK4) WriteSymbols(symbols []byte) error {
	for _, symbol := range symbols {
		if err := r.tone(symbol & 0x03); err != nil {
			return err
		}
	}

	return r.Standby()
}

func SetCorrection(offsets [4]int, length float32) {
	// Placeholder implementation
}

func (r *FSK4) write(data []byte) error {
	for _, b := range data {
		if err := r.writeByte(b); err != nil {
			return err
		}
	}

	return r.Standby()
}

func (r *FSK4) writeByte(data byte) error {
	// send symbols MSB first
	for i := 0; i < 4; i++ {
		// Extract 4FSK symbol (2 bits)
		symbol := (data & 0xC0) >> 6

		// Modulate
		if err := r.tone(symbol); err != nil {
			return err
		}

		// Shift to next symbol
		data = data << 2
	}

	return nil
}

func (r *FSK4) tone(symbol byte) error {
	start := time.Now()
	freq := r.base*100 + uint64(r.tones[symbol])
	if err := r.radio.Transmit(freq); err != nil {
		return err
	}

	// added delay to account for transmitter startup time
	transmitStartupTime := time.Since(start)
	if transmitStartupTime > r.rate {
		return baudRateTooHighError
	}
	// hold for one symbol period
	time.Sleep(r.rate - transmitStartupTime)

	return nil
}
