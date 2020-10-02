package ubazaar

type Spend struct {
	ServiceId         string `json:"serviceId"`
	ServiceCustomerId string `json:"serviceCustomerId"`
	Currency          string `json:"currency"`
	SpendAmount       string `json:"spendAmount"`
}
