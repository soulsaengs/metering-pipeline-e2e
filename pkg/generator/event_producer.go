package generator

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"

	"cloud.google.com/go/pubsub"
	"github.com/google/uuid"
)

// EventProducer generates machine events to pubsub based on a given configuration
type EventProducer struct {
	T *pubsub.Topic
}

const (
	NumOfWorkers = 10
)

// Start starts the simulation of machine events.
func (ep EventProducer) Start(fleet Fleet, transitions []Transition) error {

	// Check that the provided traversal is valid
	if err := validateTransitions(transitions); err != nil {
		return fmt.Errorf("running simulation %v", err)
	}

	jobs := make(chan int, fleet.NumOfMachines)
	results := make(chan error, fleet.NumOfMachines)
	for i := 0; i <= NumOfWorkers; i++ {
		// Start traversing the graph
		go startTraversal(fleet, transitions, ep.T, jobs, results)
	}

	for j := 0; j < fleet.NumOfMachines; j++ {
		machineId := rand.Int()
		jobs <- machineId
	}

	close(jobs)

	for a := 1; a <= fleet.NumOfMachines; a++ {
		<-results
	}

	return nil
}

func startTraversal(fleet Fleet, transitions []Transition, topic *pubsub.Topic, machines chan int, result chan error) {

	for m := range machines {

		graph := NewMachineLifecycleGraph()
		// Publish the first state.
		if err := publishToPubSub(&fleet, m, graph.CurrentVertex.State, topic); err != nil {
			result <- fmt.Errorf("publish to pub sub %v", err)
		}

		// Iterate states
		for _, t := range transitions {

			if err := graph.TransitionToState(t.State); err != nil {
				result <- fmt.Errorf("transition to state %v", err)
			}

			// Publish the new state
			if err := publishToPubSub(&fleet, m, graph.CurrentVertex.State, topic); err != nil {
				result <- fmt.Errorf("running simulation %v", err)
			}

			// Stay within that state for a given period of time.
			time.Sleep(time.Duration(t.Duration) * time.Second)
		}

		result <- nil
	}
}

// TODO, this should probably be implemented somewhere else.
func validateTransitions(t []Transition) error {
	testGraph := NewMachineLifecycleGraph()
	for _, s := range t {
		if err := testGraph.TransitionToState(s.State); err != nil {
			return fmt.Errorf("validate transitions %v", err)
		}
	}

	return nil
}

func publishToPubSub(fleet *Fleet, machineID int, state State, t *pubsub.Topic) error {
	log.Printf("publishing event fleet: %v, machineID: %v, state: %v \n", fleet, machineID, state)

	event := newEvent(fleet, machineID, state)
	data, err := json.Marshal(*event)
	if err != nil {
		return fmt.Errorf("pubsub.NewClient: %v", err)
	}

	result := t.Publish(context.Background(), &pubsub.Message{Data: data})

	_, err = result.Get(context.Background())
	if err != nil {
		return fmt.Errorf("get: %v", err)
	}

	return nil
}

func newEvent(fleet *Fleet, machineID int, state State) *Event {
	return &Event{
		ID:            uuid.New().String(),
		FleetId:       fleet.Id,
		RegionId:      fleet.RegionId,
		MachineID:     machineID,
		BillingRegion: fleet.BillingRegion,
		CustomerID:    fleet.CustomerID,
		Location:      fleet.Locations[0],
		EventType:     state,
		Timestamp:     time.Now(),
		MetaData: map[string]interface{}{
			"cores":      fleet.BaseMachineSpec.Cores,
			"memory":     fleet.BaseMachineSpec.Memory,
			"disk_size":  fleet.BaseMachineSpec.Disk,
			"os_version": fleet.BaseMachineSpec.OSVersion,
			"os_family":  fleet.BaseMachineSpec.OSFamily,
		},
	}
}
