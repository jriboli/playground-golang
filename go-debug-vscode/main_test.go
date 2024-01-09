package main

import "testing"

// Since we have setup the .vscode folder and launch config we can 'Debug" in vscode
// this includes these test files too

func Test_add(t *testing.T) {
	a, b, c := 1, 2, 3

	res := add(a, b)

	if res != c {
		t.Fail()
	}
}
