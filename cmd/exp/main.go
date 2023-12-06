package main

import (
	"html/template"
	"os"
)

type User struct {
	Name string
}

func main() {

	// first we need to parse
	t, err := template.ParseFiles("hello.html")
	if err != nil {
		panic(err)
	}

	// if everything parse correctly
	user := User{
		Name: "Kaisa",
	}

	// execute the template
	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}

}
