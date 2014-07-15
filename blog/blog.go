package blog

import (
    //"fmt"
    //"os"
    //"appengine"
    "net/http"
    "html/template"
    "github.com/gorilla/mux"
)

func init() {

  r := mux.NewRouter()
  r.HandleFunc("/", HomeHandler)
  http.Handle("/", r)

}

func HomeHandler(w http.ResponseWriter, r *http.Request) {

  var hogeTmpl = template.Must(template.New("hoge").ParseFiles("blog/base.html", "blog/hoge.html"))
  hogeTmpl.ExecuteTemplate(w, "base", "Hoge")

}
