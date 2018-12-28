package main

import "testing"

func TestTransform(t *testing.T) {
	cases := map[string]struct {
		arg  string
		want string
	}{
		"a to x": {arg: "a", want: "x"},
		"b to y": {arg: "b", want: "y"},
		"c to z": {arg: "c", want: "z"},
	}

	// setup

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			got := transform(tc.arg)
			if got != tc.want {
				t.Errorf("transform(%q) = %q; want %q", tc.arg, got, tc.want)
			}
		})
	}

	// teardown
}
