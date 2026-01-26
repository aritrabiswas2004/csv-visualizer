package cli

import "testing"

func TestParseArgs(t *testing.T) {
	opts, err := ParseArgs([]string{"viz", "-h"})
	if err != nil {
		t.Fatal(err)
	}
	if !opts.Help {
		t.Errorf("expected Help=true")
	}
}
