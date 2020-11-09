package vbox

import "fmt"

type CreateService struct {
	v *Client
}

func (vc *CreateService) exec(args ...string) ([]byte, error) {
	return vc.v.exec("vboxmanage", args...)
}

func (vc *CreateService) CreateVM(options CreateOptions) error {
	args := []string{"createvm"}
	args = append(args, options.slice()...)
	_, err := vc.exec(args...)
	return err
}

type CreateOptions struct {
	Name            string
	Basefolder      string
	OSType          ostype
	Register        bool
	DefaultHardware bool
}

func (c CreateOptions) slice() []string {
	var s []string
	if c.Name != "" {
		s = append(s, fmt.Sprintf("--name=%s", c.Name))
	}
	if c.Basefolder != "" {
		s = append(s, fmt.Sprintf("--basefolder=%s", c.Basefolder))
	}
	if c.OSType != "" {
		s = append(s, fmt.Sprintf("--ostype=%s", c.OSType))
	}
	if c.Register {
		s = append(s, "--register")
	}
	if c.DefaultHardware {
		s = append(s, "--default")
	}
	return s
}

type ostype string

const (
	Other     ostype = "other"
	WindowsXP ostype = "windowsXP"
	Debian    ostype = "debian"
	Linux     ostype = "linux"
	Linux64   ostype = "linux_64"
)
