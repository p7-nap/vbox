package vbox

import "fmt"

type MediumService struct {
	Create *MediumCreate
	Modify *MediumModify
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
	Type   mediumType
	Name   string
	Size   uint
	Format mediumFormat
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
	if c.Format != "" {
		s = append(s, fmt.Sprintf("--format=%s", c.Type))
	}
	return s
}

type MediumModify struct {
	v *Client
}
