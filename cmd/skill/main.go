package main

import (
	"fmt"
	"github.com/ANkulagin/voice_mail/cmd/skill/flags"
	"net/http"
)

// функция main вызывается автоматически при запуске приложения
func main() {
	flags.ParseFlags()
	if err := run(); err != nil {
		panic(err)
	}
}

// функция run будет полезна при инициализации зависимостей сервера перед запуском
func run() error {
	fmt.Println("Running server on", flags.FlagRunAddr)
	return http.ListenAndServe(flags.FlagRunAddr, http.HandlerFunc(webhook))
}

// функция webhook — обработчик HTTP-запроса
func webhook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		// разрешаем только POST-запросы
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// установим правильный заголовок для типа данных
	w.Header().Set("Content-Type", "application/json")
	// пока установим ответ-заглушку, без проверки ошибок
	_, _ = w.Write([]byte(`
      {
        "response": {
          "text": "Извините, я пока ничего не умею"
        },
        "version": "1.0"
      }
    `))
}
