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

