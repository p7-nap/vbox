package vbox

import "fmt"

type ControlService struct {
	v *Client
}

func (v *ControlService) exec(args ...string) ([]byte, error) {
	return v.v.exec("vboxmanage", args...)
}

func (v *ControlService) Pause(vmname string) error {
	args := []string{"controlvm", vmname, "pause"}
	_, err := v.exec(args...)
	return err
}

func (v *ControlService) Resume(vmname string) error {
	args := []string{"controlvm", vmname, "resume"}
	_, err := v.exec(args...)
	return err
}

func (v *ControlService) Reset(vmname string) error {
	args := []string{"controlvm", vmname, "reset"}
	_, err := v.exec(args...)
	return err
}

func (v *ControlService) PowerOFF(vmname string) error {
	args := []string{"controlvm", vmname, "poweroff"}
	_, err := v.exec(args...)
	return err
}

func (v *ControlService) SaveState(vmname string) error {
	args := []string{"controlvm", vmname, "savestate"}
	_, err := v.exec(args...)
	return err
}

func (v *ControlService) AcpiPowerButton(vmname string) error {
	args := []string{"controlvm", vmname, "acpipowerbutton"}
	_, err := v.exec(args...)
	return err
}

func (v *ControlService) ControlVM(vmname string, options ControlOptions) error {
	args := []string{"controlvm", vmname}
	cloneOptions := options.slice()
	args = append(args, cloneOptions...)
	_, err := v.exec(args...)
	return err
}

type ControlOptions struct {
	Nics []Nic
}

func (c ControlOptions) slice() []string {
	var s []string
	for i, n := range c.Nics {
		if n.Mode != "" {
			s = append(s, fmt.Sprintf("nic%d", i+1), fmt.Sprintf("%s", n.Mode))
		}
		if n.Mode == Bridged && n.Iface != "" {
			s = append(s, fmt.Sprintf("%s", n.Iface))
		}
	}

	return s
}
