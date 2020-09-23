package mock

import "time"

type Type string

// Machine represents a machine document
type Machine struct {
	MachineID       string    `json:"machine_id" firestore:"machine_id"`
	Billable        bool      `json:"billable" firestore:"billable"`
	RecentTimestamp time.Time `json:"recent_timestamp" firestore:"recent_timestamp"`
}

// Event represents a metering event
type Event struct {
	ID         string                 `json:"id" validate:"required"`
	FleetId    string                 `json:"fleet_id" validate:"required"`
	RegionId   string                 `json:"region_id" validate:"required"`
	MachineID  int                    `json:"machine_id" validate:"required"`
	CustomerID string                 `json:"customer_id" validate:"required"`
	Location   string                 `json:"location" validate:"required"`
	EventType  State                  `json:"type" validate:"required,oneof=scale_up scale_down spec_change start stop heartbeat_running heartbeat_stopped"`
	Timestamp  time.Time              `json:"timestamp" validate:"required"`
	MetaData   map[string]interface{} `json:"metadata" validate:"gt=0,dive,required,gt=0"`
}
