package main

import "fmt"

type building interface {
	getType() string
	accept(visitor)
}

//
type visitor interface {
	visitForChurch(*church)
	visitForHotel(*hotel)
	visitForStadium(*stadium)
}

//
type church struct {
	sideA int
	sideB int
}

func (c *church) accept(v visitor) {
	v.visitForChurch(c)
}

func (s *church) getType() string {
	return "Church"
}

//
type stadium struct {
	radius int
}

func (s *stadium) accept(v visitor) {
	v.visitForStadium(s)
}

func (s *stadium) getType() string {
	return "Stadium"
}

//
type hotel struct {
	l int
	b int
}

func (h *hotel) accept(v visitor) {
	v.visitForHotel(h)
}

func (t *hotel) getType() string {
	return "Hotel"
}

//
type areaCalculator struct {
	area int
}

func (a *areaCalculator) visitForChurch(c *church) {
	fmt.Println("Calculating area for church")
}

func (a *areaCalculator) visitForStadium(s *stadium) {
	fmt.Println("Calculating area for stadium")
}
func (a *areaCalculator) visitForHotel(h *hotel) {
	fmt.Println("Calculating area for hotel")
}

//
type middleCoordinates struct {
	x int
	y int
}

func (a *middleCoordinates) visitForChurch(c *church) {
	fmt.Println("Calculating middle point coordinates for church")
}

func (a *middleCoordinates) visitForStadium(s *stadium) {
	fmt.Println("Calculating middle point coordinates for stadium")
}
func (a *middleCoordinates) visitForHotel(h *hotel) {
	fmt.Println("Calculating middle point coordinates for hotel")
}

//
func main() {
	church := &church{sideA: 2, sideB: 4}
	stadium := &stadium{radius: 3}
	hotel := &hotel{l: 2, b: 3}

	areaCalculator := &areaCalculator{}

	church.accept(areaCalculator)
	stadium.accept(areaCalculator)
	hotel.accept(areaCalculator)

	fmt.Println()
	middleCoordinates := &middleCoordinates{}
	church.accept(middleCoordinates)
	stadium.accept(middleCoordinates)
	hotel.accept(middleCoordinates)
}
