package image // For test, we will download the alpine image filesystem from https://www.alpinelinux.org/

import (
	"errors"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/realnfcs/didactic-container/internal"
	"github.com/realnfcs/didactic-container/internal/database"
)

type Filesystem struct {
	URL      string
	fileName string
}

func AlpineImage() {
	fs := &Filesystem{
		URL:      "https://dl-cdn.alpinelinux.org/alpine/v3.17/releases/x86_64/alpine-minirootfs-3.17.1-x86_64.tar.gz",
		fileName: "alpine-fs.tar.gz",
	}

	if err := fs.PullImage(); err != nil {
		panic(err)
	}

}

func (fs *Filesystem) PullImage() error {

	response, err := http.Get(fs.URL)

	if err != nil {
		return err
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		return errors.New("Error in pull image request")
	}

	path := filepath.Join(internal.FS_FOLDER_PATH, fs.fileName)

	file, err := os.Create(path)

	if err != nil {
		return err
	}

	defer file.Close()

	_, err = io.Copy(file, response.Body)

	if err != nil {
		return err
	}

	database.InsertImage(fs.fileName, path)

	return nil

}
