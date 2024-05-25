package item

import (
	"PizzaApi/pkg/client/postgres"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func GetItemHandler(client postgres.Client) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ItemHandler(w, r, client)
	}

}

func ItemHandler(w http.ResponseWriter, r *http.Request, client postgres.Client) {
	repository := NewRepository(client)
	switch r.Method {
	case http.MethodGet:
		GetItem(w, r, repository)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func GetItem(w http.ResponseWriter, r *http.Request, repository Repository) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		log.Default().Println(err)
	}
	if id == 0 {
		i, err := repository.FindAll(r.Context())
		if err != nil {
			log.Default().Println(err)
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Content-Type", "application/json")

			if err := json.NewEncoder(w).Encode(i); err != nil {
				log.Default().Println(err)
				w.WriteHeader(http.StatusInternalServerError)
			}
		}
	} else {
		i, err := repository.FindOne(r.Context(), id)
		if err != nil {
			log.Default().Println(err)
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Access-Control-Allow-Origin", "*")

			if err := json.NewEncoder(w).Encode(i); err != nil {
				log.Default().Println(err)
				w.WriteHeader(http.StatusInternalServerError)
			}
		}
	}
}
