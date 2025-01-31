package main

import (
	"bufio"
	"fmt"
	"net/http"

	_ "github.com/apache/skywalking-go"
)

func main() {
	for {
		resp, err := http.Get("http://127.0.0.1:19091/owners/1882131")
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		fmt.Println("Response status:", resp.Status)

		scanner := bufio.NewScanner(resp.Body)
		for i := 0; scanner.Scan() && i < 5; i++ {
			fmt.Println(scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			panic(err)
		}
	}
}
