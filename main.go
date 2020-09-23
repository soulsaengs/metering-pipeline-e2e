package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/soulsaengs/metering-pipeline-e2e/internal/mock"

	"cloud.google.com/go/pubsub"
)

func main() {
	config := flag.String("config", "config.json", "a configuration file")
	flag.Parse()

	bytes, err := ioutil.ReadFile(*config)
	if err != nil {
		log.Fatalf("unable to read file : %v", err)
	}

	cfg := &mock.ProducerConfigs{}
	if err := json.Unmarshal(bytes, cfg); err != nil {
		log.Fatalf("unable to unmarshall file : %v", err)
	}

	client, err := pubsub.NewClient(context.Background(), cfg.ProjectId)
	if err != nil {
		fmt.Printf("error create client %v", err)
	}

	topic := client.Topic(cfg.TopicId)
	producer := &mock.EventProducer{
		T: topic,
	}

	if err := producer.Start(cfg.Fleet, cfg.Transitions); err != nil {
		log.Fatalf("Producer error : %v", err)
	}

}
