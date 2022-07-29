package controllers

import (
	"html/template"
	"log"
	"loja/models"
	"net/http"
	"strconv"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {

	todosOsProdutos := models.BuscaTodosOsProdutos()
	temp.ExecuteTemplate(w, "Index", todosOsProdutos)

}

func New(w http.ResponseWriter, r *http.Request) {

	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoCovertidoParaFloat, err := strconv.ParseFloat(preco, 64)

		if err != nil {
			log.Println("Erro na conversao do preço: ", err)
		}

		quantidadeConvertidaParaInt, err := strconv.Atoi(quantidade)

		if err != nil {
			log.Println("Erro na conversao do quantidade: ", err)
		}
		
		models.CriaNovoProduto (nome, descricao , precoCovertidoParaFloat , quantidadeConvertidaParaInt)

	}
	http.Redirect(w,r,"/",301)
}

func Delete (w http.ResponseWriter, r *http.Request){
	idDoProduto := r.URL.Query().Get("id")
	models.DeletaProduto(idDoProduto)
	http.Redirect(w, r, "/", 301)
}


func Edit (w http.ResponseWriter, r *http.Request){
		
	idDoProduto := r.URL.Query().Get("id")
	produto := models.EditaProduto(idDoProduto)
	temp.ExecuteTemplate(w, "Edit", produto)

}

func Update (w http.ResponseWriter, r *http.Request){
	
	if r.Method =="POST"{
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		idConvertidoParaInt, err := strconv.Atoi(id)
		if err !=  nil {
			log.Panicln("Erro na conversão do ID para int:",err)
		}

		precoCovertidoParaFloat, err := strconv.ParseFloat(preco,64)
		if err !=  nil {
			log.Panicln("Erro na conversão do preco para Float:",err)
		}

		quantidadeConvertidaParaInt, err := strconv.Atoi(quantidade)
		if err !=  nil {
			log.Panicln("Erro na conversão da quantidade para int:",err)
		}

		models.AtualizaProduto(idConvertidoParaInt, nome, descricao, precoCovertidoParaFloat, quantidadeConvertidaParaInt)

		http.Redirect(w,r, "/", 301)
	}


}
