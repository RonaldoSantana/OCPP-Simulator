package main

import (
	"fmt"
	"github.com/RonaldoSantana/ocpp-simulator/charger"
	//"os"
	"os"
)

func main() {
	method := "Authorize"
	args := []string{"Authorize"}

	if false {
		simulator := charger.Simulator{}
		response := simulator.Call(method, args) //os.Args...)

		fmt.Println(response)
	}
}
