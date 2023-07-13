package authcontroller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/ihksanghazi/go-auth-jwt/config"
	"github.com/ihksanghazi/go-auth-jwt/helpers"
	"github.com/ihksanghazi/go-auth-jwt/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Login(w http.ResponseWriter, r *http.Request) {
	// mengambil input dari json
	var userInput models.Users
	decoder:= json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInput); err!= nil {
		response:= map[string]string{"msg":err.Error()}
		helpers.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	// ambil data berdasarkan database
	var user models.Users
	if err := models.DB.Where("email = ?", userInput.Email).First(&user).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			response:= map[string]string{"msg":"Email Atau Password Salah"}
			helpers.ResponseJSON(w, http.StatusUnauthorized, response)
			return
		default :
			response:= map[string]string{"msg":err.Error()}
			helpers.ResponseJSON(w, http.StatusInternalServerError, response)
			return
		}
	}

	// pengecekan password
	if err:= bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(userInput.Password)); err != nil {
		response:= map[string]string{"msg":"Email Atau Password Salah"}
		helpers.ResponseJSON(w, http.StatusUnauthorized, response)
		return
	}

	// proses pembuatan token jwt
	expTime := time.Now().Add(time.Minute * 1)
	claims := &config.JWTClaim{
		Username: user.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: "go-auth-jwt",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	// mendeklarasi algoritma yang akan digunakan untuk login
	tokenAlgo:= jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// sign token
	token, err := tokenAlgo.SignedString(config.JWT_KEY)
	if err != nil {
		response:= map[string]string{"msg":err.Error()}
		helpers.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	// set cookie
	http.SetCookie(w,&http.Cookie{
		Name: "token",
		Path: "/",
		Value: token,
		HttpOnly: true,
	})

	response:= map[string]string{"msg":"Login Berhasil"}
	helpers.ResponseJSON(w, http.StatusOK, response)
	
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
	// set cookie
	http.SetCookie(w,&http.Cookie{
		Name: "token",
		Path: "/",
		Value: "",
		HttpOnly: true,
		MaxAge: -1,
	})

	response:= map[string]string{"msg":"Logout Berhasil"}
	helpers.ResponseJSON(w, http.StatusOK, response)
}