package main

import (
	"encoding/json"
	"log"
	"net/http"
)


func picsHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:

		res, err := client.ZRevRangeWithScores(HASH_IMAGES, 0, -1).Result()

		if err != nil {
			log.Println(err)
		} else {

			j, err := json.Marshal(res)

			if err != nil {
				log.Println(err)
			} else {
				w.Write(j)
			}
	
		}

	case http.MethodPatch:

		u := r.FormValue("u")

		err := client.ZIncrBy(HASH_IMAGES, 1, u)

		if err != nil {
			log.Println(err)
		}

	case http.MethodPost:	
	case http.MethodPut:
	case http.MethodDelete:
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
	
} // picsHandler
