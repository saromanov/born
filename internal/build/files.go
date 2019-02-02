package build

import (
	"archive/zip"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/saromanov/godownload"
)

var errDirNotFound = errors.New("cloned directory is not found")

// downloadRepo provides downloading of the repo
func downloadRepo(link, branch string) (string, error) {
	gd := &godownload.GoDownload{
		//Archive: "zip",
	}
	gd.Download(link, nil)
	err := unzip(branch, "app")
	if err != nil {
		return "", fmt.Errorf("unable to unzip repo: %v", err)
	}
	return getDownloadedPath("app")
}

func getDownloadedPath(path string) (string, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return "", fmt.Errorf("unable to open dir: %v", err)
	}
	if len(files) == 0 {
		return "", errDirNotFound
	}

	return fmt.Sprintf("%s/%s", path, files[0].Name()), nil
}

// unzip provides unzipping of the dir to output folder
func unzip(src, dest string) error {
	var filenames []string

	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {

		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer rc.Close()

		fpath := filepath.Join(dest, f.Name)
		if !strings.HasPrefix(fpath, filepath.Clean(dest)+string(os.PathSeparator)) {
			return fmt.Errorf("%s: illegal file path", fpath)
		}

		filenames = append(filenames, fpath)

		if f.FileInfo().IsDir() {

			// Make Folder
			os.MkdirAll(fpath, os.ModePerm)

		} else {

			// Make File
			if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
				return err
			}

			outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}

			_, err = io.Copy(outFile, rc)

			// Close the file without defer to close before next iteration of loop
			outFile.Close()

			if err != nil {
				return err
			}

		}
	}

	return nil
}
