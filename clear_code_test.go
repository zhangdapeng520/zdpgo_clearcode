package zdpgo_clearcode

import "testing"

func TestClearCode(t *testing.T) {
	filePath := "examples/test_data/level2_3.py"
	content, suffix, _ := GetFileContentAndSuffix(filePath)
	code, err := ClearCode(content, suffix)
	if err != nil {
		t.Error(err)
	}
	t.Log(code)
}

func TestFormat(t *testing.T) {
	filePath := "examples/test_data/level1_3.c"
	content, _, _ := GetFileContentAndSuffix(filePath)
	code := Format(content)
	t.Log(code)
}
