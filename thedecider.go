package main

import (
  //"fmt"
  "log"
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


// Handles loading of the main page "/"
func mainHandler(w http.ResponseWriter, r *http.Request) {
  log.Println("Loading main page...")

  // Quickly make a page with a basic title and no body
  p := Page{Title: "the Decision maker"}

  // Load the template file that we're going to dump that page into
  t, _ := template.ParseFiles("decider.html")
  t.Execute(w,p)

}

// Handles loading of the decision page "/decide/"
func decideHandler(httpResponseWriter http.ResponseWriter, httpRequest *http.Request) {
  log.Println("Loading decision page...")

  d := Decision{}

  d.Page1 = httpRequest.FormValue("PAGE1")
  d.Page2 = httpRequest.FormValue("PAGE2")

  // Fetch the pages for comparison
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
  t.Execute(httpResponseWriter,d)
}

// Main execution
func main() {
  // Handlers
  http.HandleFunc("/", mainHandler)
  http.HandleFunc("/decide/", decideHandler)

  // Log start of server
  log.Println("About to listen on port 8080...")

  // Start server
  httpErr := http.ListenAndServe(":8080", nil)
  if httpErr != nil {
    log.Fatal("ListenAndServe: ", httpErr)
  }


}
