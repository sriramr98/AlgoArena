package utils

import (
	"reflect"
	"testing"
)

func TestMapKeysSorted(t *testing.T) {
	tests := []struct {
		name  string
		input map[string]interface{}
		want  []string
	}{
		{
			name:  "empty map",
			input: map[string]interface{}{},
			want:  []string{},
		},
		{
			name: "map with string values",
			input: map[string]interface{}{
				"c": "value3",
				"a": "value1",
				"b": "value2",
			},
			want: []string{"a", "b", "c"},
		},
		{
			name: "map with int values",
			input: map[string]interface{}{
				"three": 3,
				"one":   1,
				"two":   2,
			},
			want: []string{"one", "three", "two"},
		},
		{
			name: "map with mixed value types",
			input: map[string]interface{}{
				"z": 26,
				"m": "middle",
				"a": true,
			},
			want: []string{"a", "m", "z"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MapKeysSorted(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapKeysSorted() = %v, want %v", got, tt.want)
			}
		})
	}
}
