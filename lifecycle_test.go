package simluation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMachineLifecycleGraph_NextStates(t *testing.T) {
	mlg := NewMachineLifecycleGraph()

	states, _ := mlg.NextStates(Start)
	expected := []State{Stop, HeartbeatRunning}
	assert.ElementsMatch(t, states, expected)

	states, _ = mlg.NextStates(Stop)
	expected = []State{ScaleDown, SpecChange, Start, HeartbeatStopped}
	assert.ElementsMatch(t, states, expected)
}

func TestMachineLifecycleGraph_SetState(t *testing.T) {
	mlg := NewMachineLifecycleGraph()

	err := mlg.TransitionToState(SpecChange)
	assert.Error(t, err)
}
