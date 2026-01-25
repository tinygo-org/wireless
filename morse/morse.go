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
	dotLength   int
	dashLength  int
	letterSpace int
	wordSpace   int
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
	m.dotLength = 1200 / m.speed
	m.dashLength = 3 * m.dotLength
	m.letterSpace = 3 * m.dotLength
	m.wordSpace = 4 * m.dotLength

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
	if (b < ' ') || (b == 0x60) || (b > 'z') {
		return ErrInvalidCharacter
	}

	// inter-word pause (space)
	if b == ' ' {
		m.radio.Standby()
		time.Sleep(time.Duration(m.wordSpace) * time.Millisecond)
		return nil
	}

	// get morse code from lookup table
	code := ASCIIToMorse(b)
	if code == 0 {
		return ErrInvalidCharacter
	}

	if code == Unsupported {
		return ErrInvalidCharacter
	}

	// transmit each symbol
	for code != guardBit {
		if (code & Dash) == 1 {
			// dash
			if err := m.radio.Transmit(m.base * 100); err != nil {
				return err
			}
			time.Sleep(time.Duration(m.dashLength) * time.Millisecond)
		} else {
			// dot
			if err := m.radio.Transmit(m.base * 100); err != nil {
				return err
			}
			time.Sleep(time.Duration(m.dotLength) * time.Millisecond)
		}

		// inter-symbol pause
		if err := m.radio.Standby(); err != nil {
			return err
		}
		time.Sleep(time.Duration(m.dotLength) * time.Millisecond)

		code >>= 1
	}

	// inter-letter pause
	if err := m.radio.Standby(); err != nil {
		return err
	}
	time.Sleep(time.Duration(m.letterSpace-m.dotLength) * time.Millisecond)

	return nil
}
