package vbox

import "fmt"

type ImportService struct {
	v *Client
}

func (is *ImportService) exec(args ...string) ([]byte, error) {
	return is.v.exec("vboxmanage", args...)
}

//ImportVM imports <ovfname> with given <options>,...
func (is *ImportService) ImportVM(ovfname string, options ImportOptions) error {
	args := []string{"import", ovfname}
	importOptions := options.slice()
	args = append(args, importOptions...)
	_, err := is.exec(args...)
	return err
}

//ImportOptions is passed to ImportVm desired options
type ImportOptions struct {
	Option          importOption
	Name            string
	Cloud           bool
	Cloudprofile    string
	Cloudinstanceid string
	Cloudbucket     string
	//additional options depend on ovf file
	//Display options with "VBoxManage import <filepath> -n"
	//These options all have a default and can be left out
	//if Vsys is not set vsys options are ignored
	Vsys   string
	OSType ostype
	//name set with vsys, enable vsys if using this
	Vmname string
	Group  string
	//path to .vbox file
	// default: <basefolder path> + <vmname> + <vmname>.vbox)
	Settingsfile string
	Basefolder   string
	//number of cpus
	CPUs string
	//memory in MBs
	Memory string
	// unit nrs that should be disabled
	Disable []int
}

func (o ImportOptions) slice() []string {
	var s []string
	if o.Option != "" {
		s = append(s, fmt.Sprintf("--options=%s", o.Option))
	}
	if o.Name != "" {
		s = append(s, fmt.Sprintf("--vmname=%s", o.Name))
	}
	if o.Cloud {
		s = append(s, fmt.Sprintf("--cloud"))
	}
	if o.Cloudprofile != "" {
		s = append(s, fmt.Sprintf("--cloudprofile=%s", o.Cloudprofile))
	}
	if o.Cloudinstanceid != "" {
		s = append(s, fmt.Sprintf("--cloudinstanceid=%s", o.Cloudinstanceid))
	}
	if o.Cloudbucket != "" {
		s = append(s, fmt.Sprintf("--cloudbucket=%s", o.Cloudbucket))
	}
	//more options / vsys options
	//if Vsys is not set vsys options are ignored
	if o.Vsys != "" {
		s = append(s, fmt.Sprintf("--vsys=%s", o.Vsys))
	} else {
		return s
	}
	if o.OSType != "" {
		s = append(s, fmt.Sprintf("--ostype=%s", o.OSType))
	}
	if o.Vmname != "" {
		s = append(s, fmt.Sprintf("--vmname=%s", o.Vmname))
	}
	if o.Group != "" {
		s = append(s, fmt.Sprintf("--group=%s", o.Group))
	}
	if o.Settingsfile != "" {
		s = append(s, fmt.Sprintf("--settingsfile=%s", o.Settingsfile))
	}
	if o.Basefolder != "" {
		s = append(s, fmt.Sprintf("--basefolder=%s", o.Basefolder))
	}
	if o.CPUs != "" {
		s = append(s, fmt.Sprintf("--cpus=%s", o.CPUs))
	}
	if o.Memory != "" {
		s = append(s, fmt.Sprintf("--memory=%s", o.Memory))
	}
	if o.Disable != nil {
		for _, unit := range o.Disable {
			s = append(s, fmt.Sprintf("--unit %d --ignore", unit))
		}
	}

	return s
}

type importOption string

//put these options in a general const somewhere?
const (
	Keepallmacs importOption = "keepallmacs"
	Keepnatmacs importOption = "keepnatmacs"
	ImportToVDI importOption = "importtovdi"
)
