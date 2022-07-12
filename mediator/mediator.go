package main

import "fmt"

type airplane interface {
	arrive()
	depart()
	permitArrival()
}

type mediator interface {
	canArrive(airplane) bool
	notifyAboutDeparture()
}

// component one, passengers
type passenger struct {
	mediator mediator
}

func (p *passenger) arrive() {
	if !p.mediator.canArrive(p) {
		fmt.Println("airplane Passenger: Arrival blocked, waiting")
		return
	}
	fmt.Println("airplane Passenger: Arrived")
}

func (p *passenger) depart() {
	fmt.Println("airplane Passenger: Leaving")
	p.mediator.notifyAboutDeparture()
}

func (p *passenger) permitArrival() {
	fmt.Println("airplane Passenger: landing permitted, arriving")
	p.arrive()
}

// component two, the airplane
type boeing737 struct {
	mediator mediator
}

func (b *boeing737) arrive() {
	if !b.mediator.canArrive(b) {
		fmt.Println("boeing 737: Landing blocked, waiting")
		return
	}
	fmt.Println("boeing 737: Arrived")
}

func (b *boeing737) depart() {
	fmt.Println("boeing 737: Taking off")
	b.mediator.notifyAboutDeparture()
}

func (g *boeing737) permitArrival() {
	fmt.Println("boeing 737: Arrival permitted")
	g.arrive()
}

// mediator, control tower
type airportManager struct {
	isRunwayFree bool
	airportQueue []airplane
}

func newStationManger() *airportManager {
	return &airportManager{
		isRunwayFree: true,
	}
}

func (s *airportManager) canArrive(t airplane) bool {
	if s.isRunwayFree {
		s.isRunwayFree = false
		return true
	}
	s.airportQueue = append(s.airportQueue, t)
	return false
}

func (s *airportManager) notifyAboutDeparture() {
	if !s.isRunwayFree {
		s.isRunwayFree = true
	}
	if len(s.airportQueue) > 0 {
		firstAirplaneInQueue := s.airportQueue[0]
		s.airportQueue = s.airportQueue[1:]
		firstAirplaneInQueue.permitArrival()
	}
}

//
func main() {
	stationManager := newStationManger()

	airplanePassenger := &passenger{
		mediator: stationManager,
	}
	boeing := &boeing737{
		mediator: stationManager,
	}

	airplanePassenger.arrive()
	boeing.arrive()
	airplanePassenger.depart()
}
