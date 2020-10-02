package validator

import (
	"testing"
)

// Test this.
func TestTestProcessor_Process(t *testing.T) {

}

func profile() *Profile {
	return &Profile{
		ServiceAccountId: "ASID:777777",
		ActiveSize:       0,
		UpperBound:       105,
		LowerBound:       80,
		Data:             []float64{84.7, 87.9, 90.7, 101.90, 89.12, 95.6, 90.1, 99.6, 102.1, 99.1, 101.2, 94.2, 91.2, 102.10, 85.5, 86.6},
	}
}
