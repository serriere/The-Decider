package main

import (
  "fmt"
  "io/ioutil"
  "net/http"
  "html/template"
)

type Page struct {
  Title string
  Body []byte
}

type Decision struct {
  Page1 string
  Page2 string
  Response1 string
}

func fetch() {
  // Do the stuff to retreive external pages
}

func loadPage(title string) (*Page, error) {
  filename := title + ".txt"
  body, err := ioutil.ReadFile(filename)
  if err != nil {
    return nil, err
  }
  return &Page{Title: title, Body: body}, nil
}

// This should probably always load the main page
func mainPageHandler(w http.ResponseWriter, r *http.Request) {
  //title := r.URL.Path[len("/view/"):]
  pageName := "main"
  p, _ := loadPage(pageName)
  fmt.Fprintf(w, "<h1>Welcome to The Decider!</h1><div>%s</div>", p.Body)
}

func decideHandler(w http.ResponseWriter, r *http.Request) {

  // Quickly make a page with a basic title and no body
  p := Page{Title: "the Decision maker"}

  // Load the template file that we're going to dump that page into
  t, _ := template.ParseFiles("decider.html")
  t.Execute(w,p)

}

func answerHandler(w http.ResponseWriter, r *http.Request) {
  d := Decision{}

  d.Page1 = r.FormValue("PAGE1")
  d.Page2 = r.FormValue("PAGE2")

  res, err1 := http.Get(d.Page1)
  if err1 != nil {
    //handle error
  }
  responseBody, err2 := ioutil.ReadAll(res.Body)
  res.Body.Close()
  if err2 != nil {
    //handle error
  }
  d.Response1 = string(responseBody[:])

  t, _ := template.ParseFiles("answer.html")
  t.Execute(w,d)
}

func main() {
  http.HandleFunc("/answer/", answerHandler)
  http.HandleFunc("/decide/", decideHandler)
  fmt.Println("Listening on port 8080...")
  http.ListenAndServe(":8080", nil)

}
