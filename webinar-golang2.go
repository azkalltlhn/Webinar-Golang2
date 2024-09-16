package main

import (
	"fmt"
)

// Struct Person
type Person struct {
	Name string
	Age  int
}

// Fungsi CreatePerson
func CreatePerson(name string, age int) *Person {
	if age < 0 {
		panic("Invalid age")
	}
	return &Person{
		Name: name,
		Age:  age,
	}
}

// Fungsi PrintPerson
func PrintPerson(p *Person) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in PrintPerson:", r)
		}
	}()
	fmt.Printf("Person: Name = %s, Age = %d\n", p.Name, p.Age)
}

// Fungsi HandlePerson
func HandlePerson(name string, age int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error:", r)
		}
	}()
	person := CreatePerson(name, age)
	PrintPerson(person)
}

func main() {
	// Contoh kasus, menangani panic
	HandlePerson("Caca", -21)  // Ini akan memicu panic karena umur negatif
	HandlePerson("Azka", 21)  // Ini akan berhasil
}
