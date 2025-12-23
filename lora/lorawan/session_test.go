package lorawan

import (
	"testing"
)

func TestSessionSetDevAddr(t *testing.T) {
	tests := []struct {
		name    string
		devAddr []uint8
		wantErr error
	}{
		{
			name:    "valid DevAddr",
			devAddr: []uint8{0x01, 0x02, 0x03, 0x04},
			wantErr: nil,
		},
		{
			name:    "too short",
			devAddr: []uint8{0x01, 0x02, 0x03},
			wantErr: ErrInvalidDevAddrLength,
		},
		{
			name:    "too long",
			devAddr: []uint8{0x01, 0x02, 0x03, 0x04, 0x05},
			wantErr: ErrInvalidDevAddrLength,
		},
		{
			name:    "empty",
			devAddr: []uint8{},
			wantErr: ErrInvalidDevAddrLength,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Session{}
			err := s.SetDevAddr(tt.devAddr)
			if err != tt.wantErr {
				t.Errorf("SetDevAddr() error = %v, want %v", err, tt.wantErr)
			}
			if err == nil {
				for i, b := range tt.devAddr {
					if s.DevAddr[i] != b {
						t.Errorf("DevAddr[%d] = %d, want %d", i, s.DevAddr[i], b)
					}
				}
			}
		})
	}
}

func TestSessionGetDevAddr(t *testing.T) {
	s := &Session{}
	s.DevAddr = [4]uint8{0xAB, 0xCD, 0xEF, 0x12}

	got := s.GetDevAddr()
	want := "abcdef12"
	if got != want {
		t.Errorf("GetDevAddr() = %q, want %q", got, want)
	}
}

func TestSessionSetNwkSKey(t *testing.T) {
	tests := []struct {
		name    string
		nwkSKey []uint8
		wantErr error
	}{
		{
			name:    "valid NwkSKey",
			nwkSKey: []uint8{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F, 0x10},
			wantErr: nil,
		},
		{
			name:    "too short",
			nwkSKey: []uint8{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08},
			wantErr: ErrInvalidNwkSKeyLength,
		},
		{
			name:    "too long",
			nwkSKey: []uint8{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F, 0x10, 0x11},
			wantErr: ErrInvalidNwkSKeyLength,
		},
		{
			name:    "empty",
			nwkSKey: []uint8{},
			wantErr: ErrInvalidNwkSKeyLength,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Session{}
			err := s.SetNwkSKey(tt.nwkSKey)
			if err != tt.wantErr {
				t.Errorf("SetNwkSKey() error = %v, want %v", err, tt.wantErr)
			}
			if err == nil {
				for i, b := range tt.nwkSKey {
					if s.NwkSKey[i] != b {
						t.Errorf("NwkSKey[%d] = %d, want %d", i, s.NwkSKey[i], b)
					}
				}
			}
		})
	}
}

func TestSessionGetNwkSKey(t *testing.T) {
	s := &Session{}
	s.NwkSKey = [16]uint8{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F, 0x10}

	got := s.GetNwkSKey()
	want := "0102030405060708090a0b0c0d0e0f10"
	if got != want {
		t.Errorf("GetNwkSKey() = %q, want %q", got, want)
	}
}

func TestSessionSetAppSKey(t *testing.T) {
	tests := []struct {
		name    string
		appSKey []uint8
		wantErr error
	}{
		{
			name:    "valid AppSKey",
			appSKey: []uint8{0xA1, 0xA2, 0xA3, 0xA4, 0xA5, 0xA6, 0xA7, 0xA8, 0xA9, 0xAA, 0xAB, 0xAC, 0xAD, 0xAE, 0xAF, 0xB0},
			wantErr: nil,
		},
		{
			name:    "too short",
			appSKey: []uint8{0xA1, 0xA2, 0xA3, 0xA4},
			wantErr: ErrInvalidAppSKeyLength,
		},
		{
			name:    "too long",
			appSKey: []uint8{0xA1, 0xA2, 0xA3, 0xA4, 0xA5, 0xA6, 0xA7, 0xA8, 0xA9, 0xAA, 0xAB, 0xAC, 0xAD, 0xAE, 0xAF, 0xB0, 0xB1},
			wantErr: ErrInvalidAppSKeyLength,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Session{}
			err := s.SetAppSKey(tt.appSKey)
			if err != tt.wantErr {
				t.Errorf("SetAppSKey() error = %v, want %v", err, tt.wantErr)
			}
			if err == nil {
				for i, b := range tt.appSKey {
					if s.AppSKey[i] != b {
						t.Errorf("AppSKey[%d] = %d, want %d", i, s.AppSKey[i], b)
					}
				}
			}
		})
	}
}

