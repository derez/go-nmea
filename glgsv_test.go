package nmea

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGLGSVGoodSentence(t *testing.T) {
	goodMsg := "$GLGSV,3,1,11,03,03,111,00,04,15,270,00,06,01,010,12,13,06,292,00*6B"
	s, err := Parse(goodMsg)

	assert.NoError(t, err, "Unexpected error parsing good sentence")
	assert.Equal(t, PrefixGLGSV, s.GetSentence().Type, "Prefix does not match")

	sentence := s.(GLGSV)
	assert.Equal(t, int64(3), sentence.TotalMessages, "Total messages does not match")
	assert.Equal(t, int64(1), sentence.MessageNumber, "Message number does not match")
	assert.Equal(t, int64(11), sentence.NumberSVsInView, "Number of SVs in view does not match")

	assert.Equal(t, int64(3), sentence.Info[0].SVPRNNumber, "Number of Info[0] SV PRN does not match")
	assert.Equal(t, int64(3), sentence.Info[0].Elevation, "Number of Info[0] Elevation does not match")
	assert.Equal(t, int64(111), sentence.Info[0].Azimuth, "Number of Info[0] Azimuth does not match")
	assert.Equal(t, int64(0), sentence.Info[0].SNR, "Number of Info[0] SNR does not match")

	assert.Equal(t, int64(4), sentence.Info[1].SVPRNNumber, "Number of Info[1] SV PRN does not match")
	assert.Equal(t, int64(15), sentence.Info[1].Elevation, "Number of Info[1] Elevation does not match")
	assert.Equal(t, int64(270), sentence.Info[1].Azimuth, "Number of Info[1] Azimuth does not match")
	assert.Equal(t, int64(0), sentence.Info[1].SNR, "Number of Info[1] SNR does not match")

	assert.Equal(t, int64(6), sentence.Info[2].SVPRNNumber, "Number of Info[2] SV PRN does not match")
	assert.Equal(t, int64(1), sentence.Info[2].Elevation, "Number of Info[2] Elevation does not match")
	assert.Equal(t, int64(10), sentence.Info[2].Azimuth, "Number of Info[2] Azimuth does not match")
	assert.Equal(t, int64(12), sentence.Info[2].SNR, "Number of Info[2] SNR does not match")

	assert.Equal(t, int64(13), sentence.Info[3].SVPRNNumber, "Number of Info[3] SV PRN does not match")
	assert.Equal(t, int64(6), sentence.Info[3].Elevation, "Number of Info[3] Elevation does not match")
	assert.Equal(t, int64(292), sentence.Info[3].Azimuth, "Number of Info[3] Azimuth does not match")
	assert.Equal(t, int64(0), sentence.Info[3].SNR, "Number of Info[3] SNR does not match")
}

func TestGLGSVShort(t *testing.T) {
	goodMsg := "$GLGSV,3,1,11,03,03,111,00,04,15,270,00,06,01,010,12*56"
	s, err := Parse(goodMsg)

	assert.NoError(t, err, "Unexpected error parsing good sentence")
	assert.Equal(t, PrefixGLGSV, s.GetSentence().Type, "Prefix does not match")

	sentence := s.(GLGSV)
	assert.Equal(t, int64(3), sentence.TotalMessages, "Total messages does not match")
	assert.Equal(t, int64(1), sentence.MessageNumber, "Message number does not match")
	assert.Equal(t, int64(11), sentence.NumberSVsInView, "Number of SVs in view does not match")

	assert.Equal(t, int64(3), sentence.Info[0].SVPRNNumber, "Number of Info[0] SV PRN does not match")
	assert.Equal(t, int64(3), sentence.Info[0].Elevation, "Number of Info[0] Elevation does not match")
	assert.Equal(t, int64(111), sentence.Info[0].Azimuth, "Number of Info[0] Azimuth does not match")
	assert.Equal(t, int64(0), sentence.Info[0].SNR, "Number of Info[0] SNR does not match")

	assert.Equal(t, int64(4), sentence.Info[1].SVPRNNumber, "Number of Info[1] SV PRN does not match")
	assert.Equal(t, int64(15), sentence.Info[1].Elevation, "Number of Info[1] Elevation does not match")
	assert.Equal(t, int64(270), sentence.Info[1].Azimuth, "Number of Info[1] Azimuth does not match")
	assert.Equal(t, int64(0), sentence.Info[1].SNR, "Number of Info[1] SNR does not match")

	assert.Equal(t, int64(6), sentence.Info[2].SVPRNNumber, "Number of Info[2] SV PRN does not match")
	assert.Equal(t, int64(1), sentence.Info[2].Elevation, "Number of Info[2] Elevation does not match")
	assert.Equal(t, int64(10), sentence.Info[2].Azimuth, "Number of Info[2] Azimuth does not match")
	assert.Equal(t, int64(12), sentence.Info[2].SNR, "Number of Info[2] SNR does not match")
}
func TestGLGSVBadSentence(t *testing.T) {
	tests := []struct {
		Input string
		Error string
	}{
		{"$GLGSV,3,1,11.2,03,03,111,00,04,15,270,00,06,01,010,12,13,06,292,00*77", "GLGSV decode number of SVs in view error: 11.2"},
		{"$GLGSV,A3,1,11,03,03,111,00,04,15,270,00,06,01,010,12,13,06,292,00*2A", "GLGSV decode total number of messages error: A3"},
		{"$GLGSV,3,A1,11,03,03,111,00,04,15,270,00,06,01,010,12,13,06,292,00*2A", "GLGSV decode message number error: A1"},
		{"$GLGSV,3,1,11,A03,03,111,00,04,15,270,00,06,01,010,12,13,06,292,00*2A", "GLGSV decode SV prn number error: A03"},
		{"$GLGSV,3,1,11,03,A03,111,00,04,15,270,00,06,01,010,12,13,06,292,00*2A", "GLGSV decode elevation error: A03"},
		{"$GLGSV,3,1,11,03,03,A111,00,04,15,270,00,06,01,010,12,13,06,292,00*2A", "GLGSV decode azimuth error: A111"},
		{"$GLGSV,3,1,11,03,03,111,A00,04,15,270,00,06,01,010,12,13,06,292,00*2A", "GLGSV decode SNR error: A00"},
	}
	for _, tc := range tests {
		_, err := Parse(tc.Input)
		assert.Error(t, err, "Parse error not returned")
		assert.Equal(t, tc.Error, err.Error(), "Incorrect error message")
	}

}

func TestGLGSVWrongSentence(t *testing.T) {
	wrongMsg := "$GPXTE,A,A,4.07,L,N*6D"
	sent := Sentence{}
	sent.parse(wrongMsg)
	msg := GLGSV{Sentence: sent}
	err := msg.parse()
	assert.Error(t, err, "Parse error not returned")
	assert.Equal(t, "GPXTE is not a GLGSV", err.Error(), "Incorrect error message")
}
