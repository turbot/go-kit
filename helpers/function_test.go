package helpers

import "testing"

func TestGetCallingFunction(t *testing.T) {
	funcA(t)
}

func funcA(t *testing.T) {
	funcB(t)
}

func funcB(t *testing.T) {

	funcName := GetCallingFunction(0)
	if funcName != "funcB" {
		t.Errorf("expected GetCallingFunction(0) to return 'funcB' but it returned %s", funcName)
	}
	funcName = GetCallingFunction(1)
	if funcName != "funcA" {
		t.Errorf("expected GetCallingFunction(1) to return 'funcA' but it returned %s", funcName)
	}
	funcName = GetCallingFunction(2)
	if funcName != "TestGetCallingFunction" {
		t.Errorf("expected GetCallingFunction(1) to return 'TestGetCallingFunction' but it returned %s", funcName)
	}
}
