package main

import (
	"context"
	"fmt"
)

func main() {

	ctx := context.Background()

	newCtx := addValue(ctx)

	s := newCtx.Value("Bootcamp")

	fmt.Println(s)

}

func addValue(ctx context.Context) context.Context {
	return context.WithValue(ctx, "Bootcamp", "Go")
}
