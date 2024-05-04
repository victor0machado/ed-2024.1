package loja

import (
	"api/dados"
	"fmt"
	"net/http"
)

func Abrir(w http.ResponseWriter, r *http.Request) {
	err := dados.AbrirLoja()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	w.Write([]byte("Loja aberta"))
	fmt.Println("Loja aberta")
}

func Fechar(w http.ResponseWriter, r *http.Request) {
	err := dados.FecharLoja()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	w.Write([]byte("Loja fechada"))
	fmt.Println("Loja fechada")
}
