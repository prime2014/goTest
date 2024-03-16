package accounts

import (
	"db"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func convertStrToInt(value string) int {
	myValue, _ := strconv.Atoi(value)
	return myValue
}

func SignUpController(w http.ResponseWriter, r *http.Request) {
	var user Users
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}

	fmt.Println(user)

	u, err := user.Save(db.Db)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	encoder := json.NewEncoder(w)
	encoder.Encode(u)
}

func GetUsersController(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	limit := convertStrToInt(r.URL.Query().Get("limit"))
	offset := convertStrToInt(r.URL.Query().Get("offset"))

	users, err := GetUsers(db.Db, offset, limit)
	fmt.Println(users)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	encoder := json.NewEncoder(w)
	encoder.Encode(users)

}

func LoginController(w http.ResponseWriter, r *http.Request) {

	var auth AuthenticationStruct
	json.NewDecoder(r.Body).Decode(&auth)

	token, err := auth.Login(db.Db)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(token)

}
