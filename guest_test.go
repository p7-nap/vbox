package vbox

import (
	"reflect"
	"testing"
)

func TestRunOptions_slice(t *testing.T) {
	var tests = []struct {
		desc string
		i    RunOptions
		out  []string
	}{
		{
			desc: "All options",
			i: RunOptions{
				Username:     "user",
				Password:     "UserPass",
				PasswordFile: "/home/userpassfile.txt",
				ExecutePath:  "/bin/ls",
				Timeout:      1000,
				STDOut:       true,
				STDErr:       true,
			},
			out: []string{
				"--exe=/bin/ls",
				"--username=user",
				"--password=UserPass",
				"--timeout=1000",
				"--wait-stdout",
				"--wait-stderr"},
		},
		{
			desc: "No options",
		},
		{
			desc: "No timeout no wait, use passwordfile",
			i: RunOptions{
				Username:     "user",
				PasswordFile: "/home/userpassfile.txt",
				ExecutePath:  "/bin/ls",
			},
			out: []string{
				"--exe=/bin/ls",
				"--username=user",
				"--passwordfile=/home/userpassfile.txt"},
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

func TestCopyOptions_slice(t *testing.T) {
	var tests = []struct {
		desc string
		i    CopyOptions
		out  []string
	}{
		{
			desc: "All options",
			i: CopyOptions{
				Username:     "user",
				Password:     "UserPass",
				PasswordFile: "/home/userpassfile.txt",
				TargetPath:   "/target/location",
				SourcePath:   "/source/file/location/file.txt",
			},
			out: []string{
				"--target-directory /target/location",
				"/source/file/location/file.txt",
				"--username user",
				"--password UserPass"},
		},
		{
			desc: "No options",
		},
		{
			desc: "Using passwordfile",
			i: CopyOptions{
				Username:     "user",
				PasswordFile: "/home/userpassfile.txt",
				TargetPath:   "/target/location",
				SourcePath:   "/source/file/location/file.txt",
			},
			out: []string{
				"--target-directory /target/location",
				"/source/file/location/file.txt",
				"--username user",
				"--passwordfile /home/userpassfile.txt"},
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
