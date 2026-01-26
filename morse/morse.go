package morse

import (
	"errors"
	"strings"
	"time"
)

var (
	// ErrInvalidCharacter is returned when an unsupported character is encountered.
	ErrInvalidCharacter = errors.New("invalid character for Morse code")
)

// Morse represents an Morse code modem.
type Morse struct {
	radio       Radio
	base        uint64
	speed       int
	dotLength   time.Duration
	dashLength  time.Duration
	letterSpace time.Duration
	wordSpace   time.Duration
}

// NewMorse creates a new Morse modem instance.
// radio: the Radio interface implementation
// base: the base frequency in Hz
// speed: the send speed in words per minute
func NewMorse(radio Radio, base uint64, speed int) *Morse {
	return &Morse{
		radio: radio,
		base:  base,
		speed: speed,
	}
}

func (m *Morse) Close() error {
	return m.radio.Close()
}

func (m *Morse) Configure() error {
	// calculate symbol lengths (using PARIS as typical word)
	m.dotLength = time.Duration(1200/m.speed) * time.Millisecond
	m.dashLength = time.Duration(3) * m.dotLength
	m.letterSpace = time.Duration(3) * m.dotLength
	m.wordSpace = time.Duration(7) * m.dotLength

	return nil
}

// GetBaseFrequency returns the current transmission frequency.
func (m *Morse) GetBaseFrequency() uint64 {
	return m.base
}

// GetSpeed returns the current send speed in words per minute.
func (m *Morse) GetSpeed() int {
	return m.speed
}

// Write sends data using Morse code modulation.
func (m *Morse) Write(data string) (int, error) {
	msg := strings.ToUpper(data)
	for _, b := range msg {
		err := m.write(byte(b))
		if err != nil {
			return 0, err
		}
	}

	return len(data), nil
}

func (m *Morse) write(b byte) error {
	// inter-word pause (space)
	if b == ' ' {
		if err := m.radio.Standby(); err != nil {
			return err
		}

		time.Sleep(m.wordSpace)
		return nil
	}

	// get morse code from lookup table
	code := ASCIIToMorse(b)
	if code == 0 {
		return ErrInvalidCharacter
	}

	// transmit each symbol
	for code != guardBit {
		if (code & Dash) == 1 {
			// dash
			if err := m.radio.Transmit(m.base * 100); err != nil {
				return err
			}
			time.Sleep(m.dashLength)
		} else {
			// dot
			if err := m.radio.Transmit(m.base * 100); err != nil {
				return err
			}
			time.Sleep(m.dotLength)
		}

		// inter-symbol pause
		if err := m.radio.Standby(); err != nil {
			return err
		}
		time.Sleep(m.dotLength)

		code >>= 1
	}

	// inter-letter pause
	if err := m.radio.Standby(); err != nil {
		return err
	}
	time.Sleep(m.letterSpace - m.dotLength)

	return nil
}
