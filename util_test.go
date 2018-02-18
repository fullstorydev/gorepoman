package gorepoman

import (
	"reflect"
	"testing"
)

func TestCombineEnv(t *testing.T) {
	orig := []string{
		"A=1",
		"B=2",
		"C=3",
		"D=",
		"EFG=HIJKLMNOP",
	}
	overrides := []string{
		"A=10",
		"EFG=",
		"QRS=TUVWXYZ",
	}
	combined := combineEnv(orig, overrides)
	expected := []string {
		"A=10",
		"B=2",
		"C=3",
		"D=",
		"EFG=",
		"QRS=TUVWXYZ",
	}
	if !reflect.DeepEqual(expected, combined) {
		t.Fatalf("combineEnv produced wrong result: expecting [%v]; got [%v]", expected, combined)
	}
}
