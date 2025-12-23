package lorawan

import (
	"errors"
	"testing"

	"tinygo.org/x/wireless/lora"
	"tinygo.org/x/wireless/lora/lorawan/region"
)

// mockRadio implements lora.Radio interface for testing
type mockRadio struct {
	frequency       uint32
	bandwidth       uint8
	codingRate      uint8
	spreadingFactor uint8
	preambleLength  uint16
	txPower         int8
	headerType      uint8
	crcEnabled      bool
	iqMode          uint8
	publicNetwork   bool
	syncWord        uint16

	txCalled    bool
	txPayload   []uint8
	txTimeout   uint32
	txError     error
	rxCalled    bool
	rxTimeout   uint32
	rxResponse  []uint8
	rxError     error
	resetCalled bool
}

func (m *mockRadio) Reset()                        { m.resetCalled = true }
func (m *mockRadio) SetFrequency(freq uint32)      { m.frequency = freq }
func (m *mockRadio) SetBandwidth(bw uint8)         { m.bandwidth = bw }
func (m *mockRadio) SetCodingRate(cr uint8)        { m.codingRate = cr }
func (m *mockRadio) SetSpreadingFactor(sf uint8)   { m.spreadingFactor = sf }
func (m *mockRadio) SetPreambleLength(pl uint16)   { m.preambleLength = pl }
func (m *mockRadio) SetTxPower(pwr int8)           { m.txPower = pwr }
func (m *mockRadio) SetHeaderType(ht uint8)        { m.headerType = ht }
func (m *mockRadio) SetCrc(enabled bool)           { m.crcEnabled = enabled }
func (m *mockRadio) SetIqMode(mode uint8)          { m.iqMode = mode }
func (m *mockRadio) SetPublicNetwork(enabled bool) { m.publicNetwork = enabled }
func (m *mockRadio) SetSyncWord(sw uint16)         { m.syncWord = sw }
func (m *mockRadio) LoraConfig(cnf lora.Config)    {}

func (m *mockRadio) Tx(pkt []uint8, timeout uint32) error {
	m.txCalled = true
	m.txPayload = pkt
	m.txTimeout = timeout
	return m.txError
}

func (m *mockRadio) Rx(timeout uint32) ([]uint8, error) {
	m.rxCalled = true
	m.rxTimeout = timeout
	return m.rxResponse, m.rxError
}

// mockChannel implements region.Channel interface for testing
type mockChannel struct {
	frequency       uint32
	bandwidth       uint8
	spreadingFactor uint8
	codingRate      uint8
	preambleLength  uint16
	txPowerDBm      int8
	nextResult      bool
	nextCalled      bool
}

func (m *mockChannel) Next() bool {
	m.nextCalled = true
	return m.nextResult
}
func (m *mockChannel) Frequency() uint32          { return m.frequency }
func (m *mockChannel) Bandwidth() uint8           { return m.bandwidth }
func (m *mockChannel) SpreadingFactor() uint8     { return m.spreadingFactor }
func (m *mockChannel) CodingRate() uint8          { return m.codingRate }
func (m *mockChannel) PreambleLength() uint16     { return m.preambleLength }
func (m *mockChannel) TxPowerDBm() int8           { return m.txPowerDBm }
func (m *mockChannel) SetFrequency(v uint32)      { m.frequency = v }
func (m *mockChannel) SetBandwidth(v uint8)       { m.bandwidth = v }
func (m *mockChannel) SetSpreadingFactor(v uint8) { m.spreadingFactor = v }
func (m *mockChannel) SetCodingRate(v uint8)      { m.codingRate = v }
func (m *mockChannel) SetPreambleLength(v uint16) { m.preambleLength = v }
func (m *mockChannel) SetTxPowerDBm(v int8)       { m.txPowerDBm = v }

// mockSettings implements region.Settings interface for testing
type mockSettings struct {
	joinRequestCh region.Channel
	joinAcceptCh  region.Channel
	uplinkCh      region.Channel
}

func (m *mockSettings) JoinRequestChannel() region.Channel { return m.joinRequestCh }
func (m *mockSettings) JoinAcceptChannel() region.Channel  { return m.joinAcceptCh }
func (m *mockSettings) UplinkChannel() region.Channel      { return m.uplinkCh }

// Helper to reset global state before each test
func resetGlobalState() {
	ActiveRadio = nil
	regionSettings = nil
	Retries = 15
}

func TestErrorDefinitions(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		expected string
	}{
		{"ErrNoJoinAcceptReceived", ErrNoJoinAcceptReceived, "no JoinAccept packet received"},
		{"ErrNoRadioAttached", ErrNoRadioAttached, "no LoRa radio attached"},
		{"ErrInvalidEuiLength", ErrInvalidEuiLength, "invalid EUI length"},
		{"ErrInvalidAppKeyLength", ErrInvalidAppKeyLength, "invalid AppKey length"},
		{"ErrInvalidPacketLength", ErrInvalidPacketLength, "invalid packet length"},
		{"ErrInvalidDevAddrLength", ErrInvalidDevAddrLength, "invalid DevAddr length"},
		{"ErrInvalidMic", ErrInvalidMic, "invalid Mic"},
		{"ErrFrmPayloadTooLarge", ErrFrmPayloadTooLarge, "FRM payload too large"},
		{"ErrInvalidNetIDLength", ErrInvalidNetIDLength, "invalid NetID length"},
		{"ErrInvalidNwkSKeyLength", ErrInvalidNwkSKeyLength, "invalid NwkSKey length"},
		{"ErrInvalidAppSKeyLength", ErrInvalidAppSKeyLength, "invalid AppSKey length"},
		{"ErrUndefinedRegionSettings", ErrUndefinedRegionSettings, "undefined Regionnal Settings "},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.err == nil {
				t.Errorf("%s is nil", tt.name)
			}
			if tt.err.Error() != tt.expected {
				t.Errorf("%s.Error() = %q, want %q", tt.name, tt.err.Error(), tt.expected)
			}
		})
	}
}

