package main

import (
    "net/http"
    "html/template"
	"fmt"

	"github.com/DTS/models"
    "github.com/go-martini/martini"
    "github.com/martini-contrib/render"
    "github.com/jinzhu/gorm"
     _ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error
var discounts map[string]*models.Discount

func main() {
    m := martini.Classic()
    CreateDBTable()
    m.Use(render.Renderer(render.Options{
		Directory:  "templates",                         // Specify what path to load the templates from.
		Layout:     "layout",                            // Specify a layout template. Layouts can call {{ yield }} to render the current template.
		Extensions: []string{".tmpl", ".html"},          // Specify extensions to load for templates.
		//Funcs:      []template.FuncMap{unescapeFuncMap}, // Specify helper function maps for templates to access.
		Charset:    "UTF-8",                             // Sets encoding for json and html content-types. Default is "UTF-8".
		IndentJSON: true,                                // Output human readable JSON
	}))

    staticOptions := martini.StaticOptions{Prefix:"static"}
    m.Use(martini.Static("static", staticOptions))
    m.Get("/", indexHandler)
    m.Get("/write", writeHandler)
    m.Get("/addsc", addServiceCentre)
    m.Run()
}

func addServiceCentre() {

}

func CreateDBTable() {
   db, err = gorm.Open("sqlite3", "./gorm.db")
    if err != nil {
       fmt.Println(err)
    }
    defer db.Close()
    db.AutoMigrate(&models.Discount{}, &models.Service_centre{}, &models.Auto{})
    fmt.Print("it's", db.HasTable(&models.Discount{}))
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
