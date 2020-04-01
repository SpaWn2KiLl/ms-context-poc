package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	printContextStatus(r.Context())

	time.Sleep(5 * time.Second)

	printContextStatus(r.Context())
}

func printContextStatus(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("Context closed")
		fmt.Println(ctx.Err())
	default:
		fmt.Println("Context is still valid")
	}
}
