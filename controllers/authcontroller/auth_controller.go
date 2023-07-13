package authcontroller

import (
	"encoding/json"
	"net/http"

	"github.com/ihksanghazi/go-auth-jwt/helpers"
	"github.com/ihksanghazi/go-auth-jwt/models"
	"golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter, r *http.Request) {

}

func Register(w http.ResponseWriter, r *http.Request) {
	// mengambil input dari json
	var userInput models.Users
	decoder:= json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInput); err!= nil {
		response:= map[string]string{"msg":err.Error()}
		helpers.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	// hash password
	hashPassword,_:= bcrypt.GenerateFromPassword([]byte(userInput.Password),bcrypt.DefaultCost)
	userInput.Password = string(hashPassword)

	// insert ke database
	if err := models.DB.Create(&userInput).Error; err != nil {
		response:= map[string]string{"msg":"Internal Server Error"}
		helpers.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	// mengembalikan json
	response:= map[string]string{"msg":"Berhasil Register"}
	helpers.ResponseJSON(w, http.StatusOK, response)
}

func Logout(w http.ResponseWriter, r *http.Request) {

}