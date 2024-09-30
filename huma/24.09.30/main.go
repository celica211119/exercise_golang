package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/go-chi/chi/v5"
)

type GreetingOutPut struct {
	Body struct {
		Message string `json:"message" example:"Hello, world!" doc:"Greeting message"`
	}
}

func main() {
	router := chi.NewMux()
	api := humachi.New(router, huma.DefaultConfig("My API", "1.0.0"))

	huma.Get(api, "/greeting_old01/{name}", func(ctx context.Context, input *struct {
		Name string `path:"name" maxLength:"30" example:"world" doc:"Name to greet"`
	}) (*GreetingOutPut, error) {
		resp := &GreetingOutPut{}
		resp.Body.Message = fmt.Sprintf("Hello, %s!", input.Name)
		return resp, nil
	})

	huma.Register(api, huma.Operation{
		OperationID: "get_greeting",
		Method:      http.MethodGet,
		Path:        "/greeting/{name}",
		Summary:     "Get a greeting",
		Description: "Get a greeting for a person by name",
		Tags:        []string{"Greetings"},
	}, func(ctx context.Context, input *struct {
		Name string `path:"name" maxLength:"30" minLength:"3" example:"world" doc:"Name to greet"`
	}) (*GreetingOutPut, error) {
		resp := &GreetingOutPut{}
		resp.Body.Message = fmt.Sprintf("Hello, %s!", input.Name)
		return resp, nil
	})

	http.ListenAndServe("127.0.0.1:8888", router)
}
