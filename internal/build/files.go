package build

import (
	"os"

	"github.com/saromanov/godownload"
)

// downloadRepo provides downloading of the repo
func downloadRepo(link string) error {
	os.MkdirAll("./repo", os.ModePerm)
	gd := &godownload.GoDownload{}
	gd.Download(link, nil)
	return nil
}
