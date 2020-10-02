package validator

import (
	"fmt"
)

type AnomalyDetector interface {
	Detect(s float64) error
}

type DailySpendAnomalyDetector struct {
	profile   *Profile
	validator *TimeSerieValidator
	notifier  Notifier
}

func (d *DailySpendAnomalyDetector) Detect(s float64) error {
	if abnormal := d.validator.ValidateDataPoint(s); abnormal {
		err := fmt.Errorf("anomaly detected on Spend %v", s)
		d.notifier.Notify(err)
		return err
	}

	return nil
}

func NewDailySpendAnomalyDetector(p *Profile, notifier Notifier) *DailySpendAnomalyDetector {
	validator := NewTimeSerieValidator(
		p.ActiveSize,
		p.UpperBound,
		p.LowerBound,
		p.Data)

	return &DailySpendAnomalyDetector{profile: p, notifier: notifier, validator: validator}
}
