package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/matejelenc/rest-api/data"
	"github.com/matejelenc/rest-api/security"
)

// swagger:route POST /login users loginUser
// Login a user
//
// responses:
//	201: noContent
//  400: badRequestResponse

//Login verifies a user
func Login(w http.ResponseWriter, r *http.Request) {

	//deserialize email and password
	var input map[string]string
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	//check if the user is exists
	var user data.Person
	foundEmail := data.DB.First(&user, "Email = ?", input["email"])
	if foundEmail.Error != nil {
		http.Error(w, "Invalid email", http.StatusBadRequest)
		return
	}

	//verify the password
	err = security.VerifyPassword(user.Password, input["password"])
	if err != nil {
		http.Error(w, "Incorrect passord", http.StatusBadRequest)
		return
	}

	//generate a jwt token for this user
	tokenString, err := security.GenerateJWT(&user)
	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	expirationTime := time.Now().Add(time.Hour * 24)

	//save the jwt token in a cookie
	cookie := &http.Cookie{
		Name:     "Token",
		Value:    tokenString,
		Expires:  expirationTime,
		HttpOnly: true,
	}

	//set the cookie
	http.SetCookie(w, cookie)
	json.NewEncoder(w).Encode("Successfully logged in")
}
