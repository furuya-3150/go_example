package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var mySigningKey = os.Getenv("MY_JWT_TOKEN")

func GenerateJWT() (string, error) {
	if mySigningKey == "" {
		return "", fmt.Errorf("environment variable MY_JWT_TOKEN is not set")
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["user"] = "hogetaro"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString([]byte(mySigningKey))
	if err != nil {
		return "", fmt.Errorf("something went wrong: %s", err.Error())
	}

	return tokenString, nil
}

func homePage(w http.ResponseWriter, r *http.Request) {
	validToken, err := GenerateJWT()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://localhost:9000", nil)
	req.Header.Set("Token", validToken)
	res, err := client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	fmt.Fprintln(w, string(body))
}

func handleRequests() {
	http.HandleFunc("/", homePage)

	http.ListenAndServe(":9001", nil)
}

func main() {
	fmt.Println("My Simple Client")

	handleRequests()
}
