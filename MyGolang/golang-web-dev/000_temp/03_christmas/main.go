package main

import (
	"fmt"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("letter.phpasp")
	if err != nil {
		fmt.Println("There was an error parsing file", err)
	}
	friend := []string{"Alex", "Conor", "Ken", "Ronnie", "Patick", "Nina", "Jeremy", "Gentry", "Christian"}

	err = tpl.Execute(os.Stdout, friend)
	if err != nil {
		fmt.Println("There was an error executing template", err)
	}
}
