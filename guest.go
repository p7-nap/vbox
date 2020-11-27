package vbox

import "fmt"

//GuestService provides the usefull guestcontrol commands.
type GuestService struct {
	v *Client
}

func (vc *GuestService) exec(args ...string) ([]byte, error) {
	return vc.v.exec("vboxmanage", args...)
}

//GuestRun uses command "guestcontrol run" from vboxmanage with RunOptions
func (vc *GuestService) GuestRun(vmname string, options RunOptions) error {
	args := []string{"guestcontrol", vmname, "run"}
	args = append(args, options.slice()...)
	_, err := vc.exec(args...)
	return err
}

//GuestCopyTo utilises the command "guestcontrol copyto" with the neccessary options
func (vc *GuestService) GuestCopyTo(vmname string, options CopyOptions) error {
	args := []string{"guestcontrol", vmname, "copyto"}
	args = append(args, options.slice()...)
	_, err := vc.exec(args...)
	return err
}

//GuestCopyFrom utilises the command "guestcontrol copyfrom" with the neccessary options
func (vc *GuestService) GuestCopyFrom(vmname string, options CopyOptions) error {
	args := []string{"guestcontrol", vmname, "copyfrom"}
	args = append(args, options.slice()...)
	_, err := vc.exec(args...)
	return err
}

//RunOptions must have the options for the command you want to run on the virtual machine.
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
		s = append(s, fmt.Sprintf("--exe=%s", o.ExecutePath))
	}

	if o.Username != "" {
		s = append(s, fmt.Sprintf("--username=%s", o.Username))
	}
	if o.Password != "" {
		s = append(s, fmt.Sprintf("--password=%s", o.Password))
	} else if o.PasswordFile != "" {
		s = append(s, fmt.Sprintf("--passwordfile=%s", o.PasswordFile))
	}

	if o.Timeout != 0 {
		s = append(s, fmt.Sprintf("--timeout=%d", o.Timeout))
	}
	if o.STDOut {
		s = append(s, "--wait-stdout")
	}
	if o.STDErr {
		s = append(s, "--wait-stderr")
	}
	return s
}

//CopyOptions must contain the options you want to include in the command to dopy things to or from the virtual machine.
type CopyOptions struct {
	//TargetPath specifies the path to the target
	TargetPath string
	//SourcePath specifies the path of the source
	SourcePath string
	//Username is the user on the vm.
	Username string
	//Password is the password of the user on the vm.
	Password string
	//PasswordFile takes the relative path to a file containing the password.
	PasswordFile string
}

func (o CopyOptions) slice() []string {
	var s []string
	if o.TargetPath != "" {
		s = append(s, fmt.Sprintf("--target-directory %s", o.TargetPath))
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
