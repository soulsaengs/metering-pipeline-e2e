package simluation

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	ubazaar "github.com/soulsaengs/metered-billing-e2e/pkg/verifier/api"

	"github.com/soulsaengs/metered-billing-e2e/pkg/verifier/persist"

	log "github.com/sirupsen/logrus"
	validator "github.com/soulsaengs/metered-billing-e2e/pkg/verifier"
)

func VerifySpend(w http.ResponseWriter, _ *http.Request) {

	p := getProfile()
	ub := initUbazaar()
	sr := initSpendRepository()

	proc := validator.NewTestRunner(ub, sr)
	if err := proc.Run(p); err != nil {
		log.Errorf("test failed : %v", err)
	}

	fmt.Fprintf(w, "Finished verifing spend")
}

func getProfile() *validator.Profile {
	bytes, err := ioutil.ReadFile("test-profiles.json")
	if err != nil {
		log.Fatalf("unable to get client config : %v\n", err)
	}

	cfg := &validator.TestProfiles{}
	if err := json.Unmarshal(bytes, cfg); err != nil {
		log.Fatalf("unable to unmarshall client file : %v\n", err)
	}

	return &cfg.Profiles[0] // Currently we default to the last profile.
}

func initUbazaar() *ubazaar.API {
	audience := os.Getenv("AUTH_AUDIENCE")
	uri := os.Getenv("UBAZAAR_URI")

	return ubazaar.NewAPI(audience, uri)
}

func initSpendRepository() *persist.FirestoreSpendRepository {
	projectID := os.Getenv("PROJECT_ID")
	return persist.NewFirestoreSpendRepository(projectID)
}
