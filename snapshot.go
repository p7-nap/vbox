package vbox

import "fmt"

type SnapshotService struct {
	v *Client
}

func (ss *SnapshotService) exec(args ...string) ([]byte, error) {
	return ss.v.exec("vboxmanage", args...)
}

//Snapshot makes snapshot related operations on <vmname> with given <options>,...
func (ss *SnapshotService) Snapshot(vmname string, options SnapshotOptions) error {
	args := []string{"snapshot", vmname}
	snapshotOptions := options.slice()
	args = append(args, snapshotOptions...)
	_, err := ss.exec(args...)
	return err
}

type SnapshotOptions interface {
	slice() []string
}

func (ss *SnapshotService) TakeSnapshot(vmname string, options SnapshotOptions) error {
	opts := options.(SnapshotTakeOptions)
	args := []string{"snapshot", vmname}
	snapshotOptions := opts.slice()
	args = append(args, snapshotOptions...)
	_, err := ss.exec(args...)
	return err
}

//SnapshotOptions is passed to snapshot as desired options
type SnapshotgeneralOptions struct {
	Operation   snapshotOperation
	takeoptions *SnapshotTakeOptions
	//edit options
	NewName string
}

type SnapshotTakeOptions struct {
	SnapshotName string
	Description  string
	Live         bool
	Uniquename   string
}

func (o SnapshotTakeOptions) slice() []string {
	var s []string
	// if o.Operation != "" {
	// 	s = append(s, fmt.Sprintf("%s", o.Operation))
	// } else {
	// 	return s
	// }
	// if o.Operation == "restorecurrent" {
	// 	return s
	// }
	// if o.Operation == "list" {

	// }
	if o.SnapshotName != "" {
		s = append(s, fmt.Sprintf("%s", o.SnapshotName))
	}
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

type snapshotOperation string

const (
	TakeSnapshotWithOptions    snapshotOperation = "take"
	DeleteSnapshot             snapshotOperation = "delete"
	RestoreFromSnapshot        snapshotOperation = "restore"
	RestoreFromCurrentSnapshot snapshotOperation = "restorecurrent"
	EditSnapshot               snapshotOperation = "edit"
	ListSnapshots              snapshotOperation = "list"
	ShowSnapshotInfo           snapshotOperation = "showvminfo"
)
