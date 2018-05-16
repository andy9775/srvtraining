package web

import (
	"context"
	"net/http"

	"github.com/dimfeld/httptreemux"
)

// A Handler is a type that handles an http request within our own little mini
// framework. Add error handling capabilities.
type Handler func(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error

// App is the entrypoint into our application and what configures our context
// object for each of our http handlers. Feel free to add any configuration
// data/logic on this App struct
type App struct {
	*httptreemux.TreeMux
}

// New creates an App value that handle a set of routes for the application.
func New() *App {
	return &App{
		TreeMux: httptreemux.New(),
	}
}

// Handle is our mechanism for mounting Handlers for a given HTTP verb and path
// pair, this makes for really easy, convenient routing.
func (a *App) Handle(verb, path string, handler Handler) {

	// The function to execute for each request.
	h := func(w http.ResponseWriter, r *http.Request, params map[string]string) {

		// Time to start thinking about context.
		ctx := context.TODO()

		// Call the wrapped handler functions.
		if err := handler(ctx, w, r, params); err != nil {
			// TODO: Error Handling
		}
	}

	// Add this handler for the specified verb and route.
	a.TreeMux.Handle(verb, path, h)
}
