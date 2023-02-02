package image

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/realnfcs/didactic-container/internal"
	"github.com/realnfcs/didactic-container/internal/database"
	"github.com/realnfcs/didactic-container/internal/models"
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

    image := models.Image{
        Name: fs.Name,
        Filename: fs.FileName,
        Path: path,
    }

	image.InsertImage()

	return nil

}

func (fs *Filesystem) PullLocalImage() error {

    if !strings.Contains(fs.FileName, ".tar.") || !strings.Contains(fs.URL, ".tar.") {
        return errors.New("The file must be compressed in any .tar")
    }

    path := filepath.Join(internal.FS_FOLDER_PATH, fs.FileName)

	file, err := os.Create(path)

	defer file.Close()
    
    bytes, err := os.ReadFile(fs.URL)
    if err != nil {
        return err
    }

    err = os.WriteFile(path, bytes, 0755)
    if err != nil {
        return err
    }

	if err != nil {
		return err
	}

    image := models.Image{
        Name: fs.Name,
        Filename: fs.FileName,
        Path: path,
    }

	image.InsertImage()

	return nil
}

func DeleteImage(id, name, path string) {

    var err error

    image := models.Image{
        ID: id,
        Name: name,
        Path: path,
    }

    if image.Path == "" {
        image.Path, err = image.SearchPath()
        if err != nil {
            log.Fatalln(err)
        }
    }

    err = image.DelImage()
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
