package vbox

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

type Client struct {
	Create   *CreateService
	Medium   *MediumService
	Modify   *ModifyService
	Storage  *StorageService
	Start    *StartService
	Clone    *CloneService
	execFunc ExecFunc
	flags    []string
	sudo     bool
	debug    bool
}

type ExecFunc func(cmd string, args ...string) ([]byte, error)

func shellExec(cmd string, args ...string) ([]byte, error) {
	return exec.Command(cmd, args...).CombinedOutput()
}

// exec executes an ExecFunc using the values from cmd and args.
// The ExecFunc may shell out to an appropriate binary, or may be swapped
// for testing.
func (c *Client) exec(cmd string, args ...string) ([]byte, error) {
	// Prepend recurring flags before arguments
	flags := append(c.flags, args...)

	// If needed, prefix sudo.
	if c.sudo {
		flags = append([]string{cmd}, flags...)
		cmd = "sudo"
	}

	c.debugf("exec: %s %v", cmd, flags)
	// Execute execFunc with all flags and clean up any whitespace or
	// newlines from its output.
	out, err := c.execFunc(cmd, flags...)
	if out != nil {
		out = bytes.TrimSpace(out)
		c.debugf("exec: %q", string(out))
	}
	if err != nil {
		// Wrap errors in Error type for further introspection
		return nil, &Error{
			Out: out,
			Err: err,
		}
	}

	return out, nil
}

//New creates a new Client
func New() *Client {
	c := &Client{
		flags:    make([]string, 0),
		execFunc: shellExec,
	}

	cs := &CreateService{
		v: c,
	}
	c.Create = cs

	ms := &ModifyService{
		v: c,
	}
	c.Modify = ms

	mes := &MediumService{
		v: c,
	}
	mes.Create = &MediumCreate{
		v: c,
	}
	mes.Modify = &MediumModify{
		v: c,
	}
	mes.Clone = &MediumClone{
		v: c,
	}
	c.Medium = mes
	cls := &CloneService{
		v: c,
	}
	c.Clone = cls
	stor := &StorageService{
		c: c,
	}
	stor.Control = &StorageControl{
		c: c,
	}
	stor.Attach = &StorageAttach{
		c: c,
	}
	c.Storage = stor

	ss := &StartService{
		v: c,
	}
	c.Start = ss

	return c
}

func (c *Client) debugf(format string, a ...interface{}) {
	if !c.debug {
		return
	}

	log.Printf("VBoxManage: "+format, a...)
}

type Error struct {
	Out []byte
	Err error
}

// Error returns the string representation of an Error.
func (e *Error) Error() string {
	return fmt.Sprintf("%s: %s", e.Err, string(e.Out))
}

type OptionFunc func(c *Client)

// Debug returns an OptionFunc which enables debugging output for the Client
// type.
func Debug(enable bool) OptionFunc {
	return func(c *Client) {
		c.debug = enable
	}
}

// Exec returns an OptionFunc which sets an ExecFunc for use with a Client.
// This function should typically only be used in tests.
func Exec(fn ExecFunc) OptionFunc {
	return func(c *Client) {
		c.execFunc = fn
	}
}
