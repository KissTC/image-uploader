package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()

	// set a value
	ctx = context.WithValue(ctx, "favorite-color", "blue")
	value := ctx.Value("favorite-color")

	fmt.Println(value)
}
