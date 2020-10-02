package persist

import (
	"context"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirestoreSpendRepository_GetSpendByCustomerId(t *testing.T) {
	r := NewFirestoreSpendRepository("mp-metered-billing-dev-273014")
	spends, err := r.GetSpendByCustomerId(context.Background(), "ASID:7777777", 30)
	if err != nil {
		log.Fatalf("error : %v", err)
	}

	assert.NoError(t, err)
	assert.NotEmpty(t, spends)
}

func TestFirestoreSpendRepository_SaveSpend(t *testing.T) {
	r := NewFirestoreSpendRepository("mp-metered-billing-dev-273014")
	err := r.SaveSpend(context.Background(), "ASID:7777777", 99.99)

	assert.NoError(t, err)

}
