package validator

import (
	log "github.com/sirupsen/logrus"
)

type Notifier interface {
	Notify(interface{})
}

type TestFailureNotifier struct {
	name string
}

// TODO : We should alert to a proper channel rather than logging it.
func (t *TestFailureNotifier) Notify(i interface{}) {
	log.Errorf("Test failure %+v", i)
}
