package main

import (
	"context"
	"fmt"
	"my-api/sdk"
)

func main() {
	ctx := context.Background()

	client, _ := sdk.NewClientWithResponses("http://localhost:8888")

	greeting, err := client.GetGreetingWithResponse(ctx, "world")
	if err != nil {
		panic(err)
	}

	if greeting.StatusCode() > 200 {
		panic(greeting.ApplicationproblemJSONDefault)
	}

	fmt.Println(greeting.JSON200.Message)
}
