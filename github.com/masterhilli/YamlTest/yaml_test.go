package YamlTest

import (
    . "gopkg.in/check.v1"
    yaml "gopkg.in/yaml.v2"
    "testing"
)

type Pair struct {
    Test    string
    Another string
}

type YamlWithList struct {
    Mylist       []string
    Mypairlist   [][]string
    Mymaplist    map[string][]string
    Myobjectlist map[string]Pair
}

type YamlTestEngine struct {
}

func TestYamlEngine(t* testing.T) {
    Suite(&YamlTestEngine{})
    TestingT(t)
}


func (this *YamlTestEngine) TestReadingInStringList(c *C) {
    yamlList := this.unmarshalList([]byte(contentForTestingYaml))
    c.Assert(len(yamlList.Mylist), Equals, 2)
    c.Assert(yamlList.Mylist[0], Equals, "test1")
    c.Assert(yamlList.Mylist[1], Equals, "test2")
    c.Assert(len(yamlList.Mypairlist), Equals, 3)
    c.Assert(len(yamlList.Mypairlist[0]), Equals, 2)
    c.Assert(len(yamlList.Mypairlist[1]), Equals, 2)
    c.Assert(len(yamlList.Mypairlist[2]), Equals, 3)
    c.Assert(len(yamlList.Mymaplist["abc"]), Equals, 2)
    c.Assert(yamlList.Mymaplist["abc"][1], Equals, "test2")
}

func (this *YamlTestEngine) unmarshalList(content []byte) YamlWithList {
    var yamlWithList YamlWithList
    err := yaml.Unmarshal(content, &yamlWithList)
    if err != nil {
        panic(err)
    }
    return yamlWithList
}



/***************************************************
Constant strings

*/

const contentForTestingYaml string = `
mylist: [test1, test2]
mypairlist: [
    [abc, def],
    [ghi, jkl],
    [mno, pqr, stv]
]
mymaplist:
    abc:
        - test
        - test2
    def:
        - again
        - again2
myobjectlist:
    mytest:
        test: my
        another: soon
    myOther:
        test: second
        another: so on
`