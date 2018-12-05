package main

import (
	"fmt"
	"log"
	"os"

	"github.com/flexzuu/exifgql/files"
	zglob "github.com/mattn/go-zglob"
)

func main() {
	argsWithoutProg := os.Args[1:]
	glob := argsWithoutProg[0]
	fmt.Println(argsWithoutProg)
	paths, err := zglob.Glob(glob)
	if err != nil {
		log.Fatal(err)
	}
	for _, p := range paths {
		photo, err := files.OpenPhoto(p)
		if err != nil {
			log.Fatal(err)
		}
		model, err := photo.Model()
		if err != nil {
			log.Fatal(err)
		}
		log.Println(photo.Name(), model)
	}
}
