package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "token", "abc123")
	bookHotel(ctx, "habbo hotel")
}

func bookHotel(ctx context.Context, name string) { // convenção sempre usar o ctx como primeiro parâmetro
	token := ctx.Value("token")
	fmt.Println(token, name)
}
