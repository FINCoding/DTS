package main

import (
    "fmt"
    "os"
    "io"

    "github.com/DTS/models"
    "github.com/jinzhu/gorm"
     _ "github.com/jinzhu/gorm/dialects/sqlite"

)

var db *gorm.DB

func main() {
    //readfile()
    upd_db_tst()
}

func print_err(err error) {
    if err != nil{
        fmt.Println(err)
        os.Exit(1)
    }
}


func upd_db_tst() {
    auto := models.Auto{Brand: "Lada", Model:"Niva"}

    fmt.Print(db.NewRecord(auto))
    db.Create(&auto)
    fmt.Print(db.NewRecord(auto))


}
func readfile() {
    file, err := os.Open("./data_from_parsing_files/lada.txt")
    print_err(err)

    defer file.Close()

    data := make([]byte, 64)

    for{
        n, err := file.Read(data)
        if err == io.EOF{
            break
        }
        fmt.Println(string(data[:n]))
    }
}
