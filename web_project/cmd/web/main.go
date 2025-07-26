package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tpl, err := template.ParseFiles("templates/hello.gohtml")
	if err != nil {
		log.Printf("parsing template: %v", err)
		http.Error(w, "There was an error parsing the template", http.StatusInternalServerError)
		return
	}
	err = tpl.Execute(w, nil)
	if err != nil {
		log.Printf("executing template: %v", err)
		http.Error(w, "There was an error executing the template", http.StatusInternalServerError)
		return
	}
}

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
