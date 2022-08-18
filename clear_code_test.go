package zdpgo_clearcode

import "testing"

func TestClearCode(t *testing.T) {
	code, err := ClearCode("examples/test_data/level2_3.py")
	if err != nil {
		t.Error(err)
	}
	t.Log(code)
}
