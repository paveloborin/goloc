# goloc

Простой и быстрый сервис для локализации писем, а также других сообщений, построенный на стандартных библиотеках go:
"golang.org/x/text/language" и "golang.org/x/text/message"


Поиск шаблонов сообщений осуществлется в папке templates проекта.

Для хранения меток локализации используются yaml файлы (каждый язык в своем файле) из дирректории translations проекта.


## Пример использования в cmd/main.go 

import "github.com/paveloborin/goloc/pkg/service"

	data := &struct {
		Name     string
		LastName string
	}{Name: "Ellis", LastName: "James"}

	emailContent, err := service.ParseTemplate("hello.html", language.Russian, data)
	fmt.Println(emailContent)


## Запуск примера
make run