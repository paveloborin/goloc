package service

import (
	"fmt"
	"io/ioutil"
	"strings"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/message/catalog"
	"gopkg.in/yaml.v2"
)

type dictionary struct {
	Data map[string]string
}

func (d *dictionary) Lookup(key string) (data string, ok bool) {
	if _, ok := d.Data[key]; !ok {
		return "", false
	}

	return "\x02" + d.Data[key], true
}

func init() {
	dict, err := parseYAMLDict()
	if err != nil {
		panic(err)
	}
	fallback := language.MustParse("en")
	cat, err := catalog.NewFromMap(dict, catalog.Fallback(fallback))
	if err != nil {
		panic(err)
	}
	message.DefaultCatalog = cat
}

func parseYAMLDict() (map[string]catalog.Dictionary, error) {
	dir := "./translations"
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	translations := map[string]catalog.Dictionary{}

	for _, f := range files {
		yamlFile, err := ioutil.ReadFile(dir + "/" + f.Name())
		if err != nil {
			return nil, err
		}
		data := map[string]string{}
		err = yaml.Unmarshal(yamlFile, &data)
		if err != nil {
			return nil, err
		}

		lang := strings.Split(f.Name(), ".")[0]

		translations[lang] = &dictionary{Data: data}
	}

	return translations, nil
}
