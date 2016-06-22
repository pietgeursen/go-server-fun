package main

import (
  "fmt"
  "net/http"
  "log"
  "html/template"
)

func main() {
  fmt.Printf("hello, piet\n")

  tmpl, err := template.New("loc").Parse(`
    <h1>You got to {{.URL.Path}} </h1>
    <h2>Method was a {{.Method}} </h2>
  `)
  
  if err != nil {panic(err)}
  http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
    tmpl.Execute(w, r)
  })

  log.Fatal(http.ListenAndServe(":8080", nil))
}

func Adder(x, y int) int {
  return x + y 
}

