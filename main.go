package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"

	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
			NewHandler,
			NewServeMux,
		),
		fx.Invoke(Register),
	)
	app.Run()

}

// NewServeMux provides a new serve mux
func NewServeMux(lifecycle fx.Lifecycle) *http.ServeMux {
	fmt.Println("construting a new serve mux")
	mux := http.NewServeMux()

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			fmt.Println("starting http server")
			// ignoring error handling for brevity
			go server.ListenAndServe()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			fmt.Println("stopping http server")
			return server.Shutdown(ctx)
		},
	})

	return mux
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

// Register the http handler
func Register(handler http.Handler, mux *http.ServeMux) {
	// TODO: do the actual registering
	fmt.Printf("register the handler %#v against %T\n", handler, mux)
	mux.Handle("/", handler)

}
