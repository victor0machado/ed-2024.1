package handlers

import (
	"fmt"
	"net/http"
)

func OlaMundo(w http.ResponseWriter, r *http.Request) {
	// Define o código da resposta à chamada
	w.WriteHeader(http.StatusOK)

	// Define o tipo de conteúdo que será enviado
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")

	// Enviar o conteúdo
	w.Write([]byte("Olá, mundo!"))

	// Fazer o log da função -- opcional
	fmt.Println("Olá, mundo!")
}
