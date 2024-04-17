package main

import (
	"fmt"
	"url-shortener-rest-practice/internal/config"
)

func main() {
	// TODO: init config: cleanenv
	cfg := config.MustLoad()

	fmt.Println(cfg)

	// TODO: init logger: slog

	// TODO: init storage: sqlite

	// TODO: init router: chi, "chi render"

	// TODO: run app

	// TODO: graceful shutdown
}
