package executeCmds

import (
	//"bufio"
	"fmt"
	"regexp"
	"testing"
	"os/exec"
	//"io/ioutil"
	"os"
	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }
type MyDockerTestEngine struct{}
var _ = Suite(&MyDockerTestEngine{})

var primitiveDockerFile = 	"FROM docker/whalesay:latest\n" +  
							"RUN apt-get -y update && apt-get install -y fortunes\n" +
							"CMD /usr/games/fortune -a | cowsay\n"

var dockCreateImage = "mydockerCreateOwnImage"

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func (s *MyDockerTestEngine) TestRunHelloWorldDocker(c *C) {
	s.DockerRunWithArgument(c, "hello-world")
}

func (s *MyDockerTestEngine) TestRunWhaleSaySomething(c *C) {
	s.DockerRunWithArgument(c, "The whale has said something", "docker/whalesay", "cowsay")
}

func (s *MyDockerTestEngine) TestImages(c *C) {
	s.DockerImages(c, "masterhilli81/docker-whale")
} 

func (s *MyDockerTestEngine) TestCreateOwnDockerImage(c *C) {
	err := os.Mkdir(dockCreateImage, 0777)
	check(err)
	os.Chdir(dockCreateImage)
	file, err := os.OpenFile("Dockerfile", os.O_APPEND | os.O_CREATE | os.O_WRONLY,0644) 
	check(err)
	n, err := file.WriteString(primitiveDockerFile)
	check(err)
	fmt.Sprintf("%d\n", n)
	err = file.Sync()
	check(err)
	err = file.Close()
	check(err)
	c.Assert(nil, Equals, err)


	out, err := exec.Command("docker", "build", "-t", "docker-whale2", ".").Output()
	if err != nil {
		c.Fatal(err)
	}

	fmt.Sprintf("%s", out)
	os.Chdir("..")
	os.RemoveAll(dockCreateImage)
	s.DockerImages(c,"docker-whale2")

	out, err = exec.Command("docker", "rmi", "-f", "docker-whale2").Output()
	if err != nil {
		c.Fatal(err)
	}
}

func (s *MyDockerTestEngine) DockerRunWithArgument(c *C, checkstring string, cmdArgumets ...string) {
	lenargs := len (cmdArgumets) + 2
	args := make([]string, lenargs, lenargs)
	args[0] = "run"
	copy(args[1:lenargs], cmdArgumets)
	args[lenargs-1] = checkstring
	out := s.ExecuteCmd(c, "docker", args...)

	c.Assert(checkstring, Equals, s.FindSearchString(checkstring, out))
}

func (s *MyDockerTestEngine) DockerImages(c *C, checkstring string) {
	out := s.ExecuteCmd(c, "docker", "images")

	c.Assert(checkstring, Equals, s.FindSearchString(checkstring, out))
}

func (s *MyDockerTestEngine)  ExecuteCmd(c *C, name string, args ...string) string {
	out, err := exec.Command(name, args...).Output()
	if err != nil {
		c.Fatal(err)
	}
	 return fmt.Sprintf("%s", out)
}

func (s *MyDockerTestEngine) FindSearchString(searchstring string, cmdOutput string) string {
	re := regexp.MustCompile(searchstring)
	returnVal := re.FindStringSubmatch(cmdOutput)
	if returnVal == nil {
		return cmdOutput
	}
	return returnVal[0]
}