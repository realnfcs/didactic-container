package main

import (
	"os"

	"github.com/realnfcs/didactic-container/src/run"
)

// This is a didactic project with learning propuse
// Cretids for
// 			Liz Rice 	@lizrice
// 			Fabio Akita @akitaonrails

// Project In Progress

// docker 			run image 	<cmd> <params>
// go run main.go 	run 		<cmd> <params>

func main() {
	switch os.Args[1] {
	case "run":
		run.Run()
	case "child":
		run.Child()
	default:
		panic("bad command")
	}
}
