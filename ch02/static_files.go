package main

import (
	_ "github.com/noissefnoc/go-for-everyone-2nd-edition/ch02/statik"
	"github.com/rakyll/statik/fs"
	"io"
	"log"
	"os"
)

func main() {
	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}
	f, err := statikFS.Open("/index.html")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	io.Copy(os.Stdout, f)
}
