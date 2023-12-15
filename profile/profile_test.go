package profile

import (
	"reflect"
	"testing"
)

func TestNewProfile(t *testing.T) {
	tests := []struct {
		name string
		want *Profile
	}{
		{name: "通常ケース", want: &Profile{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewProfile(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewProfile() = %v, want %v", got, tt.want)
			}
		})
	}
}
