package vbox

import "fmt"

type SnapshotService struct {
	v *Client
}

func (ss *SnapshotService) exec(args ...string) ([]byte, error) {
	return ss.v.exec("vboxmanage", args...)
}

//Snapshot makes snapshot related operations on <vmname> with given <options>,...
func (ss *SnapshotService) Snapshot(vmname string) error {
	args := []string{"snapshot", vmname}
	_, err := ss.exec(args...)
	return err
}

func (ss *SnapshotService) SnapshotTake(vmname string, snapshotName string, options SnapshotTakeOptions) error {
	args := []string{"snapshot", vmname, "take", snapshotName}
	Options := options.slice()
	args = append(args, Options...)
	_, err := ss.exec(args...)
	return err
}

type SnapshotTakeOptions struct {
	Description string
	Live        bool
	Uniquename  string
}

func (o SnapshotTakeOptions) slice() []string {
	var s []string
	if o.Description != "" {
		s = append(s, fmt.Sprintf("--description=%s", o.Description))
	}
	if o.Live {
		s = append(s, fmt.Sprintf("--live"))
	}
	if o.Uniquename != "" {
		s = append(s, fmt.Sprintf("--uniquename=%s", o.Uniquename))
	}
	return s
}

func (ss *SnapshotService) SnapshotDelete(vmname string, snapshotName string) error {
	args := []string{"snapshot", vmname, "delete", snapshotName}
	_, err := ss.exec(args...)
	return err
}

func (ss *SnapshotService) SnapshotRestore(vmname string, snapshotName string) error {
	args := []string{"snapshot", vmname, "restore", snapshotName}
	_, err := ss.exec(args...)
	return err
}

func (ss *SnapshotService) SnapshotRestoreCurrent(vmname string, snapshotName string) error {
	args := []string{"snapshot", vmname, "restore", snapshotName}
	_, err := ss.exec(args...)
	return err
}

func (ss *SnapshotService) SnapshotEdit(vmname string, snapshotName string, options SnapshotEditOptions) error {
	args := []string{"snapshot", vmname, "edit", snapshotName}
	Options := options.slice()
	args = append(args, Options...)
	_, err := ss.exec(args...)
	return err
}

type SnapshotEditOptions struct {
	Description string
	NewName     string
}

func (o SnapshotEditOptions) slice() []string {
	var s []string
	if o.Description != "" {
		s = append(s, fmt.Sprintf("--description=%s", o.Description))
	}
	if o.NewName != "" {
		s = append(s, fmt.Sprintf("--name=%s", o.NewName))
	}
	return s
}

func (ss *SnapshotService) SnapshotList(vmname string, snapshotName string, options SnapshotListOptions) error {
	args := []string{"snapshot", vmname, "list", snapshotName}
	Options := options.slice()
	args = append(args, Options...)
	_, err := ss.exec(args...)
	return err
}

type SnapshotListOptions struct {
	ListMode Listmode
}

func (o SnapshotListOptions) slice() []string {
	var s []string
	if o.ListMode != "" {
		s = append(s, fmt.Sprintf("--%s", o.ListMode))
	}
	return s
}

type Listmode string

const (
	Details         Listmode = "details"
	Machinereadable Listmode = "machinereadable"
)

func (ss *SnapshotService) SnapshotShowVMInfo(vmname string, snapshotName string) error {
	args := []string{"snapshot", vmname, "showvminfo", snapshotName}
	_, err := ss.exec(args...)
	return err
}

// type snapshotOperation string

// const (
// 	TakeSnapshotWithOptions    snapshotOperation = "take"
// 	DeleteSnapshot             snapshotOperation = "delete"
// 	RestoreFromSnapshot        snapshotOperation = "restore"
// 	RestoreFromCurrentSnapshot snapshotOperation = "restorecurrent"
// 	EditSnapshot               snapshotOperation = "edit"
// 	ListSnapshots              snapshotOperation = "list"
// 	ShowSnapshotInfo           snapshotOperation = "showvminfo"
// )
