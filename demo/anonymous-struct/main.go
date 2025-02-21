package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// declare the 'car' struct type
type car struct {
	make    string
	model   string
	mileage int
}

// An anonymous struct is just like a normal struct, but it is defined without a name and
// therefore cannot be referenced elsewhere in the code
// Ref: https://blog.boot.dev/golang/anonymous-structs-golang/
func main() {
	// create an instance of a car
	newCar := car{
		make:    "Ford",
		model:   "taurus",
		mileage: 200000,
	}
	fmt.Println(newCar)

	// create an anonymous struct
	newCar1 := struct {
		make    string
		model   string
		mileage int
	}{
		make:    "Ford",
		model:   "Taurus",
		mileage: 200000,
	}
	fmt.Println(newCar1)
}

// Can use anonymous structs to marshal and unmarshal JSON data in HTTP handlers
func createCarHandler1(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	decoder := json.NewDecoder(req.Body)
	newCar := struct {
		Make    string `json:"make"`
		Model   string `json:"model"`
		Mileage int    `json:"mileage"`
	}{}
	err := decoder.Decode(&newCar)
	if err != nil {
		log.Println(err)
		return
	}
	// makeCar(newCar.Make, newCar.Model, newCar.Mileage)
}

// Don’t use map[string]interface{} for JSON data if you can avoid it
func createCarHandler2(w http.ResponseWriter, req *http.Request) {
	myMap := map[string]interface{}{} // myMap := map[string]any{} --> similar to Map<String, Object> in Java
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&myMap)
	if err != nil {
		log.Println(err)
		return
	}
	model, ok := myMap["model"]
	if !ok {
		fmt.Println("field doesn't exist")
		return
	}
	modelString, ok := model.(string)
	if !ok {
		fmt.Println("model is not a string")
	}
	// do something with model field
	fmt.Println(modelString)
}
