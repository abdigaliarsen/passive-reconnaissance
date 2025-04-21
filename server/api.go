package server

import (
	"net/http"
	"net/http/httptest"
)

type API interface {
	http.Handler
	Exec(r *http.Request) (*http.Response, error)
	AddHandler(path, method string, functor http.HandlerFunc) API
}

type api struct {
	mux         *http.ServeMux
	descriptors map[string]map[string]http.HandlerFunc
}

func NewApi() API {
	return &api{
		mux:         http.NewServeMux(),
		descriptors: make(map[string]map[string]http.HandlerFunc),
	}
}

func (a *api) AddHandler(path, method string, functor http.HandlerFunc) API {
	if _, exists := a.descriptors[path]; !exists {
		a.descriptors[path] = make(map[string]http.HandlerFunc)
	}
	a.descriptors[path][method] = functor

	a.mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if handler, ok := a.descriptors[r.URL.Path][r.Method]; ok {
			handler(w, r)
		} else if _, pathExists := a.descriptors[r.URL.Path]; pathExists {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		} else {
			http.NotFound(w, r)
		}
	})

	return a
}

func (a *api) Exec(r *http.Request) (*http.Response, error) {
	handler, ok := a.descriptors[r.URL.Path][r.Method]
	if !ok {
		return &http.Response{
			StatusCode: http.StatusNotFound,
			Body:       http.NoBody,
		}, nil
	}

	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, r)
	return rec.Result(), nil
}

func (a *api) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.mux.ServeHTTP(w, r)
}
