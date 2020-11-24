package vbox

import (
	"reflect"
	"testing"
)

func TestImportOptions_slice(t *testing.T) {
	var tests = []struct {
		desc string
		i    ImportOptions
		out  []string
	}{
		{
			desc: "all options",
			i: ImportOptions{
				Option:          Keepallmacs,
				Name:            "vmname",
				Cloud:           true,
				CloudProfile:    "cloudprofile",
				CloudInstanceID: "cloudinstanceid",
				CloudBucket:     "cloudbucket",
				Vsys:            "0",
				OSType:          Linux,
				VMname:          "vmnameVsys",
				Group:           "/",
				SettingsFile:    "basefolder/Settingsfile",
				BaseFolder:      "basefolder",
				CPUs:            "1",
				Memory:          "1000",
				Disable:         []int{7, 8, 9, 10, 11},
			},
			out: []string{
				"--options=keepallmacs",
				"--vmname=vmname",
				"--cloud",
				"--cloudprofile=cloudprofile",
				"--cloudinstanceid=cloudinstanceid",
				"--cloudbucket=cloudbucket",
				"--vsys=0",
				"--ostype=linux",
				"--vmname=vmnameVsys",
				"--group=/",
				"--settingsfile=basefolder/Settingsfile",
				"--basefolder=basefolder",
				"--cpus=1",
				"--memory=1000",
				"--unit 7 --ignore",
				"--unit 8 --ignore",
				"--unit 9 --ignore",
				"--unit 10 --ignore",
				"--unit 11 --ignore",
			},
		},
		{
			desc: "machineandchildren mode with snapshot",
			i: ImportOptions{
				Option: Keepallmacs,
				OSType: Linux,
			},
			out: []string{
				"--options=keepallmacs",
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
