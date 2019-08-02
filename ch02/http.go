package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// http request is logical path so use path package
		if ok, err := path.Match("/data/*.html", r.URL.Path); err != nil || !ok {
			http.NotFound(w, r)
			return
		}

		// after this line, path is physical path so use path/filepath
		name := filepath.Join(cwd, "data", filepath.Base(r.URL.Path)) // collect
		// name := filepath.Join(cwd, "data", path.Base(r.URL.Path)) // wrong
		f, err := os.Open(name)
		if err != nil {
			http.NotFound(w, r)
			return
		}
		defer f.Close()
		io.Copy(w, f)
	})
	http.HandleFunc("/upload", upload)
	http.HandleFunc("/upload_proper", uploadProper)
	http.ListenAndServe(":8080", nil)
}

func upload(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		_, header, _ := r.FormFile("file")
		s, _ := header.Open()
		p := filepath.Join("files", header.Filename)
		buf, _ := ioutil.ReadAll(s)
		ioutil.WriteFile(p, buf, 0644)
		http.Redirect(w, r, "/" + p, 301)
	} else {
		http.Redirect(w, r, "/", 301)
	}
}

func uploadProper(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		stream, header, err := r.FormFile("file")
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError),
				http.StatusInternalServerError)
			return
		}
		p := filepath.Join("files", filepath.Base(header.Filename))
		println(p)
		f, err := os.Create(p)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError),
				http.StatusInternalServerError)
			return
		}
		defer f.Close()
		io.Copy(f, stream)
		http.Redirect(w, r, path.Join("/files", p), 301)
	} else {
		http.Redirect(w, r, "/", 301)
	}
}
