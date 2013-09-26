package invnorm

import (
	"math"
	"testing"
)

func TestInvCumulativeNormalDistribution(t *testing.T) {
	delta := 0.00000001
	var tests = []struct {
		p float64
		x float64
	}{
		{
			0.00000000001,
			-6.70602314994,
		},
		{
			0.2,
			-0.841621232727,
		},
		{
			0.5,
			0,
		},
		{
			0.6,
			0.25334710286,
		},
		{
			0.999999999999,
			7.03448690268,
		},
	}
	for _, tt := range tests {
		if actualX := InvCumulativeNormalDistribution(tt.p); math.Abs((actualX - tt.x)) > delta {
			t.Errorf("expected %.14f actual %.14f", tt.x, actualX)
		}
	}
}
