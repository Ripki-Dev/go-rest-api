package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type CategoryController interface {
	Create(writter http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindAll(writter http.ResponseWriter, request *http.Request params httprouter.Params)
}