func TestSessionGetAppSKey(t *testing.T) {
	s := &Session{}
	s.AppSKey = [16]uint8{0xA1, 0xA2, 0xA3, 0xA4, 0xA5, 0xA6, 0xA7, 0xA8, 0xA9, 0xAA, 0xAB, 0xAC, 0xAD, 0xAE, 0xAF, 0xB0}

	got := s.GetAppSKey()
	want := "a1a2a3a4a5a6a7a8a9aaabacadaeafb0"
	if got != want {
		t.Errorf("GetAppSKey() = %q, want %q", got, want)
	}
}

func TestSessionStruct(t *testing.T) {
	s := Session{
		NwkSKey:    [16]uint8{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F, 0x10},
		AppSKey:    [16]uint8{0xA1, 0xA2, 0xA3, 0xA4, 0xA5, 0xA6, 0xA7, 0xA8, 0xA9, 0xAA, 0xAB, 0xAC, 0xAD, 0xAE, 0xAF, 0xB0},
		DevAddr:    [4]uint8{0xDE, 0xAD, 0xBE, 0xEF},
		FCntDown:   100,
		FCntUp:     200,
		CFList:     [16]uint8{0x00},
		RXDelay:    1,
		DLSettings: 0x03,
	}

	if s.FCntDown != 100 {
		t.Errorf("FCntDown = %d, want 100", s.FCntDown)
	}
	if s.FCntUp != 200 {
		t.Errorf("FCntUp = %d, want 200", s.FCntUp)
	}
	if s.RXDelay != 1 {
		t.Errorf("RXDelay = %d, want 1", s.RXDelay)
	}
	if s.DLSettings != 0x03 {
		t.Errorf("DLSettings = %d, want 3", s.DLSettings)
	}
}

func TestGenMessageUplink(t *testing.T) {
	s := &Session{
		NwkSKey: [16]uint8{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F, 0x10},
		AppSKey: [16]uint8{0xA1, 0xA2, 0xA3, 0xA4, 0xA5, 0xA6, 0xA7, 0xA8, 0xA9, 0xAA, 0xAB, 0xAC, 0xAD, 0xAE, 0xAF, 0xB0},
		DevAddr: [4]uint8{0xDE, 0xAD, 0xBE, 0xEF},
		FCntUp:  0,
	}

	payload := []uint8{0x48, 0x65, 0x6C, 0x6C, 0x6F} // "Hello"
	msg, err := s.GenMessage(0, payload)

	if err != nil {
		t.Fatalf("GenMessage() error = %v", err)
	}

	if msg == nil {
		t.Fatal("GenMessage() returned nil")
	}

	// Check MHDR (first byte should be 0x40 for unconfirmed uplink)
	if msg[0] != 0x40 {
		t.Errorf("MHDR = 0x%02X, want 0x40", msg[0])
	}

	// Check DevAddr (bytes 1-4)
	for i := 0; i < 4; i++ {
		if msg[1+i] != s.DevAddr[i] {
			t.Errorf("DevAddr[%d] = 0x%02X, want 0x%02X", i, msg[1+i], s.DevAddr[i])
		}
	}

	// Check FCtrl (byte 5 should be 0x00)
	if msg[5] != 0x00 {
		t.Errorf("FCtrl = 0x%02X, want 0x00", msg[5])
	}

	// Message should include MHDR(1) + DevAddr(4) + FCtrl(1) + FCnt(2) + FPort(1) + encrypted payload + MIC(4)
	expectedMinLen := 1 + 4 + 1 + 2 + 1 + len(payload) + 4
	if len(msg) < expectedMinLen {
		t.Errorf("message length = %d, want >= %d", len(msg), expectedMinLen)
	}
}

func TestGenMessageIncrementsFCntUp(t *testing.T) {
	s := &Session{
		NwkSKey: [16]uint8{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F, 0x10},
		AppSKey: [16]uint8{0xA1, 0xA2, 0xA3, 0xA4, 0xA5, 0xA6, 0xA7, 0xA8, 0xA9, 0xAA, 0xAB, 0xAC, 0xAD, 0xAE, 0xAF, 0xB0},
		DevAddr: [4]uint8{0xDE, 0xAD, 0xBE, 0xEF},
		FCntUp:  0,
	}

	payload := []uint8{0x01, 0x02, 0x03}

	// First message
	_, err := s.GenMessage(0, payload)
	if err != nil {
		t.Fatalf("GenMessage() error = %v", err)
	}
	if s.FCntUp != 1 {
		t.Errorf("FCntUp after first message = %d, want 1", s.FCntUp)
	}

	// Second message
	_, err = s.GenMessage(0, payload)
	if err != nil {
		t.Fatalf("GenMessage() error = %v", err)
	}
	if s.FCntUp != 2 {
		t.Errorf("FCntUp after second message = %d, want 2", s.FCntUp)
	}
}

