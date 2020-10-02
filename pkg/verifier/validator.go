package validator

import (
	"github.com/lytics/anomalyzer"
)

const (
	// Magic Number
	sensitivity = 0.1
	// A result that yields 95% and above will be considered abnormal.
	acceptanceThreshold = 0.95
)

var (
	// we use a mix of different probabilistic tests. Each test will be executed sequentially and the weighted average of those tests will yield the probability.
	methods = []string{"diff", "fence", "highrank", "lowrank", "magnitude"}
)

// TimeSerieValidator Takes a daily spend amount and compares it against a normal cost curve.
type TimeSerieValidator struct {
	anomalyzer anomalyzer.Anomalyzer
}

// Detect Applies anomaly detection on the given data point returns true if data point seems abnormal, false otherwise.
func (t *TimeSerieValidator) ValidateDataPoint(datapoint float64) bool {
	confidence := t.anomalyzer.Push(datapoint)
	return confidence > acceptanceThreshold
}

// NewTimeSerieValidator creates a new daily spend validator. This validation function takes settings from previous cost curves that are deemed to be acceptable.
// We define a higher bound, lower bound and data points collected from previous billing runs. Based on those parameters, we can then test simulated daily spend for anomalies.
// The returned result will be a confidence score in percentage. The higher the confidence score, the more likely the daily spend appears normal.
func NewTimeSerieValidator(activeSize int, upperBound float64, lowerBound float64, data []float64) *TimeSerieValidator {

	conf := &anomalyzer.AnomalyzerConf{
		Sensitivity: sensitivity,
		UpperBound:  upperBound,
		LowerBound:  lowerBound,
		ActiveSize:  activeSize,
		Methods:     methods,
	}

	a, _ := anomalyzer.NewAnomalyzer(conf, data)

	return &TimeSerieValidator{anomalyzer: a}
}
