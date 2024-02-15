package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func responseWithError(w http.ResponseWriter,code int,message string){
	if (code > 499){
		log.Println("Internal server error: ",message)
	}
	type errorResponse struct {
		Error string `json:"error"` // "error":"something went wrong"
	}

	responseWithJSON(w,code,errorResponse{Error:message})
}

func responseWithJSON (w http.ResponseWriter,code int, payload interface{}){
	data,err := json.Marshal(payload)
	if err != nil{
		log.Printf("Failed to marshal payload %v",payload)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type","application/json")
	w.WriteHeader(code)
	w.Write(data)
}