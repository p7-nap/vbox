package vbox

import "fmt"

type StartService struct {
	v *Client
}

func (vc *StartService) exec(args ...string) ([]byte, error) {
	return vc.v.exec("vboxmanage", args...)
}

func (vc *StartService) StartVM(vmname string, options StartOptions) error {
	args := []string{"startvm", vmname}
	args = append(args, options.slice()...)
	_, err := vc.exec(args...)
	return err
}

type StartOptions struct {
	Type startType
}

func (o StartOptions) slice() []string {
	var s []string
	if o.Type != "" {
		s = append(s, fmt.Sprintf("--type=%s", o.Type))
	}

	return s
}

type startType string

const (
	GUI      startType = "gui"
	SDL      startType = "sdl"
	Headless startType = "headless"
	Seperate startType = "separate"
)
