package mock_test

import (
	"testing"

	"github.com/soulsaengs/metering-pipeline-e2e/internal/mock"

	"github.com/stretchr/testify/assert"
)

func TestMachineLifecycleGraph_NextStates(t *testing.T) {
	mlg := mock.NewMachineLifecycleGraph()

	states, _ := mlg.NextStates(mock.Start)
	expected := []mock.State{mock.Stop, mock.HeartbeatRunning}
	assert.ElementsMatch(t, states, expected)

	states, _ = mlg.NextStates(mock.Stop)
	expected = []mock.State{mock.ScaleDown, mock.SpecChange, mock.Start, mock.HeartbeatStopped}
	assert.ElementsMatch(t, states, expected)
}

func TestMachineLifecycleGraph_SetState(t *testing.T) {
	mlg := mock.NewMachineLifecycleGraph()

	err := mlg.TransitionToState(mock.SpecChange)
	assert.Error(t, err)
}
