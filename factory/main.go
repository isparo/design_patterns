package main

import (
	"errors"
	"fmt"
)

const (
	teacherType string = "teacherRobot"
	fighterType string = "fighterRobot"
)

type imRobot interface {
	whoAmI() string
}

// base struct
type robot struct {
	name     string
	rType    string
	autonomy float32
}

func (r robot) whoAmI() string {
	return fmt.Sprintf("Name: %s\n"+
		"rType: %s"+
		"Autonomy; %2.2f",
		r.name,
		r.rType,
		r.autonomy)
}

// Teacher Robot
type teacherRobot struct {
	robot
}

func newTeacherRobot() imRobot {
	return teacherRobot{
		robot: robot{
			name:     "Mr. Teacher",
			rType:    "teacher",
			autonomy: 100.00,
		},
	}
}

// Fighter Robot
type fighterRobot struct {
	robot
}

func newFighterRobot() imRobot {
	return fighterRobot{
		robot: robot{
			name:     "Mr. Fighter",
			rType:    "fighter",
			autonomy: 70.00,
		},
	}
}

// factory function
func createRobot(rType string) (imRobot, error) {

	if rType == teacherType {
		return newTeacherRobot(), nil
	}

	if rType == fighterType {
		return newFighterRobot(), nil
	}

	return nil, errors.New(fmt.Sprintf("could not create a robot with: %s", rType))
}

func main() {
	robotA, err := createRobot(teacherType)
	robotB, err := createRobot(fighterType)
	wrongRobot, err := createRobot("invalid")

	robotA.whoAmI()
	fmt.Println(robotA.whoAmI())
	fmt.Println("---------------------")
	fmt.Println(robotB.whoAmI())

	fmt.Println("---------------------")
	fmt.Println("wrongRobot : ", wrongRobot)
	fmt.Println("wrongRobot err : ", err.Error())

}
