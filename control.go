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
			s = append(s, fmt.Sprintf("--nic%d=%s", i+1, n.Mode))
		}
		if n.Mode == Bridged && n.Iface != "" {
			s = append(s, fmt.Sprintf("--bridgeadapter%d=%s", i+1, n.Iface))
		}
		if n.Mode == Bridged && n.Promisc != "" {
			s = append(s, fmt.Sprintf("--nicpromisc%d=%s", i+1, n.Promisc))
		}
		if n.Mode == NAT && n.PortForward == true {
			for _, r := range n.ForwardRules {
				s = append(s, fmt.Sprintf("--natpf%d=%s,%s,%s,%d,%s,%d", i+1, r.Name, r.Protocol, r.HostIP, r.HostPort, r.GuestIP, r.GuestPort))
			}
		}
	}

	return s
}
