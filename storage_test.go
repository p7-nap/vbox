package vbox

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStoragectlOptions_slice(t *testing.T) {
	var tests = []struct {
		desc string
		i    StoragectlOptions
		out  []string
		err  error
	}{
		{
			desc: "All options",
			i: StoragectlOptions{
				Name:        "SATA controller",
				Add:         SATA,
				Controller:  PIIX4,
				Portcount:   1,
				HostIOCache: true,
				Bootable:    true,
				Rename:      "newname",
				Remove:      true,
			},
			out: []string{
				"--name=SATA controller",
				"--add=sata",
				"--controller=PIIX4",
				"--portcount=1",
				"--hostiocache=on",
				"--bootable=on",
				"--rename=newname",
				"--remove"},
			err: nil,
		},
		{
			desc: "All options but remove",
			i: StoragectlOptions{
				Name:        "SATA controller",
				Add:         SATA,
				Controller:  PIIX4,
				Portcount:   1,
				HostIOCache: true,
				Bootable:    true,
				Rename:      "newname",
				Remove:      false,
			},
			out: []string{
				"--name=SATA controller",
				"--add=sata",
				"--controller=PIIX4",
				"--portcount=1",
				"--hostiocache=on",
				"--bootable=on",
				"--rename=newname"},
			err: nil,
		},
		{
			desc: "All options but No name",
			i: StoragectlOptions{
				Name:        "",
				Add:         SATA,
				Controller:  PIIX4,
				Portcount:   1,
				HostIOCache: true,
				Bootable:    true,
				Rename:      "newname",
				Remove:      true,
			},
			err: fmt.Errorf("Storage name must be diffined"),
		},
		{
			desc: "all options but remove",
			i: StoragectlOptions{
				Name:        "SATA controller",
				Add:         SATA,
				Controller:  PIIX4,
				Portcount:   1,
				HostIOCache: true,
				Bootable:    true,
				Rename:      "newname",
				Remove:      false,
			},
			out: []string{
				"--name=SATA controller",
				"--add=sata",
				"--controller=PIIX4",
				"--portcount=1",
				"--hostiocache=on",
				"--bootable=on",
				"--rename=newname"},
			err: nil,
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

func TestStorageattachOptions_slice(t *testing.T) {
	var tests = []struct {
		desc string
		i    StorageattachOptions
		out  []string
	}{
		{
			desc: "All options hdd",
			i: StorageattachOptions{
				SctlName:   "SATA Controller",
				Port:       1,
				Device:     1,
				Type:       HDD,
				MediumPath: "/path/file.vdi",
			},
			out: []string{
				"--storagectl=SATA Controller",
				"--port=1",
				"--device=1",
				"--type=hdd",
				"--medium=/path/file.vdi",
			},
		},
		{
			desc: "All options dvddrive",
			i: StorageattachOptions{
				SctlName:   "SATA Controller",
				Port:       1,
				Device:     1,
				Type:       DVDDrive,
				MediumPath: "/path/file.vdi",
			},
			out: []string{
				"--storagectl=SATA Controller",
				"--port=1",
				"--device=1",
				"--type=dvddrive",
				"--medium=/path/file.vdi",
			},
		},
		{
			desc: "All options fdd",
			i: StorageattachOptions{
				SctlName:   "SATA Controller",
				Port:       1,
				Device:     1,
				Type:       FDD,
				MediumPath: "/path/file.vdi",
			},
			out: []string{
				"--storagectl=SATA Controller",
				"--port=1",
				"--device=1",
				"--type=fdd",
				"--medium=/path/file.vdi",
			},
		},
		{
			desc: "port and device not set",
			i: StorageattachOptions{
				SctlName:   "SATA Controller",
				Type:       FDD,
				MediumPath: "/path/file.vdi",
			},
			out: []string{
				"--storagectl=SATA Controller",
				"--port=0",
				"--device=0",
				"--type=fdd",
				"--medium=/path/file.vdi",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			want := tt.out
			got := tt.i.slice()
			if !assert.Equal(t, want, got) {
				t.Fatalf("unexpected slices:\n- want: %v\n-  got: %v",
					want, got)
			}
		})
	}
}
