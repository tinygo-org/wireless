package fsk4

import (
	"testing"
)

// MockRadio implements the Radio interface for testing
type MockRadio struct {
	frequencies []uint64
	freqStep    uint64
	standby     bool
	closed      bool
}

func NewMockRadio(freqStep uint64) *MockRadio {
	return &MockRadio{
		frequencies: make([]uint64, 0),
		freqStep:    freqStep,
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

func (m *MockRadio) GetFreqStep() uint64 {
	return m.freqStep
}

func TestNewFSK4(t *testing.T) {
	radio := NewMockRadio(61.0)
	fsk := NewFSK4(radio, 433000000, 270, 100)

	if fsk.base != 433000000 {
		t.Errorf("base = %d, want 433000000", fsk.base)
	}
	if fsk.shift != 270 {
		t.Errorf("shift = %d, want 270", fsk.shift)
	}
	if fsk.rate != 100 {
		t.Errorf("rate = %d, want 100", fsk.rate)
	}
}

func TestConfigure(t *testing.T) {
	radio := NewMockRadio(61.0)
	fsk := NewFSK4(radio, 433000000, 270, 100)

	err := fsk.Configure()
	if err != nil {
		t.Errorf("Configure() error = %v", err)
	}

	// Check that tones are configured correctly
	expectedTones := [4]uint32{0, 270, 540, 810}
	for i, tone := range fsk.tones {
		if tone != expectedTones[i] {
			t.Errorf("tones[%d] = %d, want %d", i, tone, expectedTones[i])
		}
	}
}

func TestGettersAndSetters(t *testing.T) {
	radio := NewMockRadio(61.0)
	fsk := NewFSK4(radio, 433000000, 270, 100)

	// Test GetBaseFrequency
	if fsk.GetBaseFrequency() != 433000000 {
		t.Errorf("GetBaseFrequency() = %d, want 433000000", fsk.GetBaseFrequency())
	}

	// Test SetBaseFrequency
	fsk.SetBaseFrequency(144000000)
	if fsk.GetBaseFrequency() != 144000000 {
		t.Errorf("after SetBaseFrequency(), GetBaseFrequency() = %d, want 144000000", fsk.GetBaseFrequency())
	}

	// Test GetRate
	if fsk.GetRate() != 100 {
		t.Errorf("GetRate() = %d, want 100", fsk.GetRate())
	}

	// Test SetSampleRate
	fsk.SetSampleRate(200)
	if fsk.GetRate() != 200 {
		t.Errorf("after SetSampleRate(), GetRate() = %d, want 200", fsk.GetRate())
	}

	// Test SetShift
	fsk.SetShift(500)
	if fsk.shift != 500 {
		t.Errorf("after SetShift(), shift = %d, want 500", fsk.shift)
	}
}

func TestClose(t *testing.T) {
	radio := NewMockRadio(61.0)
	fsk := NewFSK4(radio, 433000000, 270, 100)

	err := fsk.Close()
	if err != nil {
		t.Errorf("Close() error = %v", err)
	}
	if !radio.closed {
		t.Error("Close() did not close the radio")
	}
}

func TestStandby(t *testing.T) {
	radio := NewMockRadio(61.0)
	fsk := NewFSK4(radio, 433000000, 270, 100)

	err := fsk.Standby()
	if err != nil {
		t.Errorf("Standby() error = %v", err)
	}
	if !radio.standby {
		t.Error("Standby() did not put radio in standby mode")
	}
}

func TestWriteByte(t *testing.T) {
	radio := NewMockRadio(61.0)
	fsk := NewFSK4(radio, 433000000, 270, 1000000) // High rate for fast test
	fsk.Configure()

	// Write a byte and check that 4 symbols were transmitted
	err := fsk.writeByte(0b11_10_01_00) // symbols: 3, 2, 1, 0
	if err != nil {
		t.Errorf("writeByte() error = %v", err)
	}

	if len(radio.frequencies) != 4 {
		t.Errorf("writeByte() transmitted %d symbols, want 4", len(radio.frequencies))
	}
}

func TestWrite(t *testing.T) {
	radio := NewMockRadio(61.0)
	fsk := NewFSK4(radio, 433000000, 270, 1000000) // High rate for fast test
	fsk.Configure()

	data := []byte{0xAB, 0xCD}
	n, err := fsk.Write(data)
	if err != nil {
		t.Errorf("Write() error = %v", err)
	}
	if n != 2 {
		t.Errorf("Write() returned %d, want 2", n)
	}

	// 2 bytes = 8 symbols
	if len(radio.frequencies) != 8 {
		t.Errorf("Write() transmitted %d symbols, want 8", len(radio.frequencies))
	}

	// Radio should be in standby after write
	if !radio.standby {
		t.Error("Write() did not put radio in standby mode")
	}
}

func TestSymbolExtraction(t *testing.T) {
	// Test that symbols are extracted MSB first
	tests := []struct {
		input   byte
		symbols []byte
	}{
		{0b00_00_00_00, []byte{0, 0, 0, 0}},
		{0b11_11_11_11, []byte{3, 3, 3, 3}},
		{0b11_10_01_00, []byte{3, 2, 1, 0}},
		{0b00_01_10_11, []byte{0, 1, 2, 3}},
		{0b10_10_10_10, []byte{2, 2, 2, 2}},
	}

	for _, tt := range tests {
		radio := NewMockRadio(1.0) // Step of 1 for easy calculation
		fsk := NewFSK4(radio, 1000, 1, 1000000)
		fsk.Configure()

		fsk.writeByte(tt.input)

		for i, expectedSymbol := range tt.symbols {
			// With base freq 1000 and tones [0, 1, 2, 3],
			// transmitted freq should be 1000 + symbol
			expectedFreq := uint32(1000 + uint32(expectedSymbol)*fsk.tones[1]/1)
			if i < len(radio.frequencies) {
				// Verify the correct symbol was selected by checking relative frequencies
				if radio.frequencies[i] != uint64(1000)+uint64(fsk.tones[expectedSymbol]) {
					t.Errorf("byte 0x%02X symbol %d: got freq %d, want %d",
						tt.input, i, radio.frequencies[i], expectedFreq)
				}
			}
		}
	}
}
