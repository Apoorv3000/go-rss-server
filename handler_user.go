package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Apoorv3000/go-server/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) createHandlerUser(w http.ResponseWriter, r *http.Request) {

	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}

	err := decoder.Decode(&params)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error parsing json: %v", err))
		return
	}

	user,err:=apiCfg.DB.CreateUser(r.Context(),database.CreateUserParams{
		ID: uuid.New(),
		Name : params.Name,
	})
	if err != nil {
		responseWithError(w, 500, fmt.Sprintf("Error creating user: %v", err))
		return
	}

	responseWithJSON(w, 200, databaseUserToUser(user))
	
}