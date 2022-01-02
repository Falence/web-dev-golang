package main

import (
	"net/http"
	"html/template"
	"io/ioutil"
)

type Page struct {
	Title string
	Body []byte
}

// save method
func (p *Page) save() error {
	f := p.Title + ".txt"
	return ioutil.WriteFile(f, p.Body, 0600) // 0600: for permissions
}

// load page content from file
func load(title string) (*Page, error) {
	f := title + ".txt"
	body, err := ioutil.ReadFile(f)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

// Render page
func view(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/test/"):]
	p, _ := load(title)
	// fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
	t, _ := template.ParseFiles("test.html")
	t.Execute(w, p)
}

func main() {
	p := &Page{Title: "Test", Body: []byte("Welcome to the Test page")}
	p.save()
	http.HandleFunc("/test/", view)
	http.ListenAndServe(":8000", nil)
}