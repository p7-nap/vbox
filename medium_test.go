package vbox

import (
	"reflect"
	"testing"
)

func TestCreateMediumOptions_slice(t *testing.T) {
	var tests = []struct {
		desc string
		i    CreateMediumOptions
		out  []string
	}{
		{
			desc: "no options",
		},
		{
			desc: "Type",
			i: CreateMediumOptions{
				Type: DVDMedium,
			},
			out: []string{
				"dvd",
			},
		},
		{
			desc: "Filename",
			i: CreateMediumOptions{
				Filename: "test/test.vdi",
			},
			out: []string{
				"--filename=test/test.vdi",
			},
		},
		{
			desc: "Size",
			i: CreateMediumOptions{
				Size: 128,
			},
			out: []string{
				"--size=128",
			},
		},
		{
			desc: "Format",
			i: CreateMediumOptions{
				Format: VDI,
			},
			out: []string{
				"--format=vdi",
			},
		},
		{
			desc: "Diffparent",
			i: CreateMediumOptions{
				Diffparrent: "test/test.vdi",
			},
			out: []string{
				"--diffparent=test/test.vdi",
			},
		},
		{
			desc: "all options",
			i: CreateMediumOptions{
				Type:        DVDMedium,
				Filename:    "test/test.vdi",
				Size:        128,
				Format:      VDI,
				Diffparrent: "test/test.vdi",
			},
			out: []string{
				"dvd",
				"--filename=test/test.vdi",
				"--size=128",
				"--diffparent=test/test.vdi",
				"--format=vdi",
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

func TestModifyMediumOptions_slice(t *testing.T) {
	var tests = []struct {
		desc string
		i    ModifyMediumOptions
		out  []string
	}{
		{
			desc: "no options",
		},
		{
			desc: "Type",
			i: ModifyMediumOptions{
				Type: DiskTypeImmutable,
			},
			out: []string{
				"--type=immutable",
				"--autoreset=off",
			},
		},
		{
			desc: "Type w autoreset",
			i: ModifyMediumOptions{
				Type:      DiskTypeImmutable,
				AutoReset: true,
			},
			out: []string{
				"--type=immutable",
			},
		},
		{
			desc: "Size",
			i: ModifyMediumOptions{
				Size: 128,
			},
			out: []string{
				"--size=128",
			},
		},
		{
			desc: "Compact",
			i: ModifyMediumOptions{
				Compact: true,
			},
			out: []string{
				"--compact",
			},
		},
		{
			desc: "Move",
			i: ModifyMediumOptions{
				NewPath: "newfolder/test.vdi",
			},
			out: []string{
				"--move=newfolder/test.vdi",
			},
		},
		{
			desc: "Location",
			i: ModifyMediumOptions{
				Location: "newlocation/test.vdi",
			},
			out: []string{
				"--setlocation=newlocation/test.vdi",
			},
		},
		{
			desc: "all options",
			i: ModifyMediumOptions{
				Type: DiskTypeNormal,
				Size: 128,
			},
			out: []string{
				"--type=normal",
				"--size=128",
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

func TestCloneMediumOptions_slice(t *testing.T) {
	var tests = []struct {
		desc string
		i    CloneMediumOptions
		out  []string
	}{
		{
			desc: "no options",
		},
		{
			desc: "Format",
			i: CloneMediumOptions{
				Format: VDI,
			},
			out: []string{
				"--format=vdi",
			},
		},
		{
			desc: "Existing",
			i: CloneMediumOptions{
				Existing: true,
			},
			out: []string{
				"--existing",
			},
		},

		{
			desc: "all options",
			i: CloneMediumOptions{
				Format:   VDI,
				Existing: true,
			},
			out: []string{
				"--format=vdi",
				"--existing",
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
