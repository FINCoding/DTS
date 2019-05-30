package main

import (
    "fmt"
    "os"
    "io"
    "bufio"

    "github.com/DTS/models"
    "github.com/jinzhu/gorm"
     _ "github.com/jinzhu/gorm/dialects/sqlite"

)

var db *gorm.DB
var err error
var auto []models.Auto

func main() {
    //readfile()
    readfile2()
    //upd_db_tst()
    //tst_upd_db()
}



func print_err(err error) {
    if err != nil{
        fmt.Println(err)
        os.Exit(1)
    }
}

func upd_db_tst() {
    db, err = gorm.Open("sqlite3", "./gorm.db")
    if err != nil {
       fmt.Println(err)
    }
    defer db.Close()

    //auto1 := models.Auto{Brand:"Lada", Model:"Niva"}
    auto := models.Auto{}
    fmt.Println(db.First(&auto))

    /*fmt.Print(db.NewRecord(auto))
    db.Create(&auto1)
    fmt.Print(db.NewRecord(auto))

    fmt.Println(db.Where("model = ?", "Niva").First(&auto))*/
    //get_all_enries(db)

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
        item_auto := models.Auto{Brand: "Lada", Model:string(data[n])}
        fmt.Println("item_auto ", item_auto)
        fmt.Println("string(data[:n]) ", string(data[:n]))
        auto = append(auto, item_auto)
        //fmt.Println(item_auto)
        //fmt.Println(string(data[:n]))
    }
    //fmt.Println(auto)
}


func tst_upd_db() {

    type Document struct{
        gorm.Model
        Body string
    }

    db, err = gorm.Open("sqlite3", "./gorm.db")
    if err != nil {
       fmt.Println(err)
    }
    defer db.Close()

    //tst_doc := Document {
    //    Body: "This is a test document",
    //}

    //db.AutoMigrate(&Document{})
    //fmt.Println(db.HasTable(&Document{}))

    //db.Create(&tst_doc)
   // var doc []Document

    if err := db.Find(&auto).Error; err != nil {
        fmt.Println(err)
    } else {
        for i,val := range(auto) {
            fmt.Println(i, val)
        }

    }
}

func readfile2() {
    file, err := os.Open("./data_from_parsing_files/lada.txt")
    print_err(err)
    defer file.Close()

    db, err = gorm.Open("sqlite3", "./gorm.db")
    if err != nil {
       fmt.Println(err)
    }
    defer db.Close()

    reader := bufio.NewReader(file)

    for{
        line, _, err := reader.ReadLine()

        if err == io.EOF{
            break
        }

        item_auto := models.Auto{Brand: "Lada", Model:string(line)}
        auto = append(auto, item_auto)
        db.Create(&auto)
    }

    //fmt.Println(db.Find(&models.Auto{}))
}