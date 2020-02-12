package tools

import (
	"fmt"
	"testing"
)

func TestGoListAll(t *testing.T) {
	pkg, err := GoList(".", false)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%+v\n", pkg)

	pkgs, err := GoList(".", true)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%+v\n", pkgs)
}
