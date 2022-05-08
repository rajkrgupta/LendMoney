package processor

import (
	"testing"
)

var getInputsFromFileMoc = func(path string) []string {
	var fileLines []string
	fileLines = append(fileLines, "LOAN IDIDI Dale 5000 1 6")
	fileLines = append(fileLines, "LOAN MBI Harry 10000 3 7")
	fileLines = append(fileLines, "LOAN UON Shelly 15000 2 9")

	fileLines = append(fileLines, "PAYMENT IDIDI Dale 1000 5")
	fileLines = append(fileLines, "PAYMENT MBI Harry 5000 10")
	fileLines = append(fileLines, "PAYMENT UON Shelly 7000 12")

	fileLines = append(fileLines, "BALANCE IDIDI Dale 3")
	fileLines = append(fileLines, "BALANCE IDIDI Dale 6")
	fileLines = append(fileLines, "BALANCE UON Shelly 12")
	fileLines = append(fileLines, "BALANCE MBI Harry 12")
	return fileLines
}

func TestProcess(t *testing.T) {

	origGetInputsFromFile := getInputsFromFile
	getInputsFromFile = getInputsFromFileMoc

	fileProc := FileProcessor{FilePath: ""}
	got := fileProc.Process()
	want := true

	getInputsFromFile = origGetInputsFromFile

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}
