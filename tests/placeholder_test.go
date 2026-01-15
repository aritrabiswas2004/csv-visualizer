// Placeholder Test Functions only
// Correct Test functions to be added later

package tests

import (
	"strings"
	"testing"
)

func TestExampleFoo(t *testing.T) {
	rowSplit := strings.Split("foo,bar,alice,bob", ",")

	if rowSplit[0] != "foo" {
		t.Fatal()
	}
}

func TestExampleBar(t *testing.T) {
	rowSplit := strings.Split("foo,bar,alice,bob", ",")

	if rowSplit[1] != "bar" {
		t.Fatal()
	}
}

func TestExampleAlice(t *testing.T) {
	rowSplit := strings.Split("foo,bar,alice,bob", ",")

	if rowSplit[2] != "alice" {
		t.Fatal("not alice value is skjfn")
	}
}

func TestExampleBob(t *testing.T) {
	rowSplit := strings.Split("foo,bar,alice,bob", ",")

	if rowSplit[3] != "bob" {
		t.Fatal()
	}
}
