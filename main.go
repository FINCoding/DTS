package main

import (
    "net/http"
    "html/template"
	"fmt"

	"github.com/DTS/models"
    "github.com/go-martini/martini"
)

var discounts map[string]*models.Discount

func main() {
    m := martini.Classic()

    staticOptions := martini.StaticOptions{Prefix:"static"}
    m.Use(martini.Static("static", staticOptions))
    m.Get("/", indexHandler)
    m.Get("/write", writeHandler)
    m.Run()
}

func indexHandler(w http.ResponseWriter, r *http.Request){
    t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
    if err != nil{
        fmt.Fprintf(w, err.Error())
    }
    t.ExecuteTemplate(w, "index", nil)
}


func writeHandler(w http.ResponseWriter, r *http.Request){
    t, err := template.ParseFiles("templates/write.html", "templates/header.html", "templates/footer.html")
    if err != nil{
        fmt.Fprintf(w, err.Error())
    }
    t.ExecuteTemplate(w, "write", nil)
}
