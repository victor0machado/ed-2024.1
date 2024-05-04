package pedidos

import (
	"api/dados"
	"encoding/json"
	"fmt"
	"net/http"
)

func Adicionar(w http.ResponseWriter, r *http.Request) {
	var novoPedido dados.Pedido

	err := json.NewDecoder(r.Body).Decode(&novoPedido)
	if err != nil {
		http.Error(w, "Erro ao decodificar dados", http.StatusBadRequest)
		return
	}

	err = dados.FilaPedidos.Adicionar(novoPedido.Delivery, novoPedido.IdProdutos)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	w.Write([]byte("Sucesso"))
	fmt.Println("Pedido adicionado com sucesso")
}

func Listar(w http.ResponseWriter, r *http.Request) {
	filaJSON, err := json.Marshal(dados.FilaPedidos.Listar())
	if err != nil {
		http.Error(w, "Erro ao converter mensagens para JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write(filaJSON)
	fmt.Println("Fila de pedidos obtida")
}
