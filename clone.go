package vbox

import "fmt"

type CloneService struct {
	v *Client
}

func (v *CloneService) exec(args ...string) ([]byte, error) {
	return v.v.exec("vboxmanage", args...)
}

func (v *CloneService) CloneVM(options CloneOptions) error {
	args := []string{"clonevm"}
	options, err := options.slice()
	args = append(args, options...)
	_, err = v.exec(args...)
	return err
}

type CloneOptions struct {
	Vmname     string
	Basefolder string
	Groups     string
	Mode       mode
	Name       string
	Options    options
	Register   bool
	Snapshot   string
	UUID       string
}

func (c CloneOptions) slice() ([]string, error) {
	var s []string
	if c.Vmname != "" {
		s = append(s, c.Name)
	}
	if c.Basefolder != "" {
		s = append(s, fmt.Sprintf("--basefolder=%s", c.Basefolder))
	}
	if c.Groups != "" {
		s = append(s, fmt.Sprintf("--groups=%s", c.Groups))
	}
	if c.Mode != "" {
		if c.Snapshot == "" && c.Mode == "machineandchildren" {
			return s, fmt.Errorf("%v mode parameter only available with --snapshot", c.Mode)
		}
		s = append(s, fmt.Sprintf("--mode=%s", c.Mode))
	}
	if c.Name != "" {
		s = append(s, fmt.Sprintf("--name=%s", c.Name))
	}
	if c.Options != "" {
		s = append(s, fmt.Sprintf("--options=%s", c.Options))
	}
	if c.Register {
		s = append(s, "--register")
	}
	if c.Snapshot != "" {
		s = append(s, fmt.Sprintf("--snapshot=%s", c.Snapshot))
	}
	if c.UUID != "" {
		s = append(s, fmt.Sprintf("--uuid=%s", c.UUID))
	}
	return s, nil
}

type mode string

const (
	Machine            mode = "machine"
	Machineandchildren mode = "machineandchildren"
	All                mode = "all"
)

type options string

const (
	Link          options = "Link"
	KeepAllMACs   options = "KeepAllMACs"
	KeepNATMACs   options = "KeepNATMACs"
	KeepDiskNames options = "KeepDiskNames"
	KeepHwUUIDs   options = "KeepDiskNames"
)
