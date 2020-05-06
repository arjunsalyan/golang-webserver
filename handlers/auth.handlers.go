package handlers

import (
	"encoding/json"
	"errors"
	"learning_project/auth"
	"learning_project/pkg/decodeJSONBody"
	"log"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var u auth.User

	err := decodeJSONBody.DecodeJSONBody(w, r, &u)
	if err != nil {
		var mr *decodeJSONBody.MalformedRequest
		if errors.As(err, &mr) {
			http.Error(w, mr.Msg, mr.Status)
		} else {
			log.Println(err.Error())
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

	if auth.SuperUser.Username != u.Username || auth.SuperUser.Password != u.Password {
		_ = json.NewEncoder(w).Encode(ResponseError{http.StatusUnauthorized, "invalid username or password"})
		return
	}

	ts, err := auth.CreateToken(auth.SuperUser.ID)
	if err != nil {
		json.NewEncoder(w).Encode(ResponseError{http.StatusUnprocessableEntity, err.Error()})
		return
	}
	saveErr := auth.CreateAuth(auth.SuperUser.ID, ts)
	if saveErr != nil {
		json.NewEncoder(w).Encode(ResponseError{http.StatusUnprocessableEntity, err.Error()})
	}
	tokens := map[string]string{
		"access_token":  ts.AccessToken,
		"refresh_token": ts.RefreshToken,
	}
	json.NewEncoder(w).Encode(tokens)
}
