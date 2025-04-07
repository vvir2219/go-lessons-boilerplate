package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

var RequestSecretKey struct{}

type SecretKey struct {
	Value string
}

type Todo struct {
	Title    string    `json:"title"`
	Done     *bool     `json:"done"`
	Deadline time.Time `json:"deadline"`
}

func (t *Todo) UnmarshalJSON_1(data []byte) error {
	var aux = &struct {
		Title    string `json:"title"`
		Done     *bool  `json:"done"`
		Deadline string `json:"deadline"`
	}{}

	err := json.Unmarshal(data, &aux)
	if err != nil {
		return err
	}

	t.Title = aux.Title
	t.Done = aux.Done
	t.Deadline, err = time.Parse(time.DateOnly, aux.Deadline)

	return err
}

func (t *Todo) UnmarshalJSON(data []byte) error {
	type alias Todo
	var aux = &struct {
		Deadline string `json:"deadline"`
		*alias
	}{
		alias: (*alias)(t),
	}

	err := json.Unmarshal(data, &aux)
	if err != nil {
		return err
	}

	t.Deadline, err = time.Parse("02.01.2006", aux.Deadline)

	return err
}

func main() {
	// GET / POST / PUT... <url> [ <body> ]
	router := chi.NewRouter()
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})

	router.Post("/todobuff", func(w http.ResponseWriter, r *http.Request) {
		buffer := bytes.Buffer{}
		buffer.ReadFrom(r.Body)

		response := fmt.Sprintf("Got body \"%s\"", buffer.String())
		w.Write([]byte(response))
	})

	router.Route("/todo", func(r chi.Router) {
		r.Use(func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				_, pass, ok := r.BasicAuth()
				if !ok {
					w.WriteHeader(http.StatusUnauthorized)
					return
				}

				if pass != "bla" {
					w.WriteHeader(http.StatusUnauthorized)
					return
				}

				fmt.Println("before check pass middleware")
				next.ServeHTTP(w, r)
				fmt.Println("after check pass middleware")
			})
		})

		r.Use(func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				ctx := r.Context()
				newCtx := context.WithValue(ctx, RequestSecretKey, SecretKey{Value: "SECRET_KEY"})

				fmt.Println("before put value middleware")
				next.ServeHTTP(w, r.WithContext(newCtx))
				fmt.Println("after put value middleware")
			})
		})

		// []middlewares
		// h
		// m3(m2(m1(h)))

		r.Post("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("func post todo")

			ctx := r.Context()
			val, ok := ctx.Value(RequestSecretKey).(SecretKey)
			if !ok {
				fmt.Println("could not get secret key out")
			} else {
				fmt.Println("got secret key", val)
			}

			var todo Todo
			err := json.NewDecoder(r.Body).Decode(&todo)
			if err != nil {
				log.Println("Got error: ", err)
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(err.Error()))
				return
			}

			if todo.Done == nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("Mandatory field \"done\""))
				return
			}

			marshaled, _ := json.Marshal(todo)

			response := fmt.Sprintf("Got body %s", marshaled)
			w.Write([]byte(response))
		})

		r.Get("/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("func get todo")
			var response string

			id := chi.URLParam(r, "id")
			page := r.URL.Query().Get("page")

			if page == "" {
				response = fmt.Sprintf("Hello with id %s", id)
			} else {
				response = fmt.Sprintf("Hello with id %s from page \"%s\"", id, page)
			}

			w.Write([]byte(response))
		})
	})

	fmt.Println("Server started")
	log.Fatal(http.ListenAndServe("localhost:3000", router))
}
