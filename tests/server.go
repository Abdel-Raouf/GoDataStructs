package main

import (
    "fmt"
    "github.com/julienschmidt/httprouter"
    "net/http"
    "log"
    "text/template"
    "os"
)

type Node struct {
  Header string
  Footer string
  Body string
}

func (me Node) PrintAll() string {
  return me.Header + me.Body + me.Footer
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

    tpl, err := template.ParseFiles(ps.ByName("index"))
    if err != nil{
      os.Exit(0)
    }

    err = tpl.Execute(w, Node{
      Header: "This is my Header",
      Body: "This is my Body",
      Footer: "This is my Footer",
    })
    if err != nil {
      os.Exit(0)
    }
}

func main() {
    router := httprouter.New()
    router.GET("/:index", Hello)

    log.Fatal(http.ListenAndServe(":8080", router))
}
