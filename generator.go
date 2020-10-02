package simluation

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/pubsub"
	"github.com/soulsaengs/metered-billing-e2e/pkg/generator"
)

func GenerateEvents(w http.ResponseWriter, _ *http.Request) {
	bytes, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatalf("unable to read file : %v", err)
	}

	cfg := &generator.ProducerConfigs{}
	if err := json.Unmarshal(bytes, cfg); err != nil {
		log.Fatalf("unable to unmarshall file : %v", err)
	}

	projectID := os.Getenv("PROJECT_ID")
	client, err := pubsub.NewClient(context.Background(), projectID)
	if err != nil {
		log.Fatalf("unable to create pubsub client : %v", err)
	}

	topic := client.Topic(cfg.TopicId)
	producer := &generator.EventProducer{
		T: topic,
	}

	if err := producer.Start(cfg.Fleet, cfg.Transitions); err != nil {
		log.Fatalf("Producer error : %v", err)
	}

	fmt.Fprintf(w, "Finished generating events.")
}
