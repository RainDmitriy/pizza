package item

import (
	"PizzaApi/pkg/client/postgres"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func Handler(client postgres.Client) func(w http.ResponseWriter, r *http.Request) {
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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(r.URL.Path[len("item/ "):])
	if err != nil {
		log.Default().Println(err)
	}
	if id == 0 {
		i, err := repository.GetAll(r.Context())
		if err != nil {
			log.Default().Println(err)
			w.WriteHeader(http.StatusInternalServerError)
		} else {

			if err := json.NewEncoder(w).Encode(i); err != nil {
				log.Default().Println(err)
				w.WriteHeader(http.StatusInternalServerError)
			}
		}
	} else {
		i, err := repository.GetOne(r.Context(), id)
		if err != nil {
			log.Default().Println(err)
			w.WriteHeader(http.StatusInternalServerError)
		} else {

			if err := json.NewEncoder(w).Encode(i); err != nil {
				log.Default().Println(err)
				w.WriteHeader(http.StatusInternalServerError)
			}
		}
	}
}
