package utils

import "testing"

func TestCodeGenerate(t *testing.T) {
	// case 1: get UUID
	uuid := CodeGenerate()
	if len(uuid) == 0 {
		t.Fail()
	}
}
