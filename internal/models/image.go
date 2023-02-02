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
		row.Scan(&id, &name,  &filename, &path)
        fmt.Printf("[%s] name: %s\tfilename: %s\tpath: %s\n", id, name, filename, path)
	}
}

func (i *Image) SearchPath() (string, error) {

    pathSearchQuerySQL := `
        SELECT path FROM image
        WHERE image.id == ? 
        OR image.name == ?
    `

    row, err := database.Query(pathSearchQuerySQL, i.ID, i.Name) 
    if err != nil {
        return "", err
    }

    defer row.Close() 

    var path string

    for row.Next() {
        row.Scan(&path)
    }
    
    if path == "" {
        return "", errors.New("error in scan")
    }

    return path, nil
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

func (i *Image)  DelImage() error {
    deleleImageSQL := `
        DELETE FROM image
        WHERE  image.id = ? 
            OR image.name = ?
            OR image.path = ?

    `

    err := database.ExecStatement(deleleImageSQL, i.ID, i.Name, i.Path)
    if err != nil {
        return err
    }

    log.Println("Image has deleted successfully!")
    return nil
}

