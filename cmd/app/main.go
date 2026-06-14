package main

import (
	"tanneryApp/internal/app"
)

// @title Tannery App
// @version 1.0
// @description Tannery app API for lightweight web messenger made for few friends
func main() {
	err := app.Run("configs")
	if err != nil {
		panic(err)
	}
}
