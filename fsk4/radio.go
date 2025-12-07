package fsk4

// Radio defines the interface for FSK4 radio transmitters.
type Radio interface {
	// Transmit sends a signal at the specified frequency.
	Transmit(freq uint32) error
	// Standby puts the radio into standby mode.
	Standby() error
	// Close releases resources associated with the radio.
	Close() error
	// GetFreqStep returns the frequency step size of the radio.
	GetFreqStep() float64
}
