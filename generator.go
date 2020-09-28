package simluation

import (
	"cloud.google.com/go/pubsub"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	client *pubsub.Client
	cfg *ProducerConfigs
)

func init() {
	bytes, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatalf("unable to read file : %v", err)
	}

	cfg = &ProducerConfigs{}
	if err := json.Unmarshal(bytes, cfg); err != nil {
		log.Fatalf("unable to unmarshall file : %v", err)
	}

	client, err = pubsub.NewClient(context.Background(), cfg.ProjectId)
	if err != nil {
		log.Fatalf("unable to create pubsub client : %v", err)
	}
}

func GenerateEvents(w http.ResponseWriter, _ *http.Request) {
	topic := client.Topic(cfg.TopicId)
	producer := &EventProducer{
		T: topic,
	}

	if err := producer.Start(cfg.Fleet, cfg.Transitions); err != nil {
		log.Fatalf("Producer error : %v", err)
	}

	fmt.Fprintf(w, "Finished generating events.")
}
