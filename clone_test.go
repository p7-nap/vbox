package vbox

import (
	"fmt"
	"reflect"
	"testing"
)

// func TestCloneVMOK(t *testing.T) {
// 	options := CloneOptions{
// 		Basefolder: ".",
// 		Groups:     "group",
// 		Mode:       "machine",
// 		Name:       "newVM",
// 		Options:    "Link, KeepAllMACs, KeepNATMACs, KeepDiskNames, KeepDiskNames",
// 		Register:   true,
// 		Snapshot:   "VmSnap",
// 		UUID:       "UUIDnew",
// 	}
// }

func TestCloneOptions_slice(t *testing.T) {
	var tests = []struct {
		desc string
		i    CloneOptions
		out  []string
		err  error
	}{
		{
			desc: "all options",
			i: CloneOptions{
				Basefolder: ".",
				Groups:     "group",
				Mode:       "machine",
				Name:       "newVM",
				Options:    "Link, KeepAllMACs, KeepNATMACs, KeepDiskNames, KeepDiskNames",
				Register:   true,
				Snapshot:   "vmSnap",
				UUID:       "UUIDnew",
			},
			out: []string{
				"--basefolder=.",
				"--groups=group",
				"--mode=machine",
				"--name=newVM",
				"--options=Link, KeepAllMACs, KeepNATMACs, KeepDiskNames, KeepDiskNames",
				"--register",
				"--snapshot=vmSnap",
				"--uuid=UUIDnew"},
			err: nil,
		},
		{
			desc: "machineandchildren mode with snapshot",
			i: CloneOptions{
				Mode:     "machineandchildren",
				Snapshot: "vmSnap",
			},
			out: []string{
				"--mode=machineandchildren",
				"--snapshot=vmSnap"},
			err: nil,
		},
		{
			desc: "machineandchildren mode without snapshot",
			i: CloneOptions{
				Mode:     "machineandchildren",
				Register: true,
			},
			out: []string{
				"--mode=machineandchildren",
			},
			err: fmt.Errorf("machineandchildren mode parameter only available with --snapshot"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			want := tt.out
			wanterr := tt.err
			got, goterr := tt.i.slice()
			if !reflect.DeepEqual(want, got) {
				t.Fatalf("unexpected slices:\n- want: %v\n-  got: %v",
					want, got)
			}
			if !reflect.DeepEqual(wanterr, goterr) {
				t.Fatalf("unexpected slices:\n- want: %v\n-  got: %v",
					wanterr, goterr)
			}
		})
	}
}
