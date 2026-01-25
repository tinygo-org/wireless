package afsk

import (
	"testing"
)

// MockRadio implements the Radio interface for testing
type MockRadio struct {
	frequencies []uint64
	standby     bool
	closed      bool
}

func NewMockRadio() *MockRadio {
	return &MockRadio{
		frequencies: make([]uint64, 0),
	}
}

func (m *MockRadio) Transmit(freq uint64) error {
	m.frequencies = append(m.frequencies, freq)
	return nil
}

func (m *MockRadio) Standby() error {
	m.standby = true
	return nil
}

func (m *MockRadio) Close() error {
	m.closed = true
	return nil
}

func TestNewAFSK(t *testing.T) {
	radio := NewMockRadio()
	afsk := NewAFSK(radio)

	if afsk == nil {
		t.Error("NewAFSK() returned nil")
	}
	if afsk.radio != radio {
		t.Error("NewAFSK() did not set radio correctly")
	}
}

func TestConfigure(t *testing.T) {
	radio := NewMockRadio()
	afsk := NewAFSK(radio)

	err := afsk.Configure()
	if err != nil {
		t.Errorf("Configure() error = %v", err)
	}
}

func TestTone(t *testing.T) {
	tests := []struct {
		name     string
		freq     uint64
		expected uint64
	}{
		{"low frequency", 1200, 120000},
		{"high frequency", 2200, 220000},
		{"zero frequency", 0, 0},
		{"fractional frequency rounds down", 1500, 150000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			radio := NewMockRadio()
			afsk := NewAFSK(radio)

			afsk.Tone(tt.freq)

			if len(radio.frequencies) != 1 {
				t.Errorf("Tone() transmitted %d frequencies, want 1", len(radio.frequencies))
				return
			}
			if radio.frequencies[0] != tt.expected {
				t.Errorf("Tone(%d) transmitted %d, want %d", tt.freq, radio.frequencies[0], tt.expected)
			}
		})
	}
}

func TestToneMultipleCalls(t *testing.T) {
	radio := NewMockRadio()
	afsk := NewAFSK(radio)

	// Simulate AFSK modulation with mark and space frequencies
	frequencies := []uint64{1200, 2200, 1200, 1200, 2200}
	for _, freq := range frequencies {
		afsk.Tone(freq)
	}

	if len(radio.frequencies) != len(frequencies) {
		t.Errorf("Tone() transmitted %d frequencies, want %d", len(radio.frequencies), len(frequencies))
	}

	for i, expected := range frequencies {
		if radio.frequencies[i] != uint64(expected)*100 {
			t.Errorf("frequencies[%d] = %d, want %d", i, radio.frequencies[i], uint64(expected)*100)
		}
	}
}

func TestStandby(t *testing.T) {
	radio := NewMockRadio()
	afsk := NewAFSK(radio)

	err := afsk.Standby()
	if err != nil {
		t.Errorf("Standby() error = %v", err)
	}
	if !radio.standby {
		t.Error("Standby() did not put radio in standby mode")
	}
}
