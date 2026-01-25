package morse

import (
	"errors"
	"testing"
	"time"
)

// mockRadio implements the Radio interface for testing.
type mockRadio struct {
	transmitted []uint64
	standbyCnt  int
	closed      bool
	failTx      bool
	failStby    bool
	failClose   bool
}

func (m *mockRadio) Transmit(freq uint64) error {
	if m.failTx {
		return errors.New("transmit error")
	}
	m.transmitted = append(m.transmitted, freq)
	return nil
}
func (m *mockRadio) Standby() error {
	if m.failStby {
		return errors.New("standby error")
	}
	m.standbyCnt++
	return nil
}
func (m *mockRadio) Close() error {
	if m.failClose {
		return errors.New("close error")
	}
	m.closed = true
	return nil
}

func TestNewMorse(t *testing.T) {
	r := &mockRadio{}
	m := NewMorse(r, 12345, 20)
	if m.radio != r {
		t.Error("radio not set correctly")
	}
	if m.base != 12345 {
		t.Error("base not set correctly")
	}
	if m.speed != 20 {
		t.Error("speed not set correctly")
	}
}

func TestConfigure(t *testing.T) {
	m := NewMorse(&mockRadio{}, 0, 20)
	err := m.Configure()
	if err != nil {
		t.Fatalf("Configure failed: %v", err)
	}
	if m.dotLength != 60 {
		t.Errorf("dotLength expected 60, got %d", m.dotLength)
	}
	if m.dashLength != 180 {
		t.Errorf("dashLength expected 180, got %d", m.dashLength)
	}
	if m.letterSpace != 180 {
		t.Errorf("letterSpace expected 180, got %d", m.letterSpace)
	}
	if m.wordSpace != 240 {
		t.Errorf("wordSpace expected 240, got %d", m.wordSpace)
	}
}

func TestGetters(t *testing.T) {
	m := NewMorse(&mockRadio{}, 555, 13)
	if m.GetBaseFrequency() != 555 {
		t.Error("GetBaseFrequency incorrect")
	}
	if m.GetSpeed() != 13 {
		t.Error("GetSpeed incorrect")
	}
}

func TestClose(t *testing.T) {
	r := &mockRadio{}
	m := NewMorse(r, 0, 10)
	err := m.Close()
	if err != nil {
		t.Errorf("Close returned error: %v", err)
	}
	if !r.closed {
		t.Error("radio not closed")
	}
}

func TestCloseError(t *testing.T) {
	r := &mockRadio{failClose: true}
	m := NewMorse(r, 0, 10)
	err := m.Close()
	if err == nil {
		t.Error("expected error from Close")
	}
}

func TestWriteUnsupportedChar(t *testing.T) {
	m := NewMorse(&mockRadio{}, 0, 10)
	m.Configure()
	_, err := m.Write("\x19") // below ' '
	if err != ErrInvalidCharacter {
		t.Errorf("expected ErrInvalidCharacter, got %v", err)
	}
	_, err = m.Write("`") // '`'
	if err != ErrInvalidCharacter {
		t.Errorf("expected ErrInvalidCharacter, got %v", err)
	}
	_, err = m.Write("{") // '{'
	if err != ErrInvalidCharacter {
		t.Errorf("expected ErrInvalidCharacter, got %v", err)
	}
}

func TestWriteSpace(t *testing.T) {
	r := &mockRadio{}
	m := NewMorse(r, 1000, 10)
	m.Configure()
	_, err := m.Write(" ")
	if err != nil {
		t.Errorf("Write space failed: %v", err)
	}
	if r.standbyCnt == 0 {
		t.Error("Standby not called for space")
	}
}

func TestWriteValidChar(t *testing.T) {
	r := &mockRadio{}
	m := NewMorse(r, 4321, 20)
	m.Configure()
	_, err := m.Write("E")
	if err != nil {
		t.Errorf("Write failed: %v", err)
	}
	if len(r.transmitted) == 0 {
		t.Error("Transmit not called")
	}
	if r.standbyCnt == 0 {
		t.Error("Standby not called")
	}
}

func TestWriteMultipleChars(t *testing.T) {
	r := &mockRadio{}
	m := NewMorse(r, 4321, 20)
	m.Configure()
	_, err := m.Write("HI")
	if err != nil {
		t.Errorf("Write failed: %v", err)
	}
	if len(r.transmitted) == 0 {
		t.Error("Transmit not called")
	}
	if r.standbyCnt == 0 {
		t.Error("Standby not called")
	}
}

func TestWriteTransmitError(t *testing.T) {
	r := &mockRadio{failTx: true}
	m := NewMorse(r, 4321, 20)
	m.Configure()
	err := m.write('E')
	if err == nil {
		t.Error("expected transmit error")
	}
}

func TestWriteStandbyError(t *testing.T) {
	r := &mockRadio{failStby: true}
	m := NewMorse(r, 4321, 20)
	m.Configure()
	err := m.write('E')
	if err == nil {
		t.Error("expected standby error")
	}
}

// Optional: Test timing (not recommended for CI, but can be useful for local checks)
func TestWriteTiming(t *testing.T) {
	r := &mockRadio{}
	m := NewMorse(r, 4321, 60)
	m.Configure()
	start := time.Now()
	_ = m.write('E')
	elapsed := time.Since(start)
	if elapsed < time.Duration(m.dotLength)*time.Millisecond {
		t.Error("timing too short")
	}
}
