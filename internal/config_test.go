package goscawa

import "testing"

func TestEnvironmentString(t *testing.T) {

	tests := []struct {
		input environment
		want  string
	}{
		{dev, "dev"},
		{test, "testing"},
		{staging, "staging"},
		{prod, "prod"},
	}
	for _, tc := range tests {
		if tc.input.String() != tc.want {
			t.Errorf("Enum environment: got %v, want %s", tc.input, tc.want)
		}
	}

}

func TestEnvironmentFromString(t *testing.T) {

	tests := []struct {
		want  string
		input environment
	}{
		{"dev", dev},
		{"testing", test},
		{"staging", staging},
		{"prod", prod},
	}
	for _, tc := range tests {
		_, err := FromString(tc.want)
		if err != nil {
			t.Errorf("Oups '%s' not possible : %v", tc.want, err)
		}
	}

	_, err := FromString("dev-in-prod")
	if err != nil {
		if err.Error() != "dev-in-prod is not a possible value" {
			t.Errorf("Incorrect error message for possible value")
		}
	}

}
