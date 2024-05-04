package produtos

import (
	"api/dados"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func Adicionar(w http.ResponseWriter, r *http.Request) {
	var novoProduto dados.Produto

	err := json.NewDecoder(r.Body).Decode(&novoProduto)
	if err != nil {
		http.Error(w, "Erro ao decodificar dados", http.StatusBadRequest)
		return
	}

	err = dados.ListaProdutos.Adicionar(novoProduto.Nome, novoProduto.Descricao, novoProduto.Valor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	w.Write([]byte("Sucesso"))
	fmt.Printf("Produto %s adicionado com sucesso\n", novoProduto.Nome)
}

func Remover(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))

	err := dados.ListaProdutos.Remover(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	w.Write([]byte("Sucesso"))
	fmt.Printf("Produto de id %d removido com sucesso\n", id)
}

func Buscar(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))

	prod, err := dados.ListaProdutos.Buscar(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	prodJSON, err := json.Marshal(prod)
	if err != nil {
		http.Error(w, "Erro ao converter mensagens para JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write(prodJSON)
	fmt.Printf("Produto com id %d localizado\n", id)
}

func Listar(w http.ResponseWriter, r *http.Request) {
	listaJSON, err := json.Marshal(dados.ListaProdutos.Listar())
	if err != nil {
		http.Error(w, "Erro ao converter mensagens para JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write(listaJSON)
	fmt.Println("Lista de produtos obtida")
}
