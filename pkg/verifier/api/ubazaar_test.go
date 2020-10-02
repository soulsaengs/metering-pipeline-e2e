package ubazaar

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAPI_Spend(t *testing.T) {
	u := NewAPI("349998405338-s59jhmvf6m77cc447idb42n8mfvcfkhk.apps.googleusercontent.com", "https://api.test.billing.corp.unity3d.com")
	s, err := u.Spend("MP", "ASID:777777")
	assert.NoError(t, err)
	assert.NotNil(t, s)
}
