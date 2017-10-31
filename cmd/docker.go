package main

import (
	"fmt"
	"github.com/korovaisdead/go-simple-membership/utils/docker"
	"os/exec"
)

var (
	mongoImage = "mongo"
)

func main() {
	fmt.Println("start deploying the docker with mongo")
	if _, err := exec.LookPath("docker"); err != nil {
		panic("Don't hace docker installed in os")
	}

	fmt.Println("Check image")
	if ok, err := docker.DockerHaveImage(mongoImage); !ok || err != nil {
		if err != nil {
			panic(fmt.Sprintf("Error running docker to check for %s: %v", mongoImage, err))
		}
		if err := docker.DockerPull(mongoImage); err != nil {
			panic(fmt.Sprintf("Error pulling %s: %v", mongoImage, err))
		}
	}

	fmt.Println("Running the image")
	_, err := docker.DockerRun("-d", "-p", "27018:27017", mongoImage)
	if err != nil {
		panic("failed to run docker container")
	}
}