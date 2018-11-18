package main

import (
	"fmt"
	"github.com/paveloborin/goloc/pkg/service"
	"github.com/rs/zerolog/log"
	"golang.org/x/text/language"
)

func main() {
	data := struct {
		Name     string
		LastName string
	}{Name: "Ellis", LastName: "James"}

	lang := language.Russian

	htmlContent, err := service.ParseTemplate("hello.html", lang, data)
	if err != nil {
		log.Warn().Err(err).Msg("cannot parse html template")
		panic(err)
	}

	fmt.Println(htmlContent)
}
