package validator

import (
	"context"
	"fmt"
	"strconv"

	ubazaar "github.com/soulsaengs/metered-billing-e2e/pkg/verifier/api"
	"github.com/soulsaengs/metered-billing-e2e/pkg/verifier/persist"
)

const (
	serviceId = "MP"
)

// TestProcessor does the heavy lifting when running a spend anomaly test.
type TestRunner struct {
	u *ubazaar.API
	f *persist.FirestoreSpendRepository
}

func (t *TestRunner) Run(p *Profile) error {

	// current spend and previous spend (yesterday).
	currentSpend, err := t.getCurrentSpendFromUbazaar(serviceId, p.ServiceAccountId)
	if err != nil {
		return fmt.Errorf("unable to compute the daily spend : %v", err)
	}

	spends, err := t.f.GetSpendByCustomerId(context.Background(), p.ServiceAccountId, 30)
	if err != nil {
		return fmt.Errorf("unable to get spend history for customerId : %v, %v", p.ServiceAccountId, err)
	}

	if len(spends) > 0 {
		lastSpend := spends[0].Amount
		if lastSpend != 0 {
			// The difference between the current spend and the last spend gives you your daily amount since the function runs once per day.
			dailySpend := currentSpend - lastSpend

			detector := NewDailySpendAnomalyDetector(p, nil)
			if err := detector.Detect(dailySpend); err != nil {
				return fmt.Errorf("detected anomaly on account %+v with daily spend %+v", p, currentSpend)
			}
		}
	}

	// Save the current spend so we can use it on the next run.
	if err := t.f.SaveSpend(context.Background(), p.ServiceAccountId, currentSpend); err != nil {
		return fmt.Errorf("unable to persist current spend : %v", err)
	}

	return nil
}

// Get the current spend amount processed by ubazaar, will be useful for determining the daily cost which is obtained by doing the delta between now and the previous run.
func (t *TestRunner) getCurrentSpendFromUbazaar(serviceId string, customerId string) (float64, error) {
	spend, err := t.u.Spend(serviceId, customerId)
	if err != nil {
		return 0, fmt.Errorf("error calling ubazaar Spend API : %v\n", err)
	}

	dailySpend, err := strconv.ParseFloat(spend.SpendAmount, 64)
	if err != nil {
		return 0, fmt.Errorf("error parsing SpendAmount from Ubazaar : %v\n", err)
	}

	return dailySpend, nil
}

func NewTestRunner(u *ubazaar.API, f *persist.FirestoreSpendRepository) *TestRunner {
	return &TestRunner{
		u: u,
		f: f,
	}
}
