package main

import (
	"reflect"
	"testing"
)

func TestCleanInput(t *testing.T) {
	tests := map[string]struct{
		input string
		want []string
	}{
		"whiteSpace": {input: " testing whiteSpace ", want: []string{"testing", "whitespace"}},
		"lowerCase": {input: "LOWER CASE TEST", want: []string{"lower", "case", "test"}},

	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T){
			got := cleanInput(tc.input)
			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("expected: %#v, got: %#v", tc.want, got)}
		})
	
	}

}
