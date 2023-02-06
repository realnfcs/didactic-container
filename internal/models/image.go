package models

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/google/uuid"
	"github.com/realnfcs/didactic-container/internal/database"
)

type Image struct {
	ID       string
	Filename string
	Path     string
	Name     string
}

func (i *Image) CreateImageTable() {
	createTableSQL := `
		CREATE TABLE IF NOT EXISTS image (
			"id" TEXT NOT NULL PRIMARY KEY,
			"name" TEXT NOT NULL,
			"filename" TEXT NOT NULL,
			"path" TEXT NOT NULL
		)
	`

	err := database.ExecStatement(createTableSQL)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Image table was created!")
}

func (i *Image) InfoImages() {

	row, err := database.Query("SELECT * FROM image ORDER BY name")
	if err != nil {
		log.Fatalln(err)
	}

	defer row.Close()

	var (
		id       string
		name     string
		filename string
		path     string
	)

	for row.Next() {
		row.Scan(&id, &name, &filename, &path)
		fmt.Printf("[%s] name: %s\tfilename: %s\tpath: %s\n", id, name, filename, path)
	}
}

func (i *Image) SearchPath() (string, error) {

	if i.ID != "" {
		pathSearchWithIdQuerySQL := `
        SELECT path FROM image
        WHERE image.id = ? 
    `

		row := database.QueryRow(pathSearchWithIdQuerySQL, i.ID)
		if row.Err() != nil {
			return "", row.Err()
		}

		var path string

        err := row.Scan(&path)
        if err != nil {
            return "", err
        }

		if path == "" {
			return "", errors.New("error in scan")
		}

		return path, nil

	} else if i.Name != "" {
		pathSearchWithNameQuerySQL := `
        SELECT path FROM image
        WHERE image.name = ? 
    `

		row := database.QueryRow(pathSearchWithNameQuerySQL, i.Name)
		if row.Err() != nil {
			return "", row.Err()
		}

		var path string

        err := row.Scan(&path)
        if err != nil {
            return "", err
        }

		if path == "" {
			return "", errors.New("error in scan")
		}

		return path, nil

	}

	return "", errors.New("invalid params: you must insert name or id of image")
}

func (i *Image) InsertImage() {
	insertImageSQL := `
		INSERT INTO image (id, name, filename, path)
		VALUES (?, ?, ?, ?)
	`

	id := uuid.New().String()
	id = strings.Replace(id, "-", "", -1)

	err := database.ExecStatement(insertImageSQL, id, i.Name, i.Filename, i.Path)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Image has inserted successfully!")
}

func (i *Image) DelImageWithId() error {
	deleleImageSQL := `
        DELETE FROM image
        WHERE  image.id = ? 
    `

	err := database.ExecStatement(deleleImageSQL, i.ID)
	if err != nil {
		return err
	}

	log.Println("Image has deleted successfully!")
	return nil
}

func (i *Image) DelImageWithName() error {
	deleleImageSQL := `
        DELETE FROM image
        WHERE  image.name = ? 

    `

	err := database.ExecStatement(deleleImageSQL, i.Name)
	if err != nil {
		return err
	}

	log.Println("Image has deleted successfully!")
	return nil
}
