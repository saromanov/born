package build

import (
	"archive/tar"
	"bytes"
	"errors"
	"fmt"
	"time"

	"github.com/fsouza/go-dockerclient"
	structs "github.com/saromanov/born/structs/v1"
)

const defaultEndpoint = "unix:///var/run/docker.sock"

var errImageNotDefined = errors.New("image is not defined")

// image defines structure for handling of Docker images
type image struct {
	client *docker.Client
	step   *structs.StepConfig
}

// newDockerClient provides initialization of docker client
func newDockerClient() (*docker.Client, error) {
	return docker.NewClient(defaultEndpoint)
}

// newImage creates init for creating of docker images
func newImage(c *docker.Client) *image {
	return &image{
		client: c,
	}
}

// pullImage provides pulling of the image
func (a *image) pullImage(name string) error {
	err := a.client.PullImage(docker.PullImageOptions{
		Repository: name,
	}, docker.AuthConfiguration{})
	if err != nil {
		return fmt.Errorf("unable to pull image: %v %s", err, name)
	}
	return nil
}
func (a *image) createImage(userID, stepName string, s BuildStep) (string, error) {
	t := time.Now()
	inputbuf, outputbuf := bytes.NewBuffer(nil), bytes.NewBuffer(nil)
	tr := tar.NewWriter(inputbuf)
	body := fmt.Sprintf("FROM %s\n", s.Image)
	body += fmt.Sprintf("COPY %s /usr/local/app\n", s.Path)
	body += fmt.Sprintf("RUN echo %s step", stepName)
	body += addCommands(s.Commands)
	fmt.Println(body)
	bodyBytes := []byte(body)
	tr.WriteHeader(&tar.Header{
		Name:       "Dockerfile",
		Size:       int64(len(bodyBytes)),
		ModTime:    t,
		AccessTime: t,
		ChangeTime: t,
	})
	tr.Write(bodyBytes)
	tr.Close()
	cointainerName := fmt.Sprintf("%s_%s", userID, randomString(imageNameLength))
	opts := docker.BuildImageOptions{
		Name:         cointainerName,
		InputStream:  inputbuf,
		OutputStream: outputbuf,
	}
	if err := a.client.BuildImage(opts); err != nil {
		return "", fmt.Errorf("unable to build image: %v", err)
	}
	return cointainerName, nil
}

func addCommands(c []string) string {
	var result string
	for i := 0; i < len(c); i++ {
		result += fmt.Sprintf("RUN %s", c[i])
	}
	return result
}
