package executeCmds

import "fmt"
import "testing"
import "os/exec"



func TestSimpleExecutionOfLSCmdWithArgument(t *testing.T) {
	out := executeCmd(t, "ls", "-la")
	fmt.Printf("output is %s \n", out)
}


func TestExecutDockerWhaleAndSayHello (t *testing.T) {
	fmt.Printf("\n%s\n", executeCmd(t, "docker", "run", "docker/whalesay", "cowsay", "This is awesome!!!"))
}


func executeCmd(t *testing.T, name string, args ...string) []byte {
	out, err := exec.Command(name, args...).Output()
	if err != nil {
		t.Fatal(err)
	}
	return out
}
