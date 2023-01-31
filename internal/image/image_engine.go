package image // For test, we will download the alpine image filesystem from https://www.alpinelinux.org/

import (
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
    "fmt"

	"github.com/realnfcs/didactic-container/internal"
	"github.com/realnfcs/didactic-container/internal/database"
)

type Filesystem struct {
	URL      string
	Name     string
	FileName string
}

func AlpineImage() {
	fs := &Filesystem{
		URL:      "https://dl-cdn.alpinelinux.org/alpine/v3.17/releases/x86_64/alpine-minirootfs-3.17.1-x86_64.tar.gz",
		Name:     "Alpine",
		FileName: "alpine-fs.tar.gz",
	}

	if err := fs.PullImage(); err != nil {
		panic(err)
	}

}

func UbuntuImage() {
	fs := &Filesystem{
		URL:      "https://cloud-images.ubuntu.com/minimal/releases/jammy/release/ubuntu-22.04-minimal-cloudimg-amd64-root.tar.xz",
		Name:     "Ubuntu",
		FileName: "ubuntu-fs.tar.xz",
	}

	if err := fs.PullImage(); err != nil {
		log.Println(err)
	}
}

// Function to download a image with the URL informed
func (fs *Filesystem) PullImage() error {

	response, err := http.Get(fs.URL)

	if err != nil {
		return err
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		return errors.New("Error in pull image request")
	}

	path := filepath.Join(internal.FS_FOLDER_PATH, fs.FileName)

	file, err := os.Create(path)

	if err != nil {
		return err
	}

	defer file.Close()

	_, err = io.Copy(file, response.Body)

	if err != nil {
		return err
	}

	database.InsertImage(fs.Name, fs.FileName, path)

	return nil

}

func DeleteImage(id, name, path string) {

    var err error

    if path == "" {
        path, err = database.SearchPath(id, name)
        if err != nil {
            log.Fatalln(err)
        }
    }

    err = database.DelImage(id, name, path)
    if err != nil {
        log.Fatalln(err)
    }

    fmt.Println("Image deleted from database")

    err = os.Remove(path)
    if err != nil {
        log.Fatalln(err)
    }

    fmt.Println("Image deleted succesfully!")
}
