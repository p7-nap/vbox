package vbox

// testClient creates a new Client with the specified OptionFuncs applied and
// using the specified ExecFunc.
func testClient(options []OptionFunc, fn ExecFunc) *Client {
	options = append(options, Exec(fn))
	c := New()
	return c
}
