package metricas

import (
	"api/dados"
	"encoding/json"
	"fmt"
	"net/http"
)

func Metricas(w http.ResponseWriter, r *http.Request) {
	metricasJSON, err := json.Marshal(dados.MetricasColetadas)
	if err != nil {
		http.Error(w, "Erro ao converter mensagens para JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write(metricasJSON)
	fmt.Println("Métricas coletadas com sucesso")
}
