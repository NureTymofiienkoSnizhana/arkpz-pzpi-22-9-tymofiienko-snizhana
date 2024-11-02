// Поганий приклад
//package my_Package
//package MyPackage
//package myPackage
//Гарний приклад
package main

import "fmt"

func main() {
	// Гарний приклад
	var interfaceValue interface{}

	//Умовний оператор switch type
	switch v := interfaceValue.(type) {
	case int:
		fmt.Printf("Ціле число %d\n", v)
	case string:
		fmt.Printf("Рядок %s\n", v)
	default:
		fmt.Printf("Невідомий тип %T\n", v)
	}

	//Comment
	/* Comment Comment
	Comment */

	//Оголошення змінних
	var student1 string = "John"
	var student2 = "Jane"
	x := 2

	//Поганий приклад: використано підкреслення
	var max_value int = 4
	// Гарний приклад
	var maxValue int = 4

	// Поганий приклад: краще використати однолітерне ім'я
	func(string string) {}()
	// Гарний приклад
	func(s string) {}()

	// Поганий приклад: не однолітерне ім'я
	for index := 0; index < 10; index++ {
	}
	// Гарний приклад
	for i := 0; i < 10; i++ {
	}

	// Поганий приклад: всі великі літери, підкреслення
	const MAX_PACKET_SIZE = 512
	const kMaxBufferSize = 1024
	// Гарний приклад
	const MaxPacketSize = 512

	// Поганий приклад:
	func ParseUrlData(urlString string) {}
	// Гарний приклад
	func ParseURLData(URLString string) {}


	// Поганий приклад: Функція з префіксом Get.
	func (u *User) GetName() string {
		return u.Name
	}
	// Гарний приклад
	func (u *User) Name() string {
		return u.Name
	}

	// Поганий приклад: Ім'я змінної повторює ім'я типу.
	var nameString string
	// Гарний приклад
	var name string

	// Умовний оператор if
	if x > 0 {
		return y
	}

	// Умовний оператор switch
	switch {
	case x < 0:
		fmt.Println("Negative")
	case x == 0:
		fmt.Println("Zero")
	default:
		fmt.Println("Positive")
	}

	// Цикл for

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	Цикл for з використанням range

	for key, value := range myMap {
		fmt.Println(key, value)
	}

	// Виклик функції
	myMessage()

	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(x)
}

// Функція
func myMessage() {
	fmt.Println("Hello world!")
}
