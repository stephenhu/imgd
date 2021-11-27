package main

import (
	//"encoding/json"
	"net/http"
)


func crawlerHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPost:

		q := r.FormValue("q")
		//p := r.FormValue("p")

		jd := Jd{}
		tb := Tb{}

		for i := 1; i < 10; i++ {
			go jd.Search(q, i)
			go tb.Search(q, i)
		}

	case http.MethodGet:		
	case http.MethodPut:
	case http.MethodDelete:
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
	
} // crawlerHandler
