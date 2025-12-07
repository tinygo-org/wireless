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

// Configure sets up the AFSK modem parameters.
func (r *AFSK) Configure() error {
	return nil
}

// SetFrequency sets the transmission frequency.
func (r *AFSK) Tone(freq float64) {
	r.radio.Transmit(uint32(freq))
}

// Standby puts the radio in standby mode.
func (r *AFSK) Standby() error {
	return r.radio.Standby()
}
