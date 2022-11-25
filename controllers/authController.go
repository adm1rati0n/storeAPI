package controllers

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"storeAPI/dbConnection"
	"storeAPI/models"
	"time"
)

type Claims struct {
	Login string `json:"login"`
	Role  int    `json:"role, omitempty"`
	jwt.StandardClaims
}

var jwtKey = []byte("BrawlStars")

func (claims Claims) Valid() error {
	var now = time.Now().Unix()
	if claims.VerifyExpiresAt(now, true) && claims.VerifyIssuer("127.0.0.0", true) {
		return nil
	}
	return fmt.Errorf("Token is invalid")
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour).Unix()
	tokenStr, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func ValidateJWT(next func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Token"] != nil {
			token, err := jwt.Parse(r.Header["Token"][0], func(t *jwt.Token) (interface{}, error) {
				_, ok := t.Method.(*jwt.SigningMethodHMAC)
				if !ok {
					w.WriteHeader(http.StatusUnauthorized)
					w.Write([]byte("Not authorized"))
				}
				return jwtKey, nil
			})
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Not authorized: " + err.Error()))
			}
			if token.Valid {
				next(w, r)
			}
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Not authorized"))
		}
	})
}

func GetJWT(w http.ResponseWriter, r *http.Request) {
	if r.Header["Access"] != nil {
		if r.Header["Access"][0] == "123" {
			token, err := GenerateJWT()
			if err != nil {
				return
			}
			fmt.Fprintf(w, token)
		}
	}
}

func SignIn(w http.ResponseWriter, r *http.Request) {

}

func getAllUsers() []models.User {
	db := dbConnection.DB
	rows, err := db.Query("select Login, `Password` from `user`")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err = rows.Scan(&user.Login, &user.Password)
		if err != nil {
			panic(err)
		}
		users = append(users, user)
	}
	return users
}
