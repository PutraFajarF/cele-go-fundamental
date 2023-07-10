package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	City  string `json:"city"`
}

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Celerates!"))
	})
	http.HandleFunc("/user", getAllUsers)

	server := new(http.Server)
	server.Addr = ":9090"

	fmt.Println("Server started at localhost:8080")
	server.ListenAndServe()
}

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	users := []User{
		{
			Name:  "Putra F",
			Email: "putraf@gmail.com",
			City:  "bogor",
		},
		{
			Name:  "Budi H",
			Email: "budih@gmail.com",
			City:  "mojokerto",
		},
		{
			Name:  "prima",
			Email: "prima@gmail.com",
			City:  "semarang",
		},
	}

	userJson, err := json.Marshal(users)
	if err != nil {
		w.Write([]byte(err.Error()))
	}

	w.Write(userJson)
}
