package morse

// Radio defines the interface for Morse radio transmitters.
type Radio interface {
	// Transmit sends a signal at the specified frequency passed in Hz * 100
	Transmit(freq uint64) error
	// Standby puts the radio into standby mode.
	Standby() error
	// Close releases resources associated with the radio.
	Close() error
}
