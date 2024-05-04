package metricas

import (
	"fmt"
	"api/dados"
	"encoding/json"
	"net/http"
)

func Metricas(w http.ResponseWriter, r *http.Request) {
	var metricas dados.Metricas

	metricasJSON, err := json.Marshal(metricas)
	if err != nil {
		http.Error(w, "Erro ao converter mensagens para JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write(metricasJSON)
	fmt.Println("Métricas coletadas com sucesso")
}
