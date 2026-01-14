package u4b

// The U4B telemetry message uses the U4B format,
// a special encoding of the callsign and location
// to carry telemetry data over WSPR.
// For more information, see:
// https://qrp-labs.com/u4b/u4bdecoding.html
// https://traquito.github.io/pro/telemetry/

import (
	"errors"

	"github.com/chewxy/math32"
	"tinygo.org/x/wireless/wspr"
)

var (
	ErrInvalidTelemetry = errors.New("invalid telemetry data")
)

// NewMessage encodes telemetry data using the U4B format
// into a WSPR message.
// The parameters are:
// channel: 2-character channel identifier
// grid: 2-character Maidenhead grid square
// altitude: altitude in meters
// temperature: temperature in degrees Celsius
// voltage: voltage in millivolts
// speed: speed in km/h
func NewMessage(channel string, grid string, altitude int, temperature int, voltage int, speed int) (wspr.Message, error) {
	callsign, err := encodeTelemetryCallSign(channel, grid, altitude)
	if err != nil {
		return 0, err
	}

	c, err := wspr.CallSign(callsign)
	if err != nil {
		return 0, err
	}

	g, p, err := encodeTelemetryGridPower(temperature, voltage, speed)
	if err != nil {
		return 0, err
	}

	l, err := wspr.Locator(g)
	if err != nil {
		return 0, err
	}

	return wspr.Message((c << 28) + (l << 13) + (wspr.Power(p) << 6)), nil
}

func encodeTelemetryCallSign(channel, grid string, altitude int) (string, error) {
	if len(grid) != 2 || len(channel) != 2 {
		return "", ErrInvalidTelemetry
	}

	grid5Val := grid[0] - 'A'
	grid6Val := grid[1] - 'A'
	alt := altitude / 20

	val := uint32(grid5Val)
	val = val*24 + uint32(grid6Val)
	val = val*1068 + uint32(alt)

	// extract
	id6Val := val % 26
	val = val / 26
	id5Val := val % 26
	val = val / 26
	id4Val := val % 26
	val = val / 26
	id2Val := val % 36
	val = val / 36

	// convert to encoded form
	id2 := encodeBase36(uint8(id2Val))
	id4 := 'A' + byte(id4Val)
	id5 := 'A' + byte(id5Val)
	id6 := 'A' + byte(id6Val)

	callsign := string([]byte{channel[0], id2, channel[1], id4, id5, id6})

	return callsign, nil
}

func encodeBase36(val uint8) byte {
	if val < 10 {
		return '0' + val
	}

	return 'A' + (val - 10)
}

var powerDbmList = [19]uint8{
	0, 3, 7,
	10, 13, 17,
	20, 23, 27,
	30, 33, 37,
	40, 43, 47,
	50, 53, 57,
	60,
}

func encodeTelemetryGridPower(temperature int, voltage int, speed int) (string, int, error) {
	// convert inputs to encoded numbers
	tempCNum := uint8(temperature + 50)
	voltageNum := (uint8(math32.Round((float32(voltage)-300)/5)) + 20) % 40
	speedKnotsNum := uint8(math32.Round(float32(speed) / 2.0))
	gpsValidNum := uint8(1)

	// convert inputs into a "big number"
	val := uint32(0)
	val = val*90 + uint32(tempCNum)
	val = val*40 + uint32(voltageNum)
	val = val*42 + uint32(speedKnotsNum)
	val = val*2 + uint32(gpsValidNum)
	val = val*2 + 1 // basic telemetry

	// extract data from big number
	powerVal := uint8(val % 19)
	val = val / 19
	g4Val := uint8(val % 10)
	val = val / 10
	g3Val := uint8(val % 10)
	val = val / 10
	g2Val := uint8(val % 18)
	val = val / 18
	g1Val := uint8(val % 18)
	val = val / 18

	// convert to encoded form
	g1 := 'A' + g1Val
	g2 := 'A' + g2Val
	g3 := '0' + g3Val
	g4 := '0' + g4Val

	// store grid
	grid := string([]byte{g1, g2, g3, g4})

	return grid, int(powerDbmList[powerVal]), nil
}
