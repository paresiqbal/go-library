// package main

// import (
// 	"errors"
// 	"fmt"
// )

// var (
// 	ValidationError = errors.New("validation error")
// 	NotFound        = errors.New("not found")
// )

// func getById(id string) error {
// 	if id == "" {
// 		return ValidationError
// 	}

// 	if id != "Pares" {
// 		return NotFound
// 	}

// 	return nil
// }

// func main() {
// 	err := getById("Pares")
// 	if err != nil {
// 		if errors.Is(err, ValidationError) {
// 			fmt.Println("Validation Error")
// 		} else if errors.Is(err, NotFound) {
// 			fmt.Println("Not Found")
// 		} else {
// 			fmt.Println("Unknown Error")
// 		}
// 	}
// }
