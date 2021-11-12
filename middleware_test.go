package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

type LoginMiddleware struct{
	http.Handler
}

func (middleware LoginMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Recive Request")
	middleware.Handler.ServeHTTP(writer, request)
}

func TestMiddleware(t *testing.T) {
	router := httprouter.New()

	router.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params){
		fmt.Fprint(writer, "Middleware")
	})

	middleware := LoginMiddleware{router}

	request := httptest.NewRequest("GET", "http://localhost:3000/", nil)
	recorder := httptest.NewRecorder()

	middleware.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ :=io.ReadAll(response.Body)

	assert.Equal(t, "Middleware", string(body))
}