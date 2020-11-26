package vbox

import "fmt"

type GuestService struct {
	v *Client
}

func (vc *GuestService) exec(args ...string) ([]byte, error) {
	return vc.v.exec("vboxmanage", args...)
}

func (vc *GuestService) GuestRun(vmname string, options RunOptions) error {
	args := []string{"guestcontrol", vmname, "run"}
	args = append(args, options.slice()...)
	_, err := vc.exec(args...)
	return err
}

func (vc *GuestService) GuestCopyTo(vmname string, options CopyOptions) error {
	args := []string{"guestcontrol", vmname, "copyto"}
	args = append(args, options.slice()...)
	_, err := vc.exec(args...)
	return err
}

func (vc *GuestService) GuestCopyFrom(vmname string, options CopyOptions) error {
	args := []string{"guestcontrol", vmname, "copyfrom"}
	args = append(args, options.slice()...)
	_, err := vc.exec(args...)
	return err
}

type RunOptions struct {
	//Username is the user on the vm.
	Username string
	//Password is the password of the user on the vm.
	Password string
	//PasswordFile takes the relative path to a file containing the password.
	PasswordFile string
	//ExecutePath takes the absolute path to the executable file which you wanna run on the guest.
	//Example: /bin/ls
	ExecutePath string
	// further options are optional
	//Timeout takes the maximum time in microseconds that the executable can run, if non specified it runs until done.
	Timeout int64
	//STDOutstdout when true the function will wait for output from guest and recieve exitcode and reason/flags
	STDOut bool
	//STDErr when true waits until process ends and recieves exitcode errors and flags.
	STDErr bool
}

func (o RunOptions) slice() []string {
	var s []string
	if o.ExecutePath != "" {
		s = append(s, fmt.Sprintf("--exe %s", o.ExecutePath))
	}

	if o.Username != "" {
		s = append(s, fmt.Sprintf("--username %s", o.Username))
	}
	if o.Password != "" {
		s = append(s, fmt.Sprintf("--password %s", o.Password))
	} else if o.PasswordFile != "" {
		s = append(s, fmt.Sprintf("--passwordfile %s", o.PasswordFile))
	}

	if o.Timeout != 0 {
		s = append(s, fmt.Sprintf("--timeout %s", o.Timeout))
	}
	if o.STDOut {
		s = append(s, "--wait-stdout")
	}
	if o.STDErr {
		s = append(s, "--wait-stderr")
	}
	return s
}

type CopyOptions struct {
	//TargetPath specifies the path to the target
	TargetPath string
	SourcePath string
	//Username is the user on the vm.
	Username string
	//Password is the password of the user on the vm.
	Password string
	//PasswordFile takes the relative path to a file containing the password.
	PasswordFile string
	//TargetPath

}

func (o CopyOptions) slice() []string {
	var s []string
	if o.TargetPath != "" {
		s = append(s, fmt.Sprintf("--target-directory %s",o.TargetPath)
	}
	if o.SourcePath != "" {
		s = append(s, o.SourcePath)
	}
	if o.Username != "" {
		s = append(s, fmt.Sprintf("--username %s", o.Username))
	}
	if o.Password != "" {
		s = append(s, fmt.Sprintf("--password %s", o.Password))
	} else if o.PasswordFile != "" {
		s = append(s, fmt.Sprintf("--passwordfile %s", o.PasswordFile))
	}

	return s
}
