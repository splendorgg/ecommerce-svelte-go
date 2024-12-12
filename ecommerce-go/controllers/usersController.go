package controllers

import (
	"encoding/json"
	"fmt"
	"main/database"
	"main/headers"
	"main/models"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var SecretKey = []byte("AJISHDFAJSDHSgsjsdfsdjfsd18E13IUODH1JE12JEIASDAKSDSDA")
var DB *gorm.DB

func Signup(w http.ResponseWriter, r *http.Request) {
	headers.SetHeaders(&w)
	//! yeni
	dsn := "DBKEY"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB.AutoMigrate(&models.User{})
	//!
	//Get the email/password from request body
	var body struct {
		Email    string
		Password string
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		http.Error(w, "failed to hash password", http.StatusBadRequest)
		return
	}
	//Create the user
	user := models.User{Email: body.Email, Password: string(hash)}

	result := database.DB.Create(&user)
	if result.Error != nil {
		http.Error(w, "error creating user", http.StatusBadRequest)
		return
	}
	w.Header().Add("content-type", "application/json")
	//w.WriteHeader(http.StatusCreated)

	//CREATE TOKEN
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 1).Unix(),
	})
	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		http.Error(w, "failed to create token", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})

}

func Login(w http.ResponseWriter, r *http.Request) {
	headers.SetHeaders(&w)
	//Get the email/password from request body
	var body struct {
		Email    string
		Password string
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Println("bodyden decodelarken hata")
		return
	}
	//Look up requested user
	var user models.User
	database.DB.First(&user, "email = ?", body.Email)
	if user.ID == 0 {
		http.Error(w, "user not found", http.StatusBadRequest)
		return
	}
	//Check password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		http.Error(w, "invalid password", http.StatusBadRequest)
		fmt.Println("invalid password")
		return
	}

	//CREATE TOKEN
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":     user.ID,
		"exp":     time.Now().Add(time.Hour * 1).Unix(),
	})
	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		http.Error(w, "failed to create token", http.StatusInternalServerError)
		return
	}
	
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})


}

func Validate(w http.ResponseWriter, r *http.Request) {
	fmt.Println("im logged in")
	fmt.Println(r.Header)

	var user models.User
	database.DB.First(&user, "tokenstring= ?", r.Header.Get("Authorization"))
	if user.ID == 0 {
		http.Error(w, "invalid email or password", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)

}


