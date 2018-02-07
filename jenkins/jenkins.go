package integrations


import (
	"fmt"

	"github.com/bndr/gojenkins"
)

func NodeStatus() {
	nodes := jenkins.GetAllNodes()

	for _, node := range nodes {

		// Fetch Node Data
		node.Poll()
		if node.IsOnline() {
			fmt.Println("Node is Online")
		}
	}
}

func main() {
	jenkins, _ := gojenkins.CreateJenkins(nil, "http://localhost:8080/", "admin", "admin").Init()


const (
	PASS = iota
	FAIL
)

type Builder interface {
	Start(job string) error
	Stop(job string)
	Status(job string) (string, error)
	History(job string) ([]Build, error)
}

type Build struct {
	BuildNumber int
	job         string
	status      string
}

type Jenkins struct {
	BaseURL string
	user    string
	pass    string
	build   []Build
}

// Start will trigger a given job on jenkins
func (j *Jenkins) Start(job string) error {
	return nil
}

// Stop will trigger a given jenkins job to abort
func (j *Jenkins) Stop(job string) {
	return
}

// Status returns either PASS or FAIL for a given job's specific build
func (j *Jenkins) Status(job string, buildNumber int) (*Status, error) {
	return nil, nil
}

// setStatus is a helper for Status
func (j *Jenkins) setStatus() {

}

// History returns a list of Build objects, with their status and number
func (j *Jenkins) History(job string) ([]Build, error) {
	return nil, nil

}
