package main

import (
	"fmt"
	"net/http"

	"api/handlers"
	"api/handlers/lista"
	"api/processamento"

	"github.com/gorilla/mux"
)

func main() {
	// router
	r := mux.NewRouter()

	// rotas
	r.HandleFunc("/olamundo", handlers.OlaMundo).Methods("GET")

	r.HandleFunc("/lista", lista.Adicionar).Methods("POST")
	r.HandleFunc("/lista", lista.Obter).Methods("GET")

	r.HandleFunc("/listacomid", lista.AdicionarComID).Methods("POST")

	// Processando dados através de goroutine
	go processamento.ProcessaDados()

	// inicialização do servidor HTTP
	fmt.Println("Servidor iniciado em http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
