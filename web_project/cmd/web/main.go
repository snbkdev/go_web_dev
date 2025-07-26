package main

import (
	"fmt"
	"os"
	"text/template"
)

type User struct {
	Name string
	Age int
	Meta UserMeta
}

type UserMeta struct {
	Visits int
}

func main() {
	fmt.Println("Starting with templates")
	t, err := template.ParseFiles("templates/hello.gohtml")
	if err != nil {
		panic(err)
	}

	user := User{Name: "John Doe", Age: 27, Meta: UserMeta{Visits: 8,}}

	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
}
