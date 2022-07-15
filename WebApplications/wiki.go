package main

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"regexp"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"

	return os.WriteFile(filename, p.Body, 0600)
}
func loadPage(title string) *Page {
	filename := title + ".txt"
	body, _ := os.ReadFile(filename)
	return &Page{Title: title, Body: body}
}
func loadPageError(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there , I love %s !", r.URL.Path[1:])
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPageError(title)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)

}
func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPageError(title)
	if err != nil {
		p = &Page{Title: title}
	}
	fmt.Fprintf(w, "<h1>Editing %s</h1>"+
		"<form action=\"/save/%s\" method=\"POST\">"+
		"<textarea name=\"body\">%s</textarea><br>"+
		"<input type=\"submit\" value=\"Save\">"+
		"</form>",
		p.Title, p.Title, p.Body)
}

/*********html/template**************/
func editTempHandler(w http.ResponseWriter, r *http.Request) {
	title, err := getTitle(w, r)
	if err != nil {
		return
	}
	p, err := loadPageError(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}
func viewTempHandler(w http.ResponseWriter, r *http.Request) {
	title, err := getTitle(w, r)
	if err != nil {
		return
	}
	p, err := loadPageError(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}
func saveHandler(w http.ResponseWriter, r *http.Request) {
	title, err := getTitle(w, r)
	if err != nil {
		return
	}
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err = p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

//template caching
var thisPath = "./WebApplications/"
var templates = template.Must(template.ParseFiles(thisPath+"edit.html", thisPath+"view.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	/*	t ,err := template.ParseFiles(tmpl + ".html")
		if err != nil{
			http.Error(w,err.Error(),http.StatusInternalServerError)
			return
		}
		err = t.Execute(w,p)
		if err != nil{
			http.Error(w,err.Error(),http.StatusInternalServerError)
		}*/
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//validation
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return "", errors.New(" invalid Page title")
	}
	return m[2], nil //The title is the second subexpression
}

//引入函数闭包 Introducing Function Literals and Closures
func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Here we will extract the page title from the Request,
		// and call the provided handler 'fn'
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}
func viewHandler2(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPageError(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}
func editHandler2(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPageError(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}
func saveHandler2(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}
func main() {
	/*p1 := &Page{Title: "TestPage",Body : []byte("This is a sample Page .")}
	p1.save()
	p2, _ := loadPageError("TestPage")
	fmt.Println(string(p2.Body))*/

	//http.HandleFunc("/",handler)
	//http.HandleFunc("/view/",viewHandler)
	//http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/edit/", makeHandler(editHandler2))
	http.HandleFunc("/save/", makeHandler(saveHandler2))
	http.HandleFunc("/view/", makeHandler(viewHandler2))
	log.Fatal(http.ListenAndServe(":8080", nil))

}
