package validator

type TestProfiles struct {
	Profiles []Profile `json:"profiles" validate:"required"`
}

type Profile struct {
	ServiceAccountId string    `json:"service_account_id" validate:"required"`
	ActiveSize       int       `json:"active_size" validate:"required"`
	UpperBound       float64   `json:"upper_bound" validate:"required"`
	LowerBound       float64   `json:"lower_bound" validate:"required"`
	Data             []float64 `json:"data" validate:"required"`
}
