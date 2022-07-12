package main

import "fmt"

// code related with computers
type computer interface {
	print()
	setPrinter(printer)
}

// desktop computer, struct definition and interface methods implementation
type desktop struct {
	printer printer
}

func (d *desktop) print() {
	fmt.Println("Print document for desktop pc")
	d.printer.printFile()
}

func (d *desktop) setPrinter(p printer) {
	d.printer = p
}

// laptop computer, struct definition and interface methods implementation
type laptop struct {
	printer printer
}

func (l *laptop) print() {
	fmt.Println("Print document for laptop")
	l.printer.printFile()
}

func (l *laptop) setPrinter(p printer) {
	l.printer = p
}

// code related wit printers
type printer interface {
	printFile()
}

type printerOne struct {
}

func (p printerOne) printFile() {
	fmt.Println("Printing by a Printer one")
}

type printerTwo struct {
}

func (p printerTwo) printFile() {
	fmt.Println("Printing by a Printer Two")
}

func main() {

	p1 := printerOne{}
	p2 := printerTwo{}

	desktopPc := desktop{}

	desktopPc.setPrinter(p1)
	desktopPc.print()
	fmt.Println()
	desktopPc.setPrinter(p2)
	desktopPc.print()
	fmt.Println()

	laptopPc := laptop{}
	laptopPc.setPrinter(p1)
	laptopPc.print()
	fmt.Println()
	laptopPc.setPrinter(p2)
	laptopPc.print()
	fmt.Println()
}
