package blic

import (
	"testing"
)

func TestBuildStandardLocations(t *testing.T) {
	t.Run("Should build without error", func(t *testing.T) {
		if _, err := buildStandardLocations(); err != nil {
			t.Error(err)
		}
	})
}
