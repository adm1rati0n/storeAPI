package controllers

import (
	"database/sql"
	"encoding/json"
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

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		panic(err)
	}
	return string(bytes)
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
	db := dbConnection.DB
	var auth models.AuthRequest
	err := json.NewDecoder(r.Body).Decode(&auth)
	if err != nil {
		panic(err)
	}
	var user models.User
	err = db.QueryRow("select * from `user` where IsDeleted = 0 and `Login` = ?", &auth.Login).Scan(
		&user.IDUser, &user.Login, &user.Password,
		&user.Employee, &user.Role, &user.IsDeleted)
	switch {
	case err == sql.ErrNoRows:
		json.NewEncoder(w).Encode("Неправильно введен логин или пароль")
	case err != nil:
		panic(err)
	default:
		if CheckPasswordHash(auth.Password, user.Password) {
			json.NewEncoder(w).Encode(user)
		} else {
			json.NewEncoder(w).Encode("Неправильно введен логин или пароль")
		}
	}
}
