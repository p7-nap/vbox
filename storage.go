package vbox

import (
	"errors"
	"fmt"
)

var (
	//ErrNoName given if no storage controller name has been given
	ErrNoName = errors.New("Storage name must be diffined")
)

//StorageService struct with entry Control and Attach of types StorageControl and StorageAttach
type StorageService struct {
	Control *StorageControl
	Attach  *StorageAttach
}

//StorageControl struct with entry c of type Client
type StorageControl struct {
	c *Client
}

//StorageAttach struct with entry c of type Client
type StorageAttach struct {
	c *Client
}

func (vc *StorageControl) exec(args ...string) ([]byte, error) {
	return vc.c.exec("vboxmanage", args...)
}

//ControlStorage takes vmname/uuid and storageoptions and creates a storage controller with given options
func (vc *StorageControl) ControlStorage(name string, options StoragectlOptions) error {
	args := []string{"storagectl", name}
	controlOptions, err := options.slice()
	args = append(args, controlOptions...)
	_, err = vc.exec(args...)
	return err
}

//StoragectlOptions contains the different options that can be used with vboxmanage storagectl
type StoragectlOptions struct {
	//Name is usualy busType + Controller
	Name       string
	Add        busType
	Controller chipsetType
	//Portcount should be int between 1 and 30
	Portcount   uint
	HostIOCache bool
	Bootable    bool
	Rename      string
	Remove      bool
}

func (c StoragectlOptions) slice() ([]string, error) {

	var s []string
	if c.Name != "" {
		s = append(s, fmt.Sprintf("--name=%s", c.Name))
	}
	if c.Name == "" {
		return nil, ErrNoName
	}
	if c.Add != "" {
		s = append(s, fmt.Sprintf("--add=%s", c.Add))
	}
	if c.Controller != "" {
		s = append(s, fmt.Sprintf("--controller=%s", c.Controller))
	}
	if c.Portcount >= 1 && c.Portcount <= 30 {
		s = append(s, fmt.Sprintf("--portcount=%d", c.Portcount))
	}
	if c.HostIOCache {
		s = append(s, "--hostiocache=on")
	}
	if c.Bootable {
		s = append(s, "--bootable=on")
	}
	if c.Rename != "" {
		s = append(s, fmt.Sprintf("--rename=%s", c.Rename))
	}
	if c.Remove {
		s = append(s, "--remove")
	}
	return s, nil
}

type busType string

const (
	//IDE is a type of bus used in command add
	IDE busType = "ide"
	//SATA is a type of bus used in command add
	SATA busType = "sata"
	//SCSI is a type of bus used in command add
	SCSI busType = "scsi"
	//FloppyStorage is a type of bus used in command add
	FloppyStorage busType = "floppy"
	//SAS is a type of bus used in command add
	SAS busType = "sas"
	//USBBus is a type of bus used in command add
	USBBus busType = "usb"
	//PCIE is a type of bus used in command add
	PCIE busType = "pcie"
)

type chipsetType string

const (
	LSILogic    chipsetType = "LSILogic"
	LSILogicSAS chipsetType = "LSILogicSAS"
	BusLogic    chipsetType = "BusLogic"
	IntelAhci   chipsetType = "IntelAhci"
	PIIX3       chipsetType = "PIIX3"
	PIIX4       chipsetType = "PIIX4"
	IHC6        chipsetType = "IHC6"
	I82078      chipsetType = "I82078"
	USBChipset  chipsetType = "USB"
	NVMe        chipsetType = "NVMe"
	VirtIO      chipsetType = "VirtIO"
)

func (vc *StorageAttach) exec(args ...string) ([]byte, error) {
	return vc.c.exec("vboxmanage", args...)
}

//AttachStorage takes vmname/uuid and storageattachoptions and attaches a storage controller with given options
func (vc *StorageAttach) AttachStorage(name string, options StorageattachOptions) error {
	args := []string{"storageattach", name}
	args = append(args, options.slice()...)
	_, err := vc.exec(args...)
	return err
}

type StorageattachOptions struct {
	SctlName   string
	Port       uint
	Device     uint
	Type       driveType
	MediumPath string
}

func (c StorageattachOptions) slice() []string {

	var s []string
	if c.SctlName != "" {
		s = append(s, fmt.Sprintf("--storagectl=%s", c.SctlName))
	}
	if c.Port >= 0 {
		s = append(s, fmt.Sprintf("--port=%d", c.Port))
	}
	if c.Device >= 0 {
		s = append(s, fmt.Sprintf("--device=%d", c.Device))
	}
	if c.Type != "" {
		s = append(s, fmt.Sprintf("--type=%s", c.Type))
	}
	if c.MediumPath != "" {
		s = append(s, fmt.Sprintf("--medium=%s", c.MediumPath))
	}

	return s
}

type driveType string

const (
	DVDDrive driveType = "dvddrive"
	HDD      driveType = "hdd"
	FDD      driveType = "fdd"
)
