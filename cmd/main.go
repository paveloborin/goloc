package main

import (
	"fmt"

	"github.com/paveloborin/goloc/pkg/service"
	"golang.org/x/text/language"
)

func main() {
	data := &struct {
		Name     string
		LastName string
	}{Name: "Ellis", LastName: "James"}

	htmlContent, err := service.ParseTemplate("hello.html", language.Russian, data)
	if err != nil {
		panic(err)
	}

	fmt.Println(htmlContent)
}
