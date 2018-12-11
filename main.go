package main

import (
    "github.com/tst_martini/models"
    "github.com/go-martini/martini"
)

var discounts map[string]*models.discounts

func main() {
    discounts := make(map[string]*models.discounts,0)
    m := martini.Classic()

    staticOptions := martini.StaticOptions{Prefix:"static"}
    m.Use(martini.Static("static", staticOptions))
    m.Get("/", indexHandler)
    m.Get("/test", func()string{
    return "That's ok!"
    })

    m.Run()
}

func indexHandler( ){

}