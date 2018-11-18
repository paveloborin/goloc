package service

import (
	"bytes"
	"fmt"
	"html/template"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func ParseTemplate(tplName string, lang language.Tag, data interface{}) (string, error) {
	p := message.NewPrinter(lang)
	fmap := template.FuncMap{
		"translate": p.Sprintf,
	}

	t, err := template.New(tplName).Funcs(fmap).ParseFiles("./templates/" + tplName)
	if err != nil {
		return "", fmt.Errorf("cannot parse template")
	}

	buf := bytes.NewBuffer([]byte{})
	if err := t.Execute(buf, data); err != nil {
		return "", fmt.Errorf("cannot execute parse template")
	}

	return buf.String(), nil
}
