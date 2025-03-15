package main

import (
	"encoding/json"
	"net/http"
)

func PostJob(w http.ResponseWriter, r *http.Request) {
	data := NewJobRequest{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		respondWithErr(w, 500, err.Error())
		return
	}
	respondWithJson(w, 200, map[string]string{"msg": "accapted"})

}
