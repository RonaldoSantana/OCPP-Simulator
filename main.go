package main

import (
	"fmt"
	"github.com/RonaldoSantana/ocpp-simulator/charger"
	//"os"
	"os"
)



func main() {

	method := os.Args[1]
	simulator := charger.Simulator{}
	response := simulator.Call(method, os.Args[2:]...) //os.Args...)

	fmt.Println(response)
}
