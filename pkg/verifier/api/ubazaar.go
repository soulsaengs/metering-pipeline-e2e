package ubazaar

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"google.golang.org/api/idtoken"
)

type API struct {
	audience   string
	ubazaarUri string
	client     *http.Client
}

func (u *API) Spend(serviceId string, serviceAccountId string) (*Spend, error) {

	resource := fmt.Sprintf("%s/v1/services/%s/customers/%s/spend", u.ubazaarUri, serviceId, serviceAccountId)
	resp, err := u.client.Get(resource)

	if err != nil {
		return nil, fmt.Errorf("get spend from ubazaar client : %v", err)
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("received status code: %d from ubazaar", resp.StatusCode)
	}

	spend := &Spend{}

	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(spend)

	if err != nil {
		return nil, fmt.Errorf("cannot decode response %v", err)
	}

	return spend, nil
}

func NewAPI(audience string, uri string) *API {

	var err error
	client, err := idtoken.NewClient(context.Background(), audience)
	if err != nil {
		log.Fatalf("unable to create ubazaar client : %v\n", err)
	}

	return &API{
		audience:   audience,
		ubazaarUri: uri,
		client:     client,
	}
}
