package main

import (
	"encoding/json"
	"fmt"
)

type MyJson struct {
	Cat `json:"cat"`
}

type Cat struct {
	Name string `json:"name"`
	Age  uint   `json:"age"`
}

type Product struct {
	ID       uint64   `json:"id"`
	Name     string   `json:"name"`
	SKU      string   `json:"sku"`
	Category Category `json:"category"`
}

type Category struct {
	ID          uint64 `json:"id"`
	Name        string `json:"name,omitempty"`
	Description string `json:"-"`
}

func main() {

	p := Product{ID: 42, Name: "Tea Pot", SKU: "TP12", Category: Category{ID: 2}}

	b, err := json.MarshalIndent(p, "", "  ")

	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))

	myJson := []byte(`
	{
		"cat": {
			"name": "Joey",
			"age": 8
		}
	}`)

	c := MyJson{}
	err = json.Unmarshal(myJson, &c)

	if err != nil {
		panic(err)
	}

	fmt.Println(c.Cat.Name)
	fmt.Println(c.Cat.Age)
}
