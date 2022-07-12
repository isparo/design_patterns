package main

import "fmt"

type robot struct {
	rType     string
	autonomy  float32
	hasGuitar bool
	hasGun    bool
	hasHelm   bool
	hasHat    bool
	hasSword  bool
}

type robotBuilder struct {
	robot robot
}

func newRobotBuilder() *robotBuilder {
	return &robotBuilder{
		robot: robot{},
	}
}

func (r *robotBuilder) setType(rType string) *robotBuilder {
	r.robot.rType = rType
	return r
}

func (r *robotBuilder) setAutonomy(autonomy float32) *robotBuilder {
	r.robot.autonomy = autonomy
	return r
}

func (r *robotBuilder) withGuitar() *robotBuilder {
	r.robot.hasGuitar = true
	return r
}

func (r *robotBuilder) withGun() *robotBuilder {
	r.robot.hasGun = true
	return r
}

func (r *robotBuilder) withHelm() *robotBuilder {
	r.robot.hasHelm = true
	return r
}

func (r *robotBuilder) withHat() *robotBuilder {
	r.robot.hasHat = true
	return r
}

func (r *robotBuilder) withSword() *robotBuilder {
	r.robot.hasSword = true
	return r
}

func (r *robotBuilder) build() robot {
	return r.robot
}

func main() {
	robotFighter := newRobotBuilder().
		setAutonomy(100.00).
		setType("fighter").
		withGun().
		withHelm().
		withSword()

	robotSinger := newRobotBuilder().
		setAutonomy(100.00).
		setType("singer").
		withHat().
		withGuitar()

	printRobotInfo(robotFighter.build())
	fmt.Println("______________________")
	printRobotInfo(robotSinger.build())
}

func printRobotInfo(robot robot) {

	robotInfMsg := fmt.Sprintf("Type: %s \n"+
		"Autonomy: %f \n"+
		"hasGuitar: %t \n"+
		"hasGun: %t \n"+
		"hasHelm: %t \n"+
		"hasHat: %t \n"+
		"hasSword: %t \n",
		robot.rType,
		robot.autonomy,
		robot.hasGuitar,
		robot.hasGun,
		robot.hasHelm,
		robot.hasHat,
		robot.hasSword)

	fmt.Println(robotInfMsg)
}
