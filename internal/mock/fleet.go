package mock

type MachineSpecs struct {
	Cores     int    `json:"cores" validate:"required"`
	Memory    int    `json:"memory" validate:"required"`
	Disk      int    `json:"disk" validate:"required"`
	OSVersion string `json:"os_version" validate:"required"`
	OSFamily  string `json:"os_family" validate:"required"`
}

type Fleet struct {
	Id              string       `json:"fleet_id" validate:"required"`
	RegionId        string       `json:"region_id" validate:"required"`
	Locations       []string     `json:"locations" validate:"required"`
	BillingRegion   string       `json:"billing_region" validate:"required"`
	CustomerID      string       `json:"customer_id" validate:"required"`
	NumOfMachines   int          `json:"num_of_machines" validate:"required"`
	BaseMachineSpec MachineSpecs `json:"base_machine_specs" validate:"required"`
}

// ProducerConfigs are the base configurations for the event producer.
type ProducerConfigs struct {
	Fleet       Fleet        `json:"fleet" validate:"required"`
	Transitions []Transition `json:"transitions" validate:"required"`
	ProjectId   string       `json:"project_id" validate:"required"`
	TopicId     string       `json:"topic_id" validate:"required"`
}

// Transition defines a state, and the duration of time that we should spend in that state.
type Transition struct {
	State    State `json:"state" validate:"required"`
	Duration int   `json:"duration" validate:"required"`
}
