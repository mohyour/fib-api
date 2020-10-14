package app

import (
	"encoding/json"
	"net/http"
	handler "pex/app/handlers"

	"github.com/gorilla/mux"
)

const (
	routePrefix = "/fibonacci"
)

// Router - router struct
type Router struct {
	*mux.Router
}

// NewRouter - new router instance
func NewRouter() *Router {
	return &Router{mux.NewRouter()}
}

// InitializeRoutes ...
func (r *Router) InitializeRoutes() {
	api := (*r).PathPrefix(routePrefix).Subrouter()

	api.HandleFunc("/current", panicRecover(handler.GetCurrentFibSequenceHandler)).
		Methods(http.MethodGet)

	api.HandleFunc("/next", panicRecover(handler.GetNextFibSequenceHandler)).
		Methods(http.MethodGet)

	api.HandleFunc("/previous", panicRecover(handler.GetPreviousFibSequenceHandler)).
		Methods(http.MethodGet)

	api.HandleFunc("/reset", panicRecover(handler.ResetFibSequenceHandler)).
		Methods(http.MethodGet)

	api.HandleFunc("/{term}", panicRecover(handler.GetNthTermOfSequenceHandler)).
		Methods(http.MethodGet)
}

// panicRecover - recover endpoint from panic
func panicRecover(restart func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			err := recover()
			if err != nil {
				handler.ErrorLogger.Println(err)
				jsonBody, _ := json.Marshal(map[string]string{
					"error": "There was an internal server error",
				})
				w.WriteHeader(http.StatusInternalServerError)
				w.Write(jsonBody)
			}
		}()
		restart(w, r)
	}
}
