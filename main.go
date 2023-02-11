/*
Copyright © 2023 Nicolas Sales <nicolas.cavalcante.dev@gmail.com>

*/
package main

import (
	"github.com/realnfcs/didactic-container/cmd"
	"github.com/realnfcs/didactic-container/internal/database"
)

func main() {

	if err := database.NewSQLiteConnection(); err != nil {
		panic(err)
	}

	cmd.Execute()
}
