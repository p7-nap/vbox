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
				"--nic1=bridged",
				"--bridgeadapter1=eth0",
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
				"--nic1=bridged",
				"--bridgeadapter1=eth0",
				"--nic2=nat",
			},
		},
		{
			desc: "PF NIC",
			i: ControlOptions{
				Nics: []Nic{
					Nic{Mode: NAT,
						Iface:       "eth0",
						PortForward: true,
						ForwardRules: []ForwardRule{
							ForwardRule{
								HostIP:    "192.168.0.1",
								HostPort:  8080,
								GuestIP:   "192.168.1.1",
								GuestPort: 80,
								Protocol:  TCP,
								Name:      "Forwardtest",
							},
						},
					},
				},
			},
			out: []string{
				"--nic1=nat",
				"--natpf1=Forwardtest,tcp,192.168.0.1,8080,192.168.1.1,80",
			},
		},
		{
			desc: "all options",
			i: ControlOptions{
				Nics: []Nic{
					Nic{Mode: Bridged, Iface: "eth0", Promisc: AllowAll},
					Nic{Mode: NAT, Iface: "eth0"}},
			},
			out: []string{
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
