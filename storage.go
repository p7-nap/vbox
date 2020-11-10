package vbox

import(
	""
)

var (
	ErrNoName = errors.New("Storage name must be diffined.")
)
type StorageSevice struct{
	Control *StorageControl
	Attach  *StorageAttach
}

type StorageControl struct {
	c *Client
}

type StorageAttach struct {
	c *Client
}

func (vc *StorageControl) exec(args ...string) ([]byte, error) {
	return vc.c.exec("vboxmanage", args...)
}
func (vc *StorageAttach) exec(args ...string) ([]byte, error) {
	return vc.c.exec("vboxmanage", args...)
}

func (vc *ModifyService) (name string, options StorageOptions) error {
	args := []string{"storagectl", name}
	args = append(args, options.slice()...)
	_, err := vc.exec(args...)
	return err
}

type StorageOptions struct{
	Name string
	Add busType
	Controller chipsetType
	//Portcount should be int between 1 and 30
	Portcount uint
	HostIOCache bool
	Bootable bool
	Rename string
	Remove bool


}
func (c StorageOptions) slice() ([]string,error){

	var s []string
	if c.Name != "" {
		s = append(s, fmt.Sprintf("--name=%s",c.Name))
	}
	if c.Name == ""{
		return nil, ErrNoName
	}
	if c.Add != "" {
		s = append(s, fmt.Sprintf("--add=%s",c.Add))
	}
	if c.Controller != "" {
		s = append(s, fmt.Sprintf("--controller=%s",c.Controller))
	}
	if 1 <= c.Portcount <= 30 {
		s = append(s, fmt.Sprintf("--portcount=%d",c.Portcount))
	}
	if c.HostIOCache {
		s = append(s,"--hostiocache=on")
	}
	if c.Bootable {
		s = append(s,"--bootable=on")
	}
	if c.Rename != "" {
		s = append(s, fmt.Sprintf("--rename=%s",c.Rename))
	}
	if c.Remove {
		s =  append(s, "--remove")
	}
	return s, nil
}

type busType string

const (
	IDE busType = "ide"
	SATA busType = "sata"
	SCSI busType = "scsi"
	FloppyStorage busType = "floppy"
	SAS busType = "sas"
	USB busType = "usb"
	PCIE busType = "pcie"
)

type chipsetType string

const (
	LSILogic chipsetType = "LSILogic"
	LSILogicSAS chipsetType = "LSILogicSAS"
	BusLogic chipsetType = "BusLogic"
	IntelAhci chipsetType = "IntelAhci"
	PIIX3 chipsetType = "PIIX3"
	PIIX4 chipsetType = "PIIX4"
	IHC6 chipsetType = "IHC6"
	I82078 chipsetType = "I82078"
	USB chipsetType = "USB"
	NVMe chipsetType = "NVMe"
	VirtIO chipsetType = "VirtIO"
)