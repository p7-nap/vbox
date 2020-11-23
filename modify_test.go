package vbox

import (
	"reflect"
	"testing"
)

func TestModifyOptions_slice(t *testing.T) {
	var tests = []struct {
		desc string
		i    ModifyOptions
		out  []string
	}{
		{
			desc: "no options",
		},
		{
			desc: "Name",
			i: ModifyOptions{
				Name: "Name",
			},
			out: []string{
				"--name=Name",
			},
		},
		{
			desc: "RAM",
			i: ModifyOptions{
				RAM: 128,
			},
			out: []string{
				"--memory=128",
			},
		},
		{
			desc: "CPU",
			i: ModifyOptions{
				CPUs: 2,
			},
			out: []string{
				"--cpus=2",
			},
		},
		{
			desc: "VRAM",
			i: ModifyOptions{
				VRAM: 128,
			},
			out: []string{
				"--vram=128",
			},
		},
		{
			desc: "OSType",
			i: ModifyOptions{
				OSType: Linux64,
			},
			out: []string{
				"--ostype=linux_64",
			},
		},
		{
			desc: "1 NIC",
			i: ModifyOptions{
				Nics: []Nic{
					Nic{Mode: Bridged, Iface: "eth0"}},
			},
			out: []string{
				"--nic1=bridged",
				"--bridgeadapter1=eth0",
			},
		},
		{
			desc: "2 NIC",
			i: ModifyOptions{
				Nics: []Nic{
					Nic{Mode: Bridged, Iface: "eth0"},
					Nic{Mode: NAT, Iface: "eth0"}},
			},
			out: []string{
				"--nic1=bridged",
				"--bridgeadapter1=eth0",
				"--nic2=nat",
			},
		},
		{
			desc: "all options",
			i: ModifyOptions{
				Name:   "Name",
				RAM:    128,
				CPUs:   2,
				VRAM:   128,
				OSType: Linux64,
				Nics: []Nic{
					Nic{Mode: Bridged, Iface: "eth0", Promisc: AllowAll},
					Nic{Mode: NAT, Iface: "eth0"}},
			},
			out: []string{
				"--name=Name",
				"--memory=128",
				"--ostype=linux_64",
				"--vram=128",
				"--cpus=2",
				"--nic1=bridged",
				"--bridgeadapter1=eth0",
				"--nicpromisc1=allow-all",
				"--nic2=nat",
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
