package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(NewHandler),
	)
	app.Run()

}

// NewHandler function instantiates our handler
func NewHandler() http.Handler {
	fmt.Println("Constructing a new http handler")

	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handler was called")
		_, err := io.WriteString(w, "hello world\n")
		if err != nil {
			log.Println("err, please do something ", err)
		}
	}

	return http.HandlerFunc(handler)
}
