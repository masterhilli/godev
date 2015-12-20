package helper

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func ReadInFile(path string) []byte {
	filename, errAbs := filepath.Abs(path)
	PanicOnError(errAbs)
	content, errReadFile := ioutil.ReadFile(filename)
	PanicOnError(errReadFile)
	return content
}

func PrintStringArrayForTables(name string, values []string) {
	fmt.Printf("Values for %s", name)
	if values != nil {
		for value := range values {
			fmt.Printf(": %s", values[value])
		}
		fmt.Printf(": LEN: %d\n", len(values))
	} else {
		fmt.Printf("no values")
	}
}

func PanicOnError(e error) {
	if e != nil {
		panic(e)
	}
}
