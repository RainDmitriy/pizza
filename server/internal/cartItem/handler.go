package cartItem

import (
	"PizzaApi/pkg/client/postgres"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func Handler(client postgres.Client) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		CartHandler(w, r, client)
	}

}

func CartHandler(w http.ResponseWriter, r *http.Request, client postgres.Client) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Authorization, Origin, X-Requested-With, Accept, X-PINGOTHER, Content-Type")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	repository := NewRepository(client)
	switch r.Method {
	case http.MethodGet:
		id, err := strconv.Atoi(r.URL.Path[len("cart/ "):])
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
		var cartItem CartItem
		if err := json.NewDecoder(r.Body).Decode(&cartItem); err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		i, err := repository.Insert(r.Context(), cartItem)
		if err != nil {
			log.Default().Println(err)
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			if err := json.NewEncoder(w).Encode(i); err != nil {
				log.Default().Println(err)
				w.WriteHeader(http.StatusInternalServerError)
			}
		}
	case http.MethodPut:
		id, err := strconv.Atoi(r.URL.Path[len("cart/ "):])
		if err != nil {
			log.Default().Println(err)
		}
		if id != 0 {
			var cartItem CartItem
			if err := json.NewDecoder(r.Body).Decode(&cartItem); err != nil {
				log.Default().Println()
				w.WriteHeader(http.StatusBadRequest)
			}
			id, err := repository.Update(r.Context(), cartItem)
			if err != nil {
				log.Default().Println(err)
				w.WriteHeader(http.StatusInternalServerError)
			} else {
				if err := json.NewEncoder(w).Encode(id); err != nil {
					log.Default().Println(err)
					w.WriteHeader(http.StatusInternalServerError)
				}
			}
		}
	case http.MethodDelete:
		id, err := strconv.Atoi(r.URL.Path[len("cart/ "):])
		if err != nil {
			log.Default().Println(err)
		}
		if id != 0 {
			i, err := repository.Delete(r.Context(), id)
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
	case http.MethodOptions:
		w.WriteHeader(http.StatusOK)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
