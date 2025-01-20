package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title string `json:"title"`
	Author `json:"author"`
}

type Author struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

func main() {
	fmt.Println("Hello World")

	author := Author{Name: "hoge", Age: 12}
	book := Book{Title: "Learning Concurrency in Python", Author: author}

	fmt.Printf("%+v\n", book)

	byteArray, err := json.MarshalIndent(book, "", "	")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(byteArray))

	var book2 Book
	bookJson := `{"title": "onigashimaTaro", "author": {"name": "hogetaro", "age": 12}}`
	err = json.Unmarshal([]byte(bookJson), &book2)
	if err != nil {
		fmt.Println(err)
	}
	
	fmt.Printf("%+v\n", book2)
}