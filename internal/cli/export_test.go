package cli

import "testing"

func TestExportFile(t *testing.T) {
	file1Err := ExportFile("failfile.py", []byte("xxx"))
	if file1Err == nil {
		t.Fatalf("invalid format file created")
	}

	//file2Err := ExportFile("passfile.md", []byte(""))
	//if file2Err != nil {
	//	t.Fatalf("empty file cannot be created")
	//}
}
