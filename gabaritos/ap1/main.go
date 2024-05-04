package main

import (
	"fmt"
	"net/http"

	"api/handlers/produtos"
	"api/handlers/pedidos"
	"api/handlers/metricas"
	"api/handlers/loja"
	"api/processamento"

	"github.com/gorilla/mux"
)

func main() {
	// router
	r := mux.NewRouter()

	// rotas
	r.HandleFunc("/produto", produtos.Adicionar).Methods("POST")
	r.HandleFunc("/produto", produtos.Buscar).Methods("GET")
	r.HandleFunc("/produto", produtos.Remover).Methods("DELETE")

	r.HandleFunc("/produtos", produtos.Listar).Methods("GET")

	r.HandleFunc("/pedido", pedidos.Adicionar).Methods("POST")

	r.HandleFunc("/pedidos", pedidos.Listar).Methods("GET")

	r.HandleFunc("/metricas", metricas.Metricas).Methods("GET")

	r.HandleFunc("/abrir", loja.Abrir).Methods("POST")
	r.HandleFunc("/fechar", loja.Fechar).Methods("POST")

	// Processando dados através de goroutine
	go processamento.ProcessaPedidos()

	// inicialização do servidor HTTP
	fmt.Println("Servidor iniciado em http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
