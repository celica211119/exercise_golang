package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/danielgtaylor/huma/v2/humacli"
	"github.com/go-chi/chi/v5"
)

type Options struct {
	Port int `help:"Port to listen on" short:"p" default:"8888"`
}

type GreetingOutPut struct {
	Body struct {
		Message string `json:"message" example:"Hello, world!" doc:"Greeting message"`
	}
}

type ReviewInput struct {
	Body struct {
		Author  string `json:"author" maxLength:"10" doc:"Author of the review"`
		Rating  int    `json:"rating" minimum:"1" maximum:"5" doc:"Rating from 1 to 5"`
		Message string `json:"message,omitempty" maxLength:"100" doc:"Review messge"`
	}
}

// https://huma.rocks/tutorial/your-first-api/
func Tutorial1() {
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

// https://huma.rocks/tutorial/service-configuration/
func Tutorial2() {
	cli := humacli.New(func(hooks humacli.Hooks, options *Options) {
		router := chi.NewMux()
		api := humachi.New(router, huma.DefaultConfig("My API", "1.0.0"))

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

		hooks.OnStart(func() {
			fmt.Printf("Starting server on Port %d...\n", options.Port)
			http.ListenAndServe(fmt.Sprintf(":%d", options.Port), router)
		})
	})

	cli.Run()
}

// https://huma.rocks/tutorial/sending-data/
func addRoutes(api huma.API) {
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

	huma.Register(api, huma.Operation{
		OperationID:   "post-review",
		Method:        http.MethodPost,
		Path:          "/reviews",
		Summary:       "Post a review",
		Tags:          []string{"Reviews"},
		DefaultStatus: http.StatusCreated,
	}, func(ctx context.Context, input *ReviewInput) (*struct{}, error) {
		fmt.Println(input)
		return nil, nil
	})
}

func Tutorial3() {
	cli := humacli.New(func(hooks humacli.Hooks, options *Options) {
		router := chi.NewMux()
		api := humachi.New(router, huma.DefaultConfig("My API", "1.0.0"))

		addRoutes(api)

		hooks.OnStart(func() {
			fmt.Printf("Starting server on Port %d...\n", options.Port)
			http.ListenAndServe(fmt.Sprintf(":%d", options.Port), router)
		})
	})

	cli.Run()
}

// https://huma.rocks/tutorial/writing-tests/
func Tutorial4() {
	cli := humacli.New(func(hooks humacli.Hooks, options *Options) {
		router := chi.NewMux()
		api := humachi.New(router, huma.DefaultConfig("My API", "1.0.0"))

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

		huma.Register(api, huma.Operation{
			OperationID:   "post-review",
			Method:        http.MethodPost,
			Path:          "/reviews",
			Summary:       "Post a review",
			Tags:          []string{"Reviews"},
			DefaultStatus: http.StatusCreated,
		}, func(ctx context.Context, input *ReviewInput) (*struct{}, error) {
			fmt.Println(input)
			return nil, nil
		})

		hooks.OnStart(func() {
			fmt.Printf("Starting server on Port %d...\n", options.Port)
			http.ListenAndServe(fmt.Sprintf(":%d", options.Port), router)
		})
	})

	cli.Run()
}

func main() {
	// Tutorial1()
	// Tutorial2()
	// Tutorial3()
	Tutorial4()
}
