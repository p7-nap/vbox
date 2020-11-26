package vbox

import (
	"reflect"
	"testing"
)

func TestStartOptions_slice(t *testing.T) {
	var tests = []struct {
		desc string
		i    StartOptions
		out  []string
	}{
		{
			desc: "no options",
		},
		{
			desc: "type gui",
			i: StartOptions{
				Type: GUI,
			},
			out: []string{
				"--type=gui",
			},
		},
		{
			desc: "type sdl",
			i: StartOptions{
				Type: SDL,
			},
			out: []string{
				"--type=sdl",
			},
		},
		{
			desc: "type headless",
			i: StartOptions{
				Type: Headless,
			},
			out: []string{
				"--type=headless",
			},
		},
		{
			desc: "type separate",
			i: StartOptions{
				Type: Seperate,
			},
			out: []string{
				"--type=separate",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			if want, got := tt.out, tt.i.slice(); !reflect.DeepEqual(want, got) {
				t.Fatalf("unexpected slices:\n- want: %v\n-  got: %v",
					want, got)
			}
		})
	}
}
