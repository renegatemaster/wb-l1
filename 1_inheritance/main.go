package main

import "fmt"

// Структура Human, его методы и конструктор

type Human struct {
	Name string
	Age  int
}

func (h *Human) GetName() {
	fmt.Printf("Human name: %s\n", h.Name)
}

func (h *Human) GetAge() {
	fmt.Printf("%s age is %d\n", h.Name, h.Age)
}

func (h *Human) CallGetName() {
	h.GetName()
}

func NewHuman(name string, age int) *Human {
	return &Human{
		Name: name,
		Age:  age,
	}
}

// Структура Action, встраивающая и переопределящая методы от структуры Human

type Action struct {
	Human
}

// Переопределённый метод
func (a *Action) GetName() {
	fmt.Println("Action name")
}

// Конструктор Action
func NewAction(h Human) *Action {
	return &Action{
		Human: h,
	}
}

func main() {
	person := NewHuman("Ivan", 24)
	action := NewAction(*person)

	// Методы встраеваемой структуры наследуются встраивающей
	// Можно обращатся к методам самой встроенной структуры
	person.GetAge()       // Ivan age is 24
	action.GetAge()       // Ivan age is 24
	action.Human.GetAge() // Ivan age is 24

	fmt.Println()

	// У встраивающей структуры можно переопределить метод
	// встраеваемой структуры, при этом остаётся возможность
	// использовать метод встраеваемой структуры
	person.GetName()       // Human name: Ivan
	action.GetName()       // Action name
	action.Human.GetName() // Human name: Ivan

	fmt.Println()

	// При вызове метода, вызывающего другой метод,
	// который переопределён во встраивающей структуре,
	// будет вызван метод встраиваемой структуры
	person.CallGetName()
	action.CallGetName()
	action.Human.CallGetName()
}
