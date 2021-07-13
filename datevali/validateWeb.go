package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

var templates = template.Must(template.ParseFiles("validate.html"))
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

type CheckLog struct{
	Title string
	Body []byte
}

func main() {
	http.HandleFunc("/validate/", makeHandler(checkHandler))
	log.Fatal(http.ListenAndServe(":8080",nil))
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

func checkHandler(w http.ResponseWriter, r *http.Request, rootPath string) {
	p, err := loadLog(rootPath)
	if err != nil {
		return
	}
	//do check process .
	beginValidate()
	renderTemplate(w, "view", p)
}

func renderTemplate(w http.ResponseWriter, title string, log *CheckLog) {
	t, err := template.ParseFiles(title+".log")
	if err !=  nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.ExecuteTemplate(w, title+".log", log)
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

func (c *CheckLog)saveLog(fileName string) error {
	fileName += ".log"
	err := ioutil.WriteFile(fileName, c.Body, 600)
	return err
}