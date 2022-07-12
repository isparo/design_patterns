package main

import "fmt"

type singletonCon struct {
	url  string
	port string
}

var singletonConInstance *singletonCon

func GetInstance() *singletonCon {
	if singletonConInstance == nil {

		fmt.Println("Creating singleton instance")

		singletonConInstance = &singletonCon{
			url:  "the url",
			port: "the port",
		}
	} else {
		fmt.Println("Singleton instance already created")
	}

	return singletonConInstance
}

func main() {
	for i := 0; i < 10; i++ {
		GetInstance()
	}
}
