package serverfx

import (
	"context"
	"fmt"
	"net/http"

	"go.uber.org/fx"
)

// Module allows the module to provide multiple constructors and functions to invoke
// through a single variable using fx.Options
var Module = fx.Options(
	fx.Provide(NewServeMux),
	fx.Invoke(Register),
)

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

// Register the http handler
func Register(handler http.Handler, mux *http.ServeMux) {
	// TODO: do the actual registering
	fmt.Printf("register the handler %#v against %T\n", handler, mux)
	mux.Handle("/", handler)

}
