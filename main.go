package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/arnavsharma93/fxdemo/fxdemofx"

	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fxdemofx.Module,
		fx.Provide(
			NewHandler,
		),
	)
	app.Run()

}

// NewHandler function instantiates our handler
func NewHandler(log *log.Logger) http.Handler {
	log.Print("Constructing a new http handler")

	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handler was called")
		_, err := io.WriteString(w, "hello world\n")
		if err != nil {
			log.Print("err, please do something ", err)
		}
	}

	return http.HandlerFunc(handler)
}
