package vbox

import (
	"reflect"
	"testing"
)

func TestSnapshotTakeOptions_slice(t *testing.T) {
	var tests = []struct {
		desc string
		i    SnapshotTakeOptions
		out  []string
	}{
		{
			desc: "no options",
		},
		{
			desc: "description",
			i: SnapshotTakeOptions{
				Description: "this_is_a_snapshot",
			},
			out: []string{
				"--description=this_is_a_snapshot",
			},
		},
		{
			desc: "live",
			i: SnapshotTakeOptions{
				Live: true,
			},
			out: []string{
				"--live",
			},
		},
		{
			desc: "uniquename",
			i: SnapshotTakeOptions{
				Uniquename: "uniquename",
			},
			out: []string{
				"--uniquename=uniquename",
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

func TestSnapshotEditOptions_slice(t *testing.T) {
	var tests = []struct {
		desc string
		i    SnapshotEditOptions
		out  []string
	}{
		{
			desc: "no options",
		},
		{
			desc: "description",
			i: SnapshotEditOptions{
				Description: "this_is_a_snapshot",
			},
			out: []string{
				"--description=this_is_a_snapshot",
			},
		},
		{
			desc: "new name",
			i: SnapshotEditOptions{
				NewName: "newname",
			},
			out: []string{
				"--name=newname",
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

func TestSnapshotListOptions_slice(t *testing.T) {
	var tests = []struct {
		desc string
		i    SnapshotListOptions
		out  []string
	}{
		{
			desc: "no options",
		},
		{
			desc: "with listmode details",
			i: SnapshotListOptions{
				ListMode: Details,
			},
			out: []string{
				"--details",
			},
		},
		{
			desc: "with listmode machinereadable",
			i: SnapshotListOptions{
				ListMode: Machinereadable,
			},
			out: []string{
				"--machinereadable",
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
