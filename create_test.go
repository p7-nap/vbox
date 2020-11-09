package vbox

// import (
// 	"reflect"
// 	"testing"
// )
//
// func TestCreateVMOK(t *testing.T) {
// 	options := CreateOptions{
// 		Name:            "test",
// 		Basefolder:      ".",
// 		Register:        true,
// 		DefaultHardware: true,
// 		OSType:          Linux64,
// 	}
//
// 	// Apply Timeout option to verify arguments
// 	c := testClient([]OptionFunc{}, func(cmd string, args ...string) ([]byte, error) {
// 		// Verify correct command and arguments passed, including option flags
// 		if want, got := "vboxmanage", cmd; want != got {
// 			t.Fatalf("incorrect command:\n- want: %v\n-  got: %v",
// 				want, got)
// 		}
//
// 		wantArgs := []string{"createvm", "--name=test", "--basefolder=.", "--ostype=linux_64", "--register", "--default"}
// 		if want, got := wantArgs, args; !reflect.DeepEqual(want, got) {
// 			t.Fatalf("incorrect arguments\n- want: %v\n-  got: %v",
// 				want, got)
// 		}
//
// 		return nil, nil
// 	})
//
// 	if err := c.Create.CreateVM(options); err != nil {
// 		t.Fatalf("unexpected error for Client.VBox.Create: %v", err)
// 	}
// }
