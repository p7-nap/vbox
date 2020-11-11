package vbox

import "fmt"

type MediumService struct {
	Create *MediumCreate
	Modify *MediumModify
	Clone  *MediumClone
	v      *Client
}

func (vc *MediumCreate) exec(args ...string) ([]byte, error) {
	return vc.v.exec("vboxmanage", args...)
}

type MediumCreate struct {
	v *Client
}

func (m *MediumCreate) CreateMedium(options CreateMediumOptions) error {
	args := []string{"createmedium"}
	args = append(args, options.slice()...)
	_, err := m.exec(args...)
	return err
}

type CreateMediumOptions struct {
	Type        mediumType
	Filename    string
	Size        uint
	Format      mediumFormat
	Diffparrent string
}

type mediumFormat string

const (
	VDI  mediumFormat = "vdi"
	VMDK mediumFormat = "vmdk"
	VHD  mediumFormat = "vhd"
)

type mediumType string

const (
	DVDMedium    mediumType = "dvd"
	DiskMedium   mediumType = "disk"
	FloppyMedium mediumType = "floppy"
)

func (c CreateMediumOptions) slice() []string {
	var s []string
	if c.Type != "" {
		s = append(s, fmt.Sprintf("%s", c.Type))
	}
	if c.Filename != "" {
		s = append(s, fmt.Sprintf("--filename=%s", c.Filename))
	}
	if c.Size != 0 {
		s = append(s, fmt.Sprintf("--size=%d", c.Size))
	}
	if c.Diffparrent != "" {
		s = append(s, fmt.Sprintf("--diffparent=%s", c.Diffparrent))
	}
	if c.Format != "" {
		s = append(s, fmt.Sprintf("--format=%s", c.Format))
	}
	return s
}

type MediumModify struct {
	v *Client
}

func (vc *MediumModify) exec(args ...string) ([]byte, error) {
	return vc.v.exec("vboxmanage", args...)
}

func (m *MediumModify) ModifyMedium(name string, options ModifyMediumOptions) error {
	args := []string{"modifymedium", name}
	args = append(args, options.slice()...)
	_, err := m.exec(args...)
	return err
}

type ModifyMediumOptions struct {
	Type      diskType
	AutoReset bool
	Size      uint
	Compact   bool
	NewPath   string
	Location  string
}

type diskType string

const (
	DiskTypeNormal    diskType = "normal"
	DiskTypeWriteT    diskType = "writethrough"
	DiskTypeImmutable diskType = "immutable"
	DiskTypeShareAble diskType = "shareable"
	DiskTypeReadonly  diskType = "readonly"
	DiskTypeMultiAtt  diskType = "multiattach"
)

func (c ModifyMediumOptions) slice() []string {
	var s []string
	if c.Type != "" {
		s = append(s, fmt.Sprintf("--type=%s", c.Type))
	}
	if c.Size != 0 {
		s = append(s, fmt.Sprintf("--size=%d", c.Size))
	}
	if c.Type == DiskTypeImmutable && !c.AutoReset {
		s = append(s, "--autoreset=off")
	}
	if c.Compact {
		s = append(s, "--compact")
	}
	if c.NewPath != "" {
		s = append(s, fmt.Sprintf("--move=%s", c.NewPath))
	}
	if c.Location != "" {
		s = append(s, fmt.Sprintf("--setlocation=%s", c.Location))
	}
	return s
}

type MediumClone struct {
	v *Client
}

func (vc *MediumClone) exec(args ...string) ([]byte, error) {
	return vc.v.exec("vboxmanage", args...)
}

func (m *MediumClone) CloneMedium(input, output string, options CloneMediumOptions) error {
	args := []string{"modifymedium", input, output}
	args = append(args, options.slice()...)
	_, err := m.exec(args...)
	return err
}

type CloneMediumOptions struct {
	Format   mediumFormat
	Existing bool
}

func (c CloneMediumOptions) slice() []string {
	var s []string
	if c.Format != "" {
		s = append(s, fmt.Sprintf("--format=%s", c.Format))
	}
	if c.Existing {
		s = append(s, "--existing")
	}
	return s
}
