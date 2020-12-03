package vbox

import (
	"reflect"
	"testing"
)

func TestControlOptions_slice(t *testing.T) {
	var tests = []struct {
		desc string
		i    ControlOptions
		out  []string
	}{
		{
			desc: "no options",
		},
		{
			desc: "1 NIC",
			i: ControlOptions{
				Nics: []Nic{
					Nic{Mode: Bridged, Iface: "eth0"}},
			},
			out: []string{
				"nic1",
				"bridged",
				"eth0",
			},
		},
		{
			desc: "2 NIC",
			i: ControlOptions{
				Nics: []Nic{
					Nic{Mode: Bridged, Iface: "eth0"},
					Nic{Mode: NAT, Iface: "eth0"}},
			},
			out: []string{
				"nic1",
				"bridged",
				"eth0",
				"nic2",
				"nat",
			},
		},
		{
			desc: "all options",
			i: ControlOptions{
				Nics: []Nic{
					Nic{Mode: Bridged, Iface: "eth0"},
					Nic{Mode: NAT}},
			},
			out: []string{
				"nic1",
				"bridged",
				"eth0",
				"nic2",
				"nat",
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
