package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (ch Handler) GetInfrastructure() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)[`id`]
		num, _ := strconv.Atoi(id)

		course, err := ch.svc.RepositoryInstance().GeInfrastructure(int64(num))
		if err != nil {
			http.Error(w, "Ошибка", http.StatusInternalServerError)
			return
		}

		// Преобразуем структуру в JSON
		jsonResponse, err := json.Marshal(course)
		if err != nil {
			http.Error(w, "Ошибка при преобразовании в JSON", http.StatusInternalServerError)
			return
		}
		// Устанавливаем заголовок Content-Type на application/json
		w.Header().Set("Content-Type", "application/json")
		// Отправляем JSON-ответ
		w.Write(jsonResponse)
	}
}
