package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDailySpendAnomalyDetector_Detect(t *testing.T) {
	acc := &Profile{
		ServiceAccountId: "ASID:777777",
		ActiveSize:       1,
		UpperBound:       112.10,
		LowerBound:       70,
		Data:             []float64{84.7, 87.9, 90.7, 101.90, 89.12, 95.6, 90.1, 99.6, 102.1, 99.1, 101.2, 94.2, 91.2, 102.10, 85.5, 86.6},
	}

	d := NewDailySpendAnomalyDetector(acc, &TestFailureNotifier{})
	err := d.Detect(150)
	assert.Error(t, err, "Anomaly detected on Spend %v", 300)

	d2 := NewDailySpendAnomalyDetector(acc, &TestFailureNotifier{})
	err = d2.Detect(103)
	assert.NoError(t, err)
}
