package router

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
)

type Box struct {
	Content string `json:"content"`
}

func NewRouter() *mux.Router {

	boxes := map[string]string{}

	r := mux.NewRouter()

	r.HandleFunc("/ping", func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "pong")
	})

	// GET /boxes/{id}
	r.HandleFunc("/boxes/{id}", func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		switch r.Method {
		case "GET":
			if box, ok := boxes[params["id"]]; ok {
				resp := Box{
					Content: box,
				}
				jResp, err := json.MarshalIndent(resp, "", " ")
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}
				w.Header().Set("Content-Type", "application/json")
				w.Write(jResp)
			} else {
				w.WriteHeader(http.StatusNotFound)
			}

		case "PUT":
			body, err := io.ReadAll(r.Body)
			if err != nil {
				log.Fatal("Error reading body", err)
			}

			var box Box
			err = json.Unmarshal(body, &box)
			if err != nil {
				log.Fatal("Error unmarshaling box", err)
			}
			log.Printf("Adding box[%s] = %+v", params["id"], box)
			boxes[params["id"]] = box.Content

			w.WriteHeader(http.StatusCreated)
		}

	}).Methods("GET", "PUT")

	return r
}
