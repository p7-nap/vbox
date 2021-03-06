package vbox

import "fmt"

type ModifyService struct {
	v *Client
}

func (vc *ModifyService) exec(args ...string) ([]byte, error) {
	return vc.v.exec("vboxmanage", args...)
}

func (vc *ModifyService) ModifyVM(name string, options ModifyOptions) error {
	args := []string{"modifyvm", name}
	args = append(args, options.slice()...)
	_, err := vc.exec(args...)
	return err
}

type ModifyOptions struct {
	// Name sets the new name of the VM
	Name string
	// RAM sets the amount of RAM in MB
	RAM uint
	// CPUs sets the number of virtual CPUs
	CPUs uint
	// VRAM sets the amount of RAM for the grapicscard
	VRAM   uint
	OSType ostype
	Nics   []Nic
	bootOptions
}

type bootOptions struct {
	Boot1 bootType
	Boot2 bootType
	Boot3 bootType
	Boot4 bootType
}

func (c ModifyOptions) slice() []string {
	var s []string
	if c.Name != "" {
		s = append(s, fmt.Sprintf("--name=%s", c.Name))
	}
	if c.RAM != 0 {
		s = append(s, fmt.Sprintf("--memory=%d", c.RAM))
	}
	if c.OSType != "" {
		s = append(s, fmt.Sprintf("--ostype=%s", c.OSType))
	}
	if c.VRAM != 0 {
		s = append(s, fmt.Sprintf("--vram=%d", c.VRAM))
	}
	if c.CPUs != 0 {
		s = append(s, fmt.Sprintf("--cpus=%d", c.CPUs))
	}
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

type bootType string

const (
	None   bootType = "none"
	Floppy bootType = "floppy"
	DVD    bootType = "dvd"
	Disk   bootType = "disk"
	Net    bootType = "net"
)

type Nic struct {
	Mode         nicMode
	Iface        string
	Promisc      promiscMode
	PortForward  bool
	ForwardRules []ForwardRule
}
type nicMode string

const (
	NONE    nicMode = "none"
	NULL    nicMode = "null"
	NAT     nicMode = "nat"
	Bridged nicMode = "bridged"
)

type promiscMode string

const (
	DENY     promiscMode = "deny"
	AllowVM  promiscMode = "allow-vms"
	AllowAll promiscMode = "allow-all"
)

type ForwardRule struct {
	Name      string
	Protocol  protocol
	HostIP    string
	HostPort  uint
	GuestIP   string
	GuestPort uint
}

type protocol string

const (
	TCP protocol = "tcp"
	UDP protocol = "udp"
)
