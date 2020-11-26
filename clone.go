package vbox

import (
	"fmt"
)

type CloneService struct {
	v *Client
}

func (v *CloneService) exec(args ...string) ([]byte, error) {
	return v.v.exec("vboxmanage", args...)
}

//CloneVM clones <vmname> with given <options>,...
func (v *CloneService) CloneVM(vmname string, options CloneOptions) error {
	args := []string{"clonevm", vmname}
	cloneOptions, err := options.slice()
	args = append(args, cloneOptions...)
	_, err = v.exec(args...)
	return err
}

type CloneOptions struct {
	Basefolder string
	Groups     string
	Mode       cloneMode
	Name       string
	//add options as []cloneOptions{cloneoption1, cloneoption2,..."
	Options  []CloneOption
	Register bool
	Snapshot string
	UUID     string
}

func (c CloneOptions) slice() ([]string, error) {
	var s []string
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
	if c.Options != nil {
		s = append(s, fmt.Sprintf("--options=%s", join(c.Options, ", ")))
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

type cloneMode string

const (
	Machine            cloneMode = "machine"
	Machineandchildren cloneMode = "machineandchildren"
	All                cloneMode = "all"
)

type CloneOption string

const (
	Link          CloneOption = "Link"
	KeepAllMACs   CloneOption = "KeepAllMACs"
	KeepNATMACs   CloneOption = "KeepNATMACs"
	KeepDiskNames CloneOption = "KeepDiskNames"
	KeepHwUUIDs   CloneOption = "KeepDiskNames"
)

func join(options []CloneOption, separator string) string {
	optionstring := string(options[0])
	for _, option := range options[1:] {
		optionstring = fmt.Sprintf("%s%s%s", optionstring, separator, option)
	}
	return optionstring
}
