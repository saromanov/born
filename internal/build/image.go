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

// newImage creates init for creating of docker images
func newImage() (*image, error) {
	client, err := docker.NewClient(defaultEndpoint)
	if err != nil {
		return nil, err
	}
	return &image{
		client: client,
	}, nil
}

func (a *image) createImage(userID string, s structs.StepConfig) (string, error) {
	t := time.Now()
	inputbuf, outputbuf := bytes.NewBuffer(nil), bytes.NewBuffer(nil)
	tr := tar.NewWriter(inputbuf)
	body := fmt.Sprintf("%s\n", s.Image)
	body += addCommands(s.Commands)
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
	cointainerName := fmt.Sprintf("%s/%s", userID, randomString(imageNameLength))
	opts := docker.BuildImageOptions{
		Name:         cointainerName,
		InputStream:  inputbuf,
		OutputStream: outputbuf,
	}
	if err := a.client.BuildImage(opts); err != nil {
		return "", err
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
