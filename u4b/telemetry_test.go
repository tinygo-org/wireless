package u4b

import (
	"testing"
)

func TestEncodeBase36(t *testing.T) {
	tests := []struct {
		val      uint8
		expected byte
	}{
		{0, '0'},
		{9, '9'},
		{10, 'A'},
		{15, 'F'},
		{35, 'Z'},
	}
	for _, tt := range tests {
		got := encodeBase36(tt.val)
		if got != tt.expected {
			t.Errorf("encodeBase36(%d) = %c, want %c", tt.val, got, tt.expected)
		}
	}
}

func TestEncodeTelemetryCallSign(t *testing.T) {
	callsign, err := encodeTelemetryCallSign("AB", "CD", 100)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if len(callsign) != 6 {
		t.Errorf("callsign length = %d, want 6", len(callsign))
	}

	_, err = encodeTelemetryCallSign("A", "CD", 100)
	if err == nil {
		t.Error("expected error for short channel, got nil")
	}
	_, err = encodeTelemetryCallSign("AB", "C", 100)
	if err == nil {
		t.Error("expected error for short grid, got nil")
	}
}

func TestEncodeTelemetryGridPower(t *testing.T) {
	grid, power, err := encodeTelemetryGridPower(27, 500, 10)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if len(grid) != 4 {
		t.Errorf("grid length = %d, want 4", len(grid))
	}
	if power < 0 || power > 60 {
		t.Errorf("power = %d, out of expected range", power)
	}
}

func TestEncodeTelemetryGridPower2(t *testing.T) {
	grid, power, err := encodeTelemetryGridPower(-27, 3700, 10)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if len(grid) != 4 {
		t.Errorf("grid length = %d, want 4", len(grid))
	}
	if power < 0 || power > 60 {
		t.Errorf("power = %d, out of expected range", power)
	}
}

func TestNewTelemetryMessage(t *testing.T) {
	msg, err := NewMessage("AB", "CD", 100, 27, 3700, 10)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if msg == 0 {
		t.Error("expected non-zero message")
	}

	_, err = NewMessage("A", "CD", 100, 27, 500, 10)
	if err == nil {
		t.Error("expected error for short channel, got nil")
	}
	_, err = NewMessage("AB", "C", 100, 27, 500, 10)
	if err == nil {
		t.Error("expected error for short grid, got nil")
	}
}
