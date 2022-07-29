package main

import (
	"loja/routes"
	"net/http"
	"text/template"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)

}
