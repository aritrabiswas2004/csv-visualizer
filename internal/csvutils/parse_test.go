package csvutils

import "testing"

// This is a bad test
// TODO: Fix this test, make it better
func TestCleanRow(t *testing.T) {
	singleRow := []string{`"some string"`, "23", `"test"`} // Just add test cases here
	output := CleanRow(singleRow)

	if output[0] != "some string" || output[1] != "23" || output[2] != "test" {
		t.Fatalf("incorrect output of CleanRow() -> %v\n", output)
	}
}
