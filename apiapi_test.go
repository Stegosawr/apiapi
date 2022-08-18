package apiapi

import "testing"

func TestGetByCode(t *testing.T) {
	tests := []struct {
		Name     string
		In       string
		CodeType string
	}{
		{
			Name:     "SCode",
			In:       "FIGURE-131176",
			CodeType: string(CodeTypeS),
		}, {
			Name:     "GCode",
			In:       "FIG-DOL-7508",
			CodeType: string(CodeTypeG),
		},
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			details, err := GetItemByCode(CodeType(tt.CodeType), tt.In)
			if err != nil {
				t.Error(err)
			}

			if !details.RSuccess {
				t.Errorf("Got: %v - want: %v", "false", "true")
			}
		})
	}
}

func TestGetByKeywords(t *testing.T) {
	tests := []struct {
		Name     string
		Keywords string
		Want     int
	}{
		{
			Name:     "Default",
			Keywords: "B-STYLE%20Rent-A-Girlfriend",
			Want:     3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			details, err := GetItemsByKeywords(tt.Keywords)
			if err != nil {
				t.Error(err)
			}

			if len(details) < tt.Want {
				t.Errorf("Got: %v - want: %v", len(details), tt.Want)
			}
		})
	}
}

func TestGetCurrencyLayer(t *testing.T) {
	t.Run("Default Request Currency Exchange Rate", func(t *testing.T) {
		cLayer, err := GetCurrencyLayer()
		if err != nil {
			t.Error(err)
		}

		if !cLayer.Success {
			t.Errorf("Got: %v - want: %v", "false", "true")
		}
	})
}
