package lora

import (
	"testing"
)

func TestRadioEventConstants(t *testing.T) {
	tests := []struct {
		name     string
		constant int
		expected int
	}{
		{"RadioEventRxDone", RadioEventRxDone, 0},
		{"RadioEventTxDone", RadioEventTxDone, 1},
		{"RadioEventTimeout", RadioEventTimeout, 2},
		{"RadioEventWatchdog", RadioEventWatchdog, 3},
		{"RadioEventCrcError", RadioEventCrcError, 4},
		{"RadioEventUnhandled", RadioEventUnhandled, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.constant != tt.expected {
				t.Errorf("%s = %d, want %d", tt.name, tt.constant, tt.expected)
			}
		})
	}
}

func TestNewRadioEvent(t *testing.T) {
	tests := []struct {
		name      string
		eventType int
		irqStatus uint16
		eventData []byte
	}{
		{
			name:      "RxDone with data",
			eventType: RadioEventRxDone,
			irqStatus: 0x01,
			eventData: []byte{0x48, 0x65, 0x6C, 0x6C, 0x6F},
		},
		{
			name:      "TxDone no data",
			eventType: RadioEventTxDone,
			irqStatus: 0x02,
			eventData: nil,
		},
		{
			name:      "Timeout empty data",
			eventType: RadioEventTimeout,
			irqStatus: 0x00,
			eventData: []byte{},
		},
		{
			name:      "CrcError with status",
			eventType: RadioEventCrcError,
			irqStatus: 0xFFFF,
			eventData: []byte{0x01},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			event := NewRadioEvent(tt.eventType, tt.irqStatus, tt.eventData)

			if event.EventType != tt.eventType {
				t.Errorf("EventType = %d, want %d", event.EventType, tt.eventType)
			}
			if event.IRQStatus != tt.irqStatus {
				t.Errorf("IRQStatus = %d, want %d", event.IRQStatus, tt.irqStatus)
			}
			if len(event.EventData) != len(tt.eventData) {
				t.Errorf("EventData length = %d, want %d", len(event.EventData), len(tt.eventData))
			}
			for i, b := range event.EventData {
				if b != tt.eventData[i] {
					t.Errorf("EventData[%d] = %d, want %d", i, b, tt.eventData[i])
				}
			}
		})
	}
}

func TestRadioEventStruct(t *testing.T) {
	data := []byte{0xDE, 0xAD, 0xBE, 0xEF}
	event := RadioEvent{
		EventType: RadioEventRxDone,
		IRQStatus: 0x1234,
		EventData: data,
	}

	if event.EventType != RadioEventRxDone {
		t.Errorf("EventType = %d, want %d", event.EventType, RadioEventRxDone)
	}
	if event.IRQStatus != 0x1234 {
		t.Errorf("IRQStatus = 0x%04X, want 0x1234", event.IRQStatus)
	}
	if len(event.EventData) != 4 {
		t.Errorf("EventData length = %d, want 4", len(event.EventData))
	}
}

func TestNewRadioEventReturnsValue(t *testing.T) {
	// Verify that NewRadioEvent returns a value, not a pointer
	event := NewRadioEvent(RadioEventTxDone, 0x00, nil)

	// Modify the returned event to ensure it's a copy
	event.EventType = RadioEventCrcError

	// Create another event to verify independence
	event2 := NewRadioEvent(RadioEventTxDone, 0x00, nil)

	if event2.EventType != RadioEventTxDone {
		t.Errorf("event2.EventType = %d, want %d", event2.EventType, RadioEventTxDone)
	}
}
