package main

import (
	"fmt"
	"weather-aggregation-service/models"
)

func main() {
	m := &models.ExampleModel{Name: "test"}

	fmt.Println(m)
}
