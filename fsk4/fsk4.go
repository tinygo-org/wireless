package fsk4

import (
	"math"
	"time"
)

// FSK4 represents an FSK4 modem.
type FSK4 struct {
	radio     Radio
	frequency float64
	shift     uint32
	rate      uint32
	tones     [4]int32
}

// NewFSK4 creates a new FSK4 modem instance.
func NewFSK4(radio Radio, frequency float64, shift, rate uint32) *FSK4 {
	return &FSK4{
		radio:     radio,
		frequency: frequency,
		shift:     shift,
		rate:      rate,
	}
}

// Configure sets up the FSK4 modem parameters.
func (r *FSK4) Configure() error {
	shiftFreq := r.getRawShift(int32(r.shift))
	for i := range 4 {
		r.tones[i] = shiftFreq * int32(i)
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

// Close releases resources associated with the FSK4 modem.
func (r *FSK4) Close() error {
	return r.radio.Close()
}

// GetFrequency returns the current transmission frequency.
func (r *FSK4) GetFrequency() float64 {
	return r.frequency
}

// GetRate returns the current sample rate.
func (r *FSK4) GetRate() uint32 {
	return r.rate
}

// SetFrequency sets the transmission frequency.
func (r *FSK4) SetFrequency(freq float64) {
	r.frequency = freq
}

// SetSampleRate sets the sample rate.
func (r *FSK4) SetSampleRate(rate uint32) {
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
		r.tone(symbol)

		// Shift to next symbol
		data = data << 2
	}

	return nil
}

func (r *FSK4) tone(symbol byte) {
	freq := r.frequency + float64(r.tones[symbol])
	r.radio.Transmit(uint32(freq))
	// hold for one symbol period
	time.Sleep(time.Duration(1000000/r.rate) * time.Microsecond)
}

func (r *FSK4) getRawShift(shift int32) int32 {
	// calculate module carrier frequency resolution
	step := int32(math.Round(r.radio.GetFreqStep()))

	// check minimum shift value
	if abs(shift) < step/2 {
		return 0
	}

	// round shift to multiples of frequency step size
	if abs(shift)%step < (step / 2) {
		return shift / step
	}
	if shift < 0 {
		return (shift / step) - 1
	}
	return (shift / step) + 1
}

func abs(x int32) int32 {
	if x < 0 {
		return -x
	}
	return x
}
