package golden_files

import (
	"os"
	"testing"
)

// TODO: Implement a function that generates a summary report from input data.
// Then, write tests that compare the output against golden files.
//
// Your task:
// - Implement the GenerateReport function to produce a formatted report
// - Write a test that reads the expected output from testdata/expected.txt
// - Compare the generated output with the golden file
// - The test should pass when outputs match

func GenerateReport(data map[string]int) string {
	// Your implementation here
	if len(data) == 0 {
		return ""
	}
	return ""
}

func TestGenerateReport(t *testing.T) {
	data := map[string]int{
		"Apples":   10,
		"Bananas":  20,
		"Cherries": 15,
	}

	// TODO: Read from golden file and compare
	got := GenerateReport(data)
	_ = got

	// expected, err := os.ReadFile("testdata/expected.txt")
	// if err != nil {
	//     t.Fatal(err)
	// }
	//
	// if string(expected) != got {
	//     t.Errorf("mismatch:\ngot:\n%s\nwant:\n%s", got, string(expected))
	// }

	_ = os.Stderr // use os package
	t.Skip("Golden file test not yet implemented")
}
