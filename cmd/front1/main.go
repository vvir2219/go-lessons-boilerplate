package main

import (
	"curs1_boilerplate/cmd/front1/view"
	"fmt"
	"net/http"

	"github.com/a-h/templ"
)

func main() {
	// component := view.Hello("World")
	component := view.Base()

	http.Handle("/", templ.Handler(component))

	fmt.Println("Listening on :3000")
	http.ListenAndServe(
		":3000",
		nil,
	)
}
