package main

import (
    "fmt"
    "net/http"
    "net/url"
   // "io"
    "os"

    "github.com/PuerkitoBio/goquery"
)

var (
    WORKERS       int             = 1            //кол-во "потоков"
    REPORT_PERIOD int             = 10           //частота отчетов (сек)
    DUP_TO_STOP   int             = 500          //максимум повторов до останова
    HASH_FILE     string          = "hash.bin"   //файл с хешами
    FILE          string          = "data.txt" //файл с цитатами
)

func checkError(err error){
  if err != nil {
    panic(err)
    os.Exit(1)
  }
}

func parse() <-chan string{
    c := make(chan string)
    for i:=0; i < WORKERS; i++ {
        func() {
            for{
                response, err := http.Get("https://www.tts.ru/service/lada/")
                checkError(err)
                fmt.Print(response)

                u, err := url.Parse("https://www.tts.ru/service/lada/")
                checkError(err)
                fmt.Print(u)

                doc, err := goquery.NewDocument("https://www.tts.ru/service/lada/")
	            checkError(err)
	            fmt.Println(doc)

                /*defer response.Body.Close()
                doc, err := goquery.NewDocumentFromReader(io.Reader(response.Body))
                checkError(err)

                doc.Find("h3.r a").Each(func(i int, s *goquery.Selection) {
                    str, exists := s.Attr("href")
                    if exists {
                      u, err := url.Parse(str)
                      checkError(err)
                      m, _ := url.ParseQuery(u.RawQuery)
                      fmt.Println("\033[1;35m"+s.Text()+"\033[0m", m["q"][0])
                    } else {
                       fmt.Println(s.Text())
                    }
                }) */

            }
        }()
    }

    fmt.Println("Запущено потоков: ", WORKERS)
	return c
}

func main() {
    //parse()
    response, err := http.Get("https://www.yandex.ru")
    checkError(err)
    fmt.Print(response)

    //doc, err := goquery.NewDocument("http://yandex.ru/")
    //checkError(err)
    //fmt.Println(doc)
}

