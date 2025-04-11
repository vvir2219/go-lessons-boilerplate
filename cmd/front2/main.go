package main

import (
	"curs1_boilerplate/cmd/front2/views/base"
	"curs1_boilerplate/cmd/front2/views/pages"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

/*
wgo -file '\.templ$' templ generate &
wgo -file .go go run .
*/

type TodoItem struct {
	Description string `json:"description"`
}

type User struct {
	FirstName       string `json:"firstName"`
	LastName        string `json:"lastName"`
	Email           string `json:"email"`
	Username        string `json:"username"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
	Date            string `json:"date"`
	Time            string `json:"time"`
}

func main() {
	router := chi.NewRouter()
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		todoList := []pages.TodoItemSpecial{
			{Id: "1", Selected: true, Value: "Cumparaturi"},
			{Id: "2", Selected: true, Value: "Jucarii"},
			{Id: "3", Selected: true, Value: "Tot felu"},
		}

		base.PageSkeleton(pages.Page4(todoList)).Render(r.Context(), w)
	})

	router.Post("/todo", func(w http.ResponseWriter, r *http.Request) {
		var todoItem TodoItem
		json.NewDecoder(r.Body).Decode(&todoItem)

		pages.Item(todoItem.Description).Render(r.Context(), w)
	})

	router.Post("/register", func(w http.ResponseWriter, r *http.Request) {
		var user User
		json.NewDecoder(r.Body).Decode(&user)

		fmt.Println(user.Date)
		fmt.Println(user.Time)
	})

	fmt.Println("Server started")
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
