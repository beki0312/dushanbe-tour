package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

func (ch Handler) GetLanguage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(11111111)
		language, err := ch.svc.RepositoryInstance().GeLanguage()
		if err != nil {
			http.Error(w, "Ошибка при преобразовании в JSON", http.StatusInternalServerError)
			return
		}

		// Преобразуем структуру в JSON
		jsonResponse, err := json.Marshal(language)
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
