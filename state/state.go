package main

import (
	"errors"
	"fmt"
)

// state interface
type state interface {
	requestItem() error
	dispenseItem() error
}

// select item state
type selectItemMachineState struct {
	vendingMachine *vendingMachine
}

func (si *selectItemMachineState) requestItem() error {
	fmt.Println("Preparing selected item")
	si.vendingMachine.currentState = si.vendingMachine.itemRequested
	return nil
}

func (si *selectItemMachineState) dispenseItem() error {
	return errors.New("first select an item")
}

// dispense item state
type dispenseItemMachineState struct {
	vendingMachine *vendingMachine
}

func (di *dispenseItemMachineState) requestItem() error {
	return errors.New("dispensing item, please wait")
}

func (di *dispenseItemMachineState) dispenseItem() error {
	fmt.Println("Item dispensed")
	di.vendingMachine.currentState = di.vendingMachine.selectionAvailable
	return nil
}

//vendingMachine struct, handle all the posible states
type vendingMachine struct {
	selectionAvailable state
	itemRequested      state

	currentState state
}

func (v *vendingMachine) requestItem() error {
	return v.currentState.requestItem()
}

func (v *vendingMachine) dispenseItem() error {
	return v.currentState.dispenseItem()
}

func newVendingMachine() *vendingMachine {

	v := &vendingMachine{}

	selectItem := &selectItemMachineState{
		vendingMachine: v,
	}

	dispenseItemState := &dispenseItemMachineState{
		vendingMachine: v,
	}

	v.selectionAvailable = selectItem
	v.itemRequested = dispenseItemState
	v.currentState = selectItem

	return v
}

func main() {
	vm := newVendingMachine()

	fmt.Println("Please select an item")

	err := vm.requestItem()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("---")

	err = vm.dispenseItem()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("---")

	fmt.Println("Please select an item")

	err = vm.requestItem()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("---")
	err = vm.dispenseItem()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("---")

}
