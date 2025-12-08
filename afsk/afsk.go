package afsk

// AFSK represents an AFSK 	.
type AFSK struct {
	radio Radio
}

// NewAFSK creates a new AFSK modem instance.
func NewAFSK(radio Radio) *AFSK {
	return &AFSK{
		radio: radio,
	}
}

// Close releases resources associated with the AFSK modem.
func (r *AFSK) Close() error {
	return r.radio.Close()
}

// Configure sets up the AFSK modem parameters.
func (r *AFSK) Configure() error {
	return nil
}

// SetFrequency sets the transmission frequency.
func (r *AFSK) Tone(freq uint64) {
	r.radio.Transmit(freq)
}

// Standby puts the radio in standby mode.
func (r *AFSK) Standby() error {
	return r.radio.Standby()
}
