package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func AllUsers(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Could not connect to the database")
	}
	defer db.Close()

	var users []User
	db.Find(&users)

	json.NewEncoder(w).Encode(users)
}

func NewUser(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Could not connect to the database")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	db.Create(&User{Name: name, Email: email})

	fmt.Fprintf(w, "New User Successfully Created")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Could not connect to the database")
	}
	defer db.Close()

	var user User
	vars := mux.Vars(r)
	name := vars["name"]

	db.Where("name = ?", name).Find(&user)
	db.Delete(&user)

	fmt.Fprintf(w, "User Successfully Deleted")
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Could not connect to the database")
	}
	defer db.Close()

	var user User
	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	db.Where("name = ?" , name).Find(&user)

	user.Email = email
	db.Save(&user)

	fmt.Fprintf(w, "User Successfully Updated")
}