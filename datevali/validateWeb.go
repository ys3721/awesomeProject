package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

var templates = template.Must(template.ParseFiles("validate.html"))
var validPath = regexp.MustCompile("^/(edit|save|view|validate|index)/([a-zA-Z0-9]+)$")

type CheckLog struct {
	Title string
	Body  []byte
}

func main() {
	//http.HandleFunc("/validate/", makeHandler(checkHandler))
	http.HandleFunc("/validate/", makeHandler(dispatcherHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func dispatcherHandler(w http.ResponseWriter, r *http.Request, action string) {
	switch action {
	case "index":
		renderTemplate(w, action, nil)
	case "golang":
		checkHandler(w, r, action)
	default:
		http.NotFound(w, r)
	}
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		m := validPath.FindStringSubmatch(request.URL.Path)
		if m == nil {
			http.NotFound(writer, request)
		}
		fn(writer, request, m[2])
	}
}

func checkHandler(w http.ResponseWriter, r *http.Request, action string) {
	rootPath := action
	//do check process .
	bytes, _ := BeginValidate()
	fmt.Printf("%s", bytes)
	p, err := loadLog(rootPath)
	if err != nil {
		p = &CheckLog{Title: rootPath, Body: bytes}
	} else {
		p = &CheckLog{Title: rootPath, Body: bytes}
	}
	p.saveLog(rootPath)
	renderTemplate(w, "validate", p)
}

func renderTemplate(w http.ResponseWriter, title string, log *CheckLog) {
	t, err := template.ParseFiles(title + ".html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.ExecuteTemplate(w, title+".html", log)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func loadLog(fileName string) (*CheckLog, error) {
	fileName += ".log"
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	cl := CheckLog{Title: fileName, Body: content}
	return &cl, nil
}

func (c *CheckLog) saveLog(fileName string) error {
	fileName += ".log"
	err := ioutil.WriteFile(fileName, c.Body, 755)
	return err
}
