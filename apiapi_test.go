package apiapi

import "testing"

func TestGetBySCode(t *testing.T) {
	t.Run("Default test", func(t *testing.T) {
		details, err := GetItemBySCode("FIGURE-131176")
		if err != nil {
			t.Error(err)
		}

		if !details.RSuccess {
			t.Errorf("Got: %v - want: %v", "false", "true")
		}
	})
}