func TestGenMessageDownlinkDoesNotIncrementFCntUp(t *testing.T) {
	s := &Session{
		NwkSKey:  [16]uint8{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F, 0x10},
		AppSKey:  [16]uint8{0xA1, 0xA2, 0xA3, 0xA4, 0xA5, 0xA6, 0xA7, 0xA8, 0xA9, 0xAA, 0xAB, 0xAC, 0xAD, 0xAE, 0xAF, 0xB0},
		DevAddr:  [4]uint8{0xDE, 0xAD, 0xBE, 0xEF},
		FCntUp:   5,
		FCntDown: 10,
	}

	payload := []uint8{0x01, 0x02, 0x03}

	_, err := s.GenMessage(1, payload) // dir=1 for downlink
	if err != nil {
		t.Fatalf("GenMessage() error = %v", err)
	}

	if s.FCntUp != 5 {
		t.Errorf("FCntUp should not change for downlink, got %d, want 5", s.FCntUp)
	}
}

func TestGenMessageEmptyPayload(t *testing.T) {
	s := &Session{
		NwkSKey: [16]uint8{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F, 0x10},
		AppSKey: [16]uint8{0xA1, 0xA2, 0xA3, 0xA4, 0xA5, 0xA6, 0xA7, 0xA8, 0xA9, 0xAA, 0xAB, 0xAC, 0xAD, 0xAE, 0xAF, 0xB0},
		DevAddr: [4]uint8{0xDE, 0xAD, 0xBE, 0xEF},
		FCntUp:  0,
	}

	payload := []uint8{}
	msg, err := s.GenMessage(0, payload)

	if err != nil {
		t.Fatalf("GenMessage() error = %v", err)
	}

	if msg == nil {
		t.Fatal("GenMessage() returned nil")
	}

	// MHDR(1) + DevAddr(4) + FCtrl(1) + FCnt(2) + FPort(1) + MIC(4) = 13 bytes minimum
	if len(msg) < 13 {
		t.Errorf("message length = %d, want >= 13", len(msg))
	}
}

func TestGenMessageLargePayload(t *testing.T) {
	s := &Session{
		NwkSKey: [16]uint8{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F, 0x10},
		AppSKey: [16]uint8{0xA1, 0xA2, 0xA3, 0xA4, 0xA5, 0xA6, 0xA7, 0xA8, 0xA9, 0xAA, 0xAB, 0xAC, 0xAD, 0xAE, 0xAF, 0xB0},
		DevAddr: [4]uint8{0xDE, 0xAD, 0xBE, 0xEF},
		FCntUp:  0,
	}

	// Create a payload that spans multiple AES blocks (each block is 16 bytes)
	payload := make([]uint8, 50)
	for i := range payload {
		payload[i] = uint8(i)
	}

	msg, err := s.GenMessage(0, payload)

	if err != nil {
		t.Fatalf("GenMessage() error = %v", err)
	}

	if msg == nil {
		t.Fatal("GenMessage() returned nil")
	}

	// Verify message is properly formed
	expectedMinLen := 1 + 4 + 1 + 2 + 1 + len(payload) + 4
	if len(msg) < expectedMinLen {
		t.Errorf("message length = %d, want >= %d", len(msg), expectedMinLen)
	}
}

func TestSessionZeroValues(t *testing.T) {
	s := Session{}

	if s.FCntUp != 0 {
		t.Errorf("FCntUp zero value = %d, want 0", s.FCntUp)
	}
	if s.FCntDown != 0 {
		t.Errorf("FCntDown zero value = %d, want 0", s.FCntDown)
	}
	if s.RXDelay != 0 {
		t.Errorf("RXDelay zero value = %d, want 0", s.RXDelay)
	}
	if s.DLSettings != 0 {
		t.Errorf("DLSettings zero value = %d, want 0", s.DLSettings)
	}
}

func TestGetDevAddrZeroValue(t *testing.T) {
	s := &Session{}
	got := s.GetDevAddr()
	want := "00000000"
	if got != want {
		t.Errorf("GetDevAddr() zero value = %q, want %q", got, want)
	}
}

func TestGetNwkSKeyZeroValue(t *testing.T) {
	s := &Session{}
	got := s.GetNwkSKey()
	want := "00000000000000000000000000000000"
	if got != want {
		t.Errorf("GetNwkSKey() zero value = %q, want %q", got, want)
	}
}

func TestGetAppSKeyZeroValue(t *testing.T) {
	s := &Session{}
	got := s.GetAppSKey()
	want := "00000000000000000000000000000000"
	if got != want {
		t.Errorf("GetAppSKey() zero value = %q, want %q", got, want)
	}
}
