package blic

import (
	"testing"
)

func TestStandardGameLocationSetup(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Error(r)
		}
	}()

	mustBuildStandardLocations()
}
