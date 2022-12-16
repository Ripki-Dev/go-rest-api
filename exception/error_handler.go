package exception

import (
	"net/http"

	"restGo/helper"
	"restGo/model/web"

	"github.com/go-playground/validator/v10"
)

func ErrorHandler(writter http.ResponseWriter, request *http.Request, err interface{}) {
	if notFoundError(writter, request, err) {
		return
	}

}

func notFoundError(writter http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		writter.Header().Set("Content-Type", "application/json")
		writter.WriteHeader(http.StatusNotFound)

		webReasponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   exception.Error,
		}

		helper.WriteToResponseBody(writter, webReasponse)
		return true
	} else {
		return false
	}
}

func validationErrors(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		webReasponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   exception.Error(),
		}
		helper.WriteToResponseBody(writer, webReasponse)
		return true
	} else {
		return false
	}
}

func internalServerError(writter http.ResponseWriter, request *http.Request, err interface{}) {
	writter.Header().Set("Content-Type", "application/json")
	writter.WriteHeader(http.StatusInternalServerError)

	webReasponse := web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Data:   err,
	}

	helper.WriteToResponseBody(writter, webReasponse)
}
