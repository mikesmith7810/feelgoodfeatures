package main

import "feelgoodfeatures/internal/api"

func main() {
	r := api.CreateRouter()

	err := r.Run(":8080")
	if err != nil {
		return
	}
}
