package persist

import (
	"context"
	"fmt"
	"log"
	"time"

	"cloud.google.com/go/firestore"
)

const (
	spendCollection = "spend"
)

type CustomerSpend struct {
	CustomerID string    `firestore:"customer_id,omitempty"`
	Timestamp  time.Time `firestore:"timestamp,omitempty"`
	Amount     float64   `firestore:"amount"`
}

// FirestoreSpendRepository Simple repository to keep track of ubazaar expenditure per client.
type FirestoreSpendRepository struct {
	c *firestore.Client
}

func (f *FirestoreSpendRepository) GetSpendByCustomerId(ctx context.Context, customerId string, limit int) ([]CustomerSpend, error) {
	spendCollection := f.c.Collection(spendCollection)

	docs, err := spendCollection.Where("customer_id", "==", customerId).
		Limit(limit).
		OrderBy("timestamp", firestore.Desc).
		Documents(ctx).GetAll()

	if err != nil {
		return nil, fmt.Errorf("query collection : %v", err)
	}

	var spends []CustomerSpend
	for _, d := range docs {
		var entry CustomerSpend
		d.DataTo(&entry)
		spends = append(spends, entry)
	}

	return spends, nil
}

func (f *FirestoreSpendRepository) SaveSpend(ctx context.Context, customerId string, spend float64) error {
	spends := f.c.Collection(spendCollection)
	if spends == nil {
		return fmt.Errorf("cannot find collection %v", spends)
	}

	cs := &CustomerSpend{
		CustomerID: customerId,
		Timestamp:  time.Now(),
		Amount:     spend,
	}
	if _, _, err := spends.Add(ctx, cs); err != nil {
		return fmt.Errorf("error persisting spend %v", err)
	}

	return nil
}

func NewFirestoreSpendRepository(projectId string) *FirestoreSpendRepository {
	client, err := firestore.NewClient(context.Background(), projectId)
	if err != nil {
		log.Fatalf("unable to create firestore client : %v", err)
	}

	return &FirestoreSpendRepository{c: client}
}
