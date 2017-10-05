package handlers

import (
	"sync"
	"text/template"
	"net/http"
	"log"
	"path/filepath"
	"os"
)

const templatePath string = "views"

type TemplateHandler struct {
	once     sync.Once
	filename string
	tmpl     *template.Template
}

func NewTemplateHandler(filename string) *TemplateHandler {
	return &TemplateHandler{
		filename: filename,
	}
}

func (t *TemplateHandler) Parse() error {
	path := filepath.Join(templatePath, t.filename)

	_, err := os.Stat(path)
	if err != nil {
		return err
	}


	t.tmpl, err = template.New("index.html").Delims("<{", "}>").ParseFiles(path)
	if err != nil{
		return err
	}

	t.tmpl = template.Must(t.tmpl, err)
	return nil
}

func (t *TemplateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := t.Parse()
	if err != nil{
		log.Fatal(err)
	}
	/*
	t.once.Do(func() {
		err := t.Parse()
		if err != nil {
			log.Fatal(err)
		}
	})
	*/

	err = t.tmpl.Execute(w, r)
	if err != nil {
		log.Fatal(err)
	}
}
