package arguments

import (
    "fmt"
)

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
        return TestMockUp{value:argument}
    } else {
        return Console{value:argument}
    }
}


type ArgOut interface {
    PrintLn()
}


// console output implementation


type Console struct {
    value string
}

func (this Console) PrintLn( ) {
    fmt.Println(this.value)
}


// The test mock up
type TestMockUp struct {
    value string
}

func (this TestMockUp) PrintLn() {

}
