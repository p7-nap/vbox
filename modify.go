package vbox

import "fmt"

type ModifyService struct {
	v *Client
}

func (vc *ModifyService) exec(args ...string) ([]byte, error) {
	return vc.v.exec("vboxmanage", args...)
}

func (vc *ModifyService) ModigfyVM(name string, options ModifyOptions) error {
	args := []string{"modifyvm"}
	args = append(args, options.slice()...)
	_, err := vc.exec(args...)
	return err
}

type ModifyOptions struct {
	// Name sets the new name of the VM
	Name string
	// RAM sets the amount of RAM in MB
	RAM uint
	// CPUs sets the number of virtual CPUs
	CPUs uint
	// VRAM sets the amount of RAM for the grapicscard
	VRAM   uint
	OSType ostype
	bootOptions
    networkOptions
}

type bootOptions struct {
	Boot1 bootType
	Boot2 bootType
	Boot3 bootType
	Boot4 bootType
}

func (c ModifyOptions) slice() []string {
	var s []string
	if c.Name != "" {
		s = append(s, fmt.Sprintf("--name=%s", c.Name))
	}
	if c.RAM != 0 {
		s = append(s, fmt.Sprintf("--memory=%d", c.RAM))
	}
	if c.OSType != "" {
		s = append(s, fmt.Sprintf("--ostype=%s", c.OSType))
	}
	if c.VRAM != 0 {
		s = append(s, fmt.Sprintf("--vram=%d", c.VRAM))
	}
	if c.CPUs != 0 {
		s = append(s, fmt.Sprintf("--cpus=%d", c.CPUs))
	}

	return s
}

type bootType string

const (
	None   bootType = "none"
	Floppy bootType = "floppy"
	DVD    bootType = "dvd"
	Disk   bootType = "disk"
	Net    bootType = "net"
)

type networkOptions struct{
    NICs []
}
