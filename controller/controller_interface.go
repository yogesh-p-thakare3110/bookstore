package controller

import "net/http"

type IController interface {
	AllBook() func(w http.ResponseWriter, r *http.Request)
	InsertBook() func(w http.ResponseWriter, r *http.Request)
	DeleteBook() func(w http.ResponseWriter, r *http.Request)
	UpdateBook() func(w http.ResponseWriter, r *http.Request)
}
