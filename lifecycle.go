package simluation

import (
	"fmt"
)

// MachineLifecycleGraph Simple directed graph that represents a machine lifecycle.
type MachineLifecycleGraph struct {
	// Vertices are the set of states in the Graph.
	Vertices map[State]*Vertex
	// CurrentState are the set of states in the Graph.
	CurrentVertex *Vertex
}

type Vertex struct {
	State State
	Edges []Edge
}

type Edge struct {
	From *Vertex
	To   *Vertex
}

type State string

const (
	Start            = "start"
	Stop             = "stop"
	ScaleUp          = "scale_up"
	ScaleDown        = "scale_down"
	SpecChange       = "spec_change"
	HeartbeatRunning = "heartbeat_running"
	HeartbeatStopped = "heartbeat_stopped"
)

// addEdge adds a directed edge from and to the origin vertex.
func (mlg *MachineLifecycleGraph) addEdge(from *Vertex, to *Vertex) error {
	_, present := mlg.Vertices[from.State]
	if !present {
		return fmt.Errorf("from vertex %v does not exist", from)
	}
	_, present = mlg.Vertices[to.State]
	if !present {
		return fmt.Errorf("to vertex %v does not exist", to)
	}

	existing := mlg.Vertices[from.State]
	existing.Edges = append(existing.Edges, Edge{from, to})

	return nil
}

// addVertex Adds a new vertex to the Graph
func (mlg *MachineLifecycleGraph) addVertex(vertex *Vertex) {
	_, present := mlg.Vertices[vertex.State]
	// If the vertex already exists remove it.
	if present {
		delete(mlg.Vertices, vertex.State)
	}

	mlg.Vertices[vertex.State] = vertex
}

// NextStates returns the next states the machine can transition to based on the given state.
func (mlg *MachineLifecycleGraph) NextStates(state State) ([]State, error) {
	source := mlg.Vertices[state]
	destinations := make([]State, 0, 10)
	for _, e := range source.Edges {
		destinations = append(destinations, e.To.State)
	}

	return destinations, nil
}

// TransitionToState Moves the current state to the next desired state. An error is thrown there are no connecting edges.
func (mlg *MachineLifecycleGraph) TransitionToState(state State) error {
	states, err := mlg.NextStates(mlg.CurrentVertex.State)
	if err != nil {
		return fmt.Errorf("error transition to state %v", err)
	}
	if !containsState(states, state) {
		return fmt.Errorf("cannot transition to state %v", state)
	}

	mlg.CurrentVertex = mlg.Vertices[state]

	return nil
}

func containsState(states []State, state State) bool {
	contains := false
	for _, s := range states {
		if s == state {
			contains = true
		}
	}

	return contains
}

// NewMachineLifecycleGraph Generates a new machine lifecycle graph with the scale up state as the current vertex.
func NewMachineLifecycleGraph() *MachineLifecycleGraph {

	g := &MachineLifecycleGraph{make(map[State]*Vertex), nil}

	start := &Vertex{State: Start}
	stop := &Vertex{State: Stop}
	scaleUp := &Vertex{State: ScaleUp}
	scaleDown := &Vertex{State: ScaleDown}
	specChange := &Vertex{State: SpecChange}
	hbr := &Vertex{State: HeartbeatRunning}
	hbs := &Vertex{State: HeartbeatStopped}

	g.addVertex(start)
	g.addVertex(stop)
	g.addVertex(scaleUp)
	g.addVertex(scaleDown)
	g.addVertex(specChange)
	g.addVertex(hbr)
	g.addVertex(hbs)

	g.addEdge(scaleUp, stop)
	g.addEdge(scaleUp, hbr)
	g.addEdge(start, stop)
	g.addEdge(start, hbr)
	g.addEdge(hbr, stop)
	g.addEdge(stop, start)
	g.addEdge(stop, specChange)
	g.addEdge(stop, scaleDown)
	g.addEdge(stop, hbs)
	g.addEdge(specChange, start)
	g.addEdge(hbs, start)
	g.addEdge(hbs, scaleDown)

	g.CurrentVertex = scaleUp

	return g
}
