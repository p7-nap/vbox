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

// ChangeNIC changes the the NIC with the given number, however only one thing can be changed at a time
func (v *ControlService) ChangeNIC(vmname string, nicNumber uint, nic Nic) error {
	args := []string{"controlvm", vmname}
	options := nic.slice(nicNumber)
	args = append(args, options...)
	_, err := v.exec(args...)
	return err

}

func (n Nic) slice(nicNumber uint) []string {
	var s []string
	if n.PortForward == true {
		for _, r := range n.ForwardRules {
			s = append(s, fmt.Sprintf("--natpf%d=%s,%s,%s,%d,%s,%d", nicNumber, r.Name, r.Protocol, r.HostIP, r.HostPort, r.GuestIP, r.GuestPort))
		}
		return s
	}
	if n.Mode != "" {
		s = append(s, fmt.Sprintf("nic%d", nicNumber), fmt.Sprintf("%s", n.Mode))
	}
	if n.Mode == Bridged && n.Iface != "" {
		s = append(s, fmt.Sprintf("%s", n.Iface))
	}
	return s
}
