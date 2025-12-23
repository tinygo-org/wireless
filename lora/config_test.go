package lora

import (
	"testing"
)

func TestConfigStruct(t *testing.T) {
	cfg := Config{
		Freq:           MHz_868_1,
		Cr:             CodingRate4_5,
		Sf:             SpreadingFactor7,
		Bw:             Bandwidth_125_0,
		Ldr:            LowDataRateOptimizeOff,
		Preamble:       8,
		SyncWord:       0x3444,
		HeaderType:     HeaderExplicit,
		Crc:            CRCOn,
		Iq:             IQStandard,
		LoraTxPowerDBm: 14,
	}

	if cfg.Freq != 868100000 {
		t.Errorf("Freq = %d, want 868100000", cfg.Freq)
	}
	if cfg.Cr != CodingRate4_5 {
		t.Errorf("Cr = %d, want %d", cfg.Cr, CodingRate4_5)
	}
	if cfg.Sf != SpreadingFactor7 {
		t.Errorf("Sf = %d, want %d", cfg.Sf, SpreadingFactor7)
	}
	if cfg.LoraTxPowerDBm != 14 {
		t.Errorf("LoraTxPowerDBm = %d, want 14", cfg.LoraTxPowerDBm)
	}
}

func TestSpreadingFactorConstants(t *testing.T) {
	tests := []struct {
		name     string
		constant uint8
		expected uint8
	}{
		{"SpreadingFactor5", SpreadingFactor5, 0x05},
		{"SpreadingFactor6", SpreadingFactor6, 0x06},
		{"SpreadingFactor7", SpreadingFactor7, 0x07},
		{"SpreadingFactor8", SpreadingFactor8, 0x08},
		{"SpreadingFactor9", SpreadingFactor9, 0x09},
		{"SpreadingFactor10", SpreadingFactor10, 0x0A},
		{"SpreadingFactor11", SpreadingFactor11, 0x0B},
		{"SpreadingFactor12", SpreadingFactor12, 0x0C},
	}

	for _, tt := range tests {
		if tt.constant != tt.expected {
			t.Errorf("%s = 0x%02X, want 0x%02X", tt.name, tt.constant, tt.expected)
		}
	}
}

func TestCodingRateConstants(t *testing.T) {
	if CodingRate4_5 != 0x01 {
		t.Errorf("CodingRate4_5 = %d, want 1", CodingRate4_5)
	}
	if CodingRate4_6 != 0x02 {
		t.Errorf("CodingRate4_6 = %d, want 2", CodingRate4_6)
	}
	if CodingRate4_7 != 0x03 {
		t.Errorf("CodingRate4_7 = %d, want 3", CodingRate4_7)
	}
	if CodingRate4_8 != 0x04 {
		t.Errorf("CodingRate4_8 = %d, want 4", CodingRate4_8)
	}
}

func TestBandwidthConstants(t *testing.T) {
	if Bandwidth_7_8 != 0 {
		t.Errorf("Bandwidth_7_8 = %d, want 0", Bandwidth_7_8)
	}
	if Bandwidth_125_0 != 7 {
		t.Errorf("Bandwidth_125_0 = %d, want 7", Bandwidth_125_0)
	}
	if Bandwidth_500_0 != 9 {
		t.Errorf("Bandwidth_500_0 = %d, want 9", Bandwidth_500_0)
	}
}

func TestHeaderTypeConstants(t *testing.T) {
	if HeaderExplicit != 0x00 {
		t.Errorf("HeaderExplicit = %d, want 0", HeaderExplicit)
	}
	if HeaderImplicit != 0x01 {
		t.Errorf("HeaderImplicit = %d, want 1", HeaderImplicit)
	}
}

func TestCRCConstants(t *testing.T) {
	if CRCOff != 0x00 {
		t.Errorf("CRCOff = %d, want 0", CRCOff)
	}
	if CRCOn != 0x01 {
		t.Errorf("CRCOn = %d, want 1", CRCOn)
	}
}

func TestIQConstants(t *testing.T) {
	if IQStandard != 0x00 {
		t.Errorf("IQStandard = %d, want 0", IQStandard)
	}
	if IQInverted != 0x01 {
		t.Errorf("IQInverted = %d, want 1", IQInverted)
	}
}

func TestFrequencyConstants(t *testing.T) {
	if MHz_868_1 != 868100000 {
		t.Errorf("MHz_868_1 = %d, want 868100000", MHz_868_1)
	}
	if MHZ_915_0 != 915000000 {
		t.Errorf("MHZ_915_0 = %d, want 915000000", MHZ_915_0)
	}
}

func TestErrUndefinedLoraConf(t *testing.T) {
	if ErrUndefinedLoraConf == nil {
		t.Error("ErrUndefinedLoraConf is nil")
	}
	if ErrUndefinedLoraConf.Error() != "Undefined Lora configuration" {
		t.Errorf("ErrUndefinedLoraConf.Error() = %q, want %q",
			ErrUndefinedLoraConf.Error(), "Undefined Lora configuration")
	}
}

func TestSyncConstants(t *testing.T) {
	if SyncPublic != 0 {
		t.Errorf("SyncPublic = %d, want 0", SyncPublic)
	}
	if SyncPrivate != 1 {
		t.Errorf("SyncPrivate = %d, want 1", SyncPrivate)
	}
}
