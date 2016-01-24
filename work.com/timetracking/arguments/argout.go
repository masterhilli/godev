package arguments

import (
    "fmt"
)

type ArgOut interface {
    Println(format string, argument string,  args ...interface{})
    getValue() string
}

// console output implementation
type Console struct {
    value string
}

func (this *Console) Println(format string, argument string, args ...interface{}) {
    newArgList := make([]interface{}, 1)
    newArgList[0] = argument
    for i := range args {
        newArgList = append(newArgList, args[i])
    }
    fmt.Printf(format + "\n", newArgList...)
    this.value = argument
}

func (this Console) getValue() string { return this.value}

// The test mock up
type TestMockUp struct {
    value string
}

func (this *TestMockUp) Println(format string,argument string, args ...interface{}) {
    this.value = argument
}

func (this *TestMockUp) getValue() string {
    return this.value
}
