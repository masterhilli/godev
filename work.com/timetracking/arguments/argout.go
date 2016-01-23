package arguments

import (
    "fmt"
)

/*
var ArgOutGetter ArgOutFactory

type ArgOutFactory struct {
    mockEnabled bool
    lastValue string
}

func (this *ArgOutFactory) EnableTestMockUp() {
    this.mockEnabled = true
}

func (this *ArgOutFactory) DisableTestMockUp() {
    this.mockEnabled = false
}

func (this *ArgOutFactory) SetLastArgument(value string) {
    this.lastValue = value
}

func (this *ArgOutFactory) GetLastArgument() string {
    return this.lastValue
}

func (this *ArgOutFactory) GetPrinter(argument string) ArgOut {
    if this.mockEnabled {
        return &TestMockUp{value:argument}
    } else {
        return &Console{value:argument}
    }
}

*/
type ArgOut interface {
    Println(format string, argument string,  args ...string)
    getValue() string
}


// console output implementation


type Console struct {
    value string
}

func (this *Console) Println(format string, argument string, args ...string) {
    //arguments := []string(args)
    fmt.Printf(format + "\n", argument, args)
    this.value = argument
}

func (this *Console) getValue() string {
    return this.value
}


// The test mock up
type TestMockUp struct {
    value string
}

func (this *TestMockUp) Println(format string,argument string, args ...string) {
    this.value = argument
}

func (this *TestMockUp) getValue() string {
    return this.value
}
