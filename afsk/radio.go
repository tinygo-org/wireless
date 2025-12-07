package afsk

// Radio defines the interface for AFSK radio transmitters.
type Radio interface {
	// Transmit sends a signal at the specified frequency.
	Transmit(freq uint32) error
	// Standby puts the radio into standby mode.
	Standby() error
	// Close releases resources associated with the radio.
	Close() error
}
