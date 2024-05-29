package order

import (
	"PizzaApi/pkg/client/postgres"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func Handler(client postgres.Client) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		OrderHandler(w, r, client)
	}

}

func OrderHandler(w http.ResponseWriter, r *http.Request, client postgres.Client) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	repository := NewRepository(client)
	switch r.Method {
	case http.MethodGet:
		id, err := strconv.Atoi(r.URL.Path[len("orders/ "):])
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
	case http.MethodPost:
		var order Order
		if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		i, err := repository.Insert(r.Context(), order)
		if err != nil {
			log.Default().Println(err)
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			if err := json.NewEncoder(w).Encode(i); err != nil {
				log.Default().Println(err)
				w.WriteHeader(http.StatusInternalServerError)
			}
		}
	case http.MethodOptions:
		w.WriteHeader(http.StatusOK)
	default:
		log.Default().Println(r.Body)
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