func TestErrorsAreDistinct(t *testing.T) {
	allErrors := []error{
		ErrNoJoinAcceptReceived,
		ErrNoRadioAttached,
		ErrInvalidEuiLength,
		ErrInvalidAppKeyLength,
		ErrInvalidPacketLength,
		ErrInvalidDevAddrLength,
		ErrInvalidMic,
		ErrFrmPayloadTooLarge,
		ErrInvalidNetIDLength,
		ErrInvalidNwkSKeyLength,
		ErrInvalidAppSKeyLength,
		ErrUndefinedRegionSettings,
	}

	for i, err1 := range allErrors {
		for j, err2 := range allErrors {
			if i != j && errors.Is(err1, err2) {
				t.Errorf("errors at index %d and %d should be distinct", i, j)
			}
		}
	}
}

func TestConstants(t *testing.T) {
	if LORA_TX_TIMEOUT != 2000 {
		t.Errorf("LORA_TX_TIMEOUT = %d, want 2000", LORA_TX_TIMEOUT)
	}
	if LORA_RX_TIMEOUT != 10000 {
		t.Errorf("LORA_RX_TIMEOUT = %d, want 10000", LORA_RX_TIMEOUT)
	}
}

func TestDefaultRetries(t *testing.T) {
	if Retries != 15 {
		t.Errorf("Retries = %d, want 15", Retries)
	}
}

func TestUseRadio(t *testing.T) {
	resetGlobalState()
	defer resetGlobalState()

	radio := &mockRadio{}
	UseRadio(radio)

	if ActiveRadio != radio {
		t.Error("UseRadio did not set ActiveRadio correctly")
	}
}

func TestUseRadioPanicsWhenAlreadySet(t *testing.T) {
	resetGlobalState()
	defer resetGlobalState()

	defer func() {
		if r := recover(); r == nil {
			t.Error("UseRadio did not panic when radio already set")
		}
	}()

	ActiveRadio = &mockRadio{}
	UseRadio(&mockRadio{})
}

func TestUseRegionSettings(t *testing.T) {
	resetGlobalState()
	defer resetGlobalState()

	settings := &mockSettings{}
	UseRegionSettings(settings)

	if regionSettings != settings {
		t.Error("UseRegionSettings did not set regionSettings correctly")
	}
}

func TestSetPublicNetworkTrue(t *testing.T) {
	resetGlobalState()
	defer resetGlobalState()

	radio := &mockRadio{}
	ActiveRadio = radio

	SetPublicNetwork(true)
	if !radio.publicNetwork {
		t.Error("SetPublicNetwork(true) did not set publicNetwork to true")
	}
}

func TestSetPublicNetworkFalse(t *testing.T) {
	resetGlobalState()
	defer resetGlobalState()

	radio := &mockRadio{publicNetwork: true}
	ActiveRadio = radio

	SetPublicNetwork(false)
	if radio.publicNetwork {
		t.Error("SetPublicNetwork(false) did not set publicNetwork to false")
	}
}

func TestJoinWithNoRadio(t *testing.T) {
	resetGlobalState()
	defer resetGlobalState()

	otaa := &Otaa{}
	session := &Session{}

	err := Join(otaa, session)
	if err != ErrNoRadioAttached {
		t.Errorf("Join() error = %v, want %v", err, ErrNoRadioAttached)
	}
}

func TestJoinWithNoRegionSettings(t *testing.T) {
	resetGlobalState()
	defer resetGlobalState()

	radio := &mockRadio{}
	ActiveRadio = radio

	otaa := &Otaa{}
	session := &Session{}

	err := Join(otaa, session)
	if err != ErrUndefinedRegionSettings {
		t.Errorf("Join() error = %v, want %v", err, ErrUndefinedRegionSettings)
	}
}

func TestSendUplinkWithNoRegionSettings(t *testing.T) {
	resetGlobalState()
	defer resetGlobalState()

	session := &Session{}
	err := SendUplink([]byte("test"), session)
	if err != ErrUndefinedRegionSettings {
		t.Errorf("SendUplink() error = %v, want %v", err, ErrUndefinedRegionSettings)
	}
}

func TestListenDownlink(t *testing.T) {
	err := ListenDownlink()
	if err != nil {
		t.Errorf("ListenDownlink() error = %v, want nil", err)
	}
}

func TestMockRadioImplementsInterface(t *testing.T) {
	var _ lora.Radio = (*mockRadio)(nil)
}

func TestMockChannelImplementsInterface(t *testing.T) {
	var _ region.Channel = (*mockChannel)(nil)
}

func TestMockSettingsImplementsInterface(t *testing.T) {
	var _ region.Settings = (*mockSettings)(nil)
}

func TestRetriesCanBeModified(t *testing.T) {
	originalRetries := Retries
	defer func() { Retries = originalRetries }()

	Retries = 5
	if Retries != 5 {
		t.Errorf("Retries = %d, want 5", Retries)
	}

	Retries = 20
	if Retries != 20 {
		t.Errorf("Retries = %d, want 20", Retries)
	}
}
