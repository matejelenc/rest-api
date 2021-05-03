package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/matejelenc/rest-api/data"
	"github.com/matejelenc/rest-api/security"
)

func Login(w http.ResponseWriter, r *http.Request) {

	var input map[string]string
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	var user data.Person
	foundEmail := data.DB.First(&user, "Email = ?", input["email"])
	if foundEmail.Error != nil {
		http.Error(w, "Invalid email", http.StatusBadRequest)
		return
	}

	err = security.VerifyPassword(user.Password, input["password"])
	if err != nil {
		http.Error(w, "Incorrect passord", http.StatusBadRequest)
		return
	}

	tokenString, err := security.GenerateJWT(&user)
	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	expirationTime := time.Now().Add(time.Minute * 30)

	cookie := &http.Cookie{
		Name:     "Token",
		Value:    tokenString,
		Expires:  expirationTime,
		HttpOnly: true,
	}

	http.SetCookie(w, cookie)
}
