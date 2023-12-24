// Реализовать паттерн адаптер на любом примере

package main

import "fmt"

// OldPrinter - старый компонент с несовместимым интерфейсом
type OldPrinter struct {
}

func (op *OldPrinter) Print(str string) {
	fmt.Println(str)
}

// Printer - новый интерфейс, ожидаемый клиентом
type Printer interface {
	PrintMessage(message string)
}

// PrinterAdapter - адаптер для использования OldPrinter в новом интерфейсе
type PrinterAdapter struct {
	OldPrinter *OldPrinter
}

func (pa *PrinterAdapter) PrintMessage(message string) {
	pa.OldPrinter.Print(message)
}

func main() {
	// Используем новый интерфейс
	printer := &PrinterAdapter{OldPrinter: &OldPrinter{}}
	printer.PrintMessage("Hello, Adapter Pattern!")
}
