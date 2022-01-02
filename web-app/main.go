package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

type Page struct {
	Title string
	Body []byte
}

func (p *Page) save() error {
	f := p.Title + ".txt"
	return ioutil.WriteFile(f, p.Body, 0600) // 0600: for permissions
}

func load(title string) (*Page, error) {
	f := title + ".txt"
	body, err := ioutil.ReadFile(f)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func view(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/test/"):]
	p, _ := load(title)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}