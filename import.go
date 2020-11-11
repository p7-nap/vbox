package vbox

type ImportService struct {
	c *Client
}

func (is *ImportService) exec(args ...string) ([]byte, error) {
	return is.c.exec("vboxmanage", args...)
}

//ImportVM
func (is *ImportService) ImportVM(vmname string, options ImportOptions) error {
	args := []string{"clonevm", vmname}
	importOptions := options.slice()
	args = append(args, importOptions...)
	_, err := is.exec(args...)
	return err
}

type ImportOptions struct {
}

func (c ImportOptions) slice() []string {
	var s []string
	return s
}
