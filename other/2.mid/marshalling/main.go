package main

import (
	"encoding/json"
	"fmt"
)

// Что будет выведено и почему? Подумать, после запустить программу. Исправить.
// (вспомнить особенность работы маршалинга из стандартного пакета).

type Data struct {
	Age  int    `json:"age"`
	name string `json:"name"` //написанно c маденькой буквой
}

func testData() {
	in := Data{33, "Ivan"}
	fmt.Printf("%#v\n", in)

	encoded, _ := json.Marshal(in)
	fmt.Println(string(encoded))

	var out Data
	json.Unmarshal(encoded, &out)
	fmt.Printf("%#v\n", out)
}

func main() {
	testData()
}
