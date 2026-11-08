package main

import "fmt"

// Что будет выведено? Почему? Подумать, после запустить программу. Исправить.

type Person struct {
	Name string
}

func changeName(person *Person) {

	*person = Person{Name: "Alice"}

}

func main() {
	person := &Person{
		Name: "Bob",
	}
	fmt.Println(person.Name)
	changeName(person)
	fmt.Println(person.Name)
}
