package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

// Endpoint handler for convert endpoint on server
func indexHandler(w http.ResponseWriter, req *http.Request) {
    b, err := os.ReadFile("indextemplate.tpl")
    if (err != nil) {
        log.Fatal("Could not read index template file")
    }

    tpl := string(b)
    
    title := "Test page"

    t, err := template.New("webpage").Parse(tpl)
    
    if (err != nil) {
        log.Fatal("Could not initialize template")
    }

    body := "This works"

    data := struct {
        Title string
        BodyText string
    }{
        Title: title,
        BodyText: body,
    }

    w.WriteHeader(http.StatusOK)
    w.Header().Set("Content-Type", "text/html; charset=utf-8")

    err = t.Execute(w, data)
}
