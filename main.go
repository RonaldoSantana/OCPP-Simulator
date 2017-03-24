package main

import (
	"fmt"
	"github.com/RonaldoSantana/ocpp-simulator/charger"
	//"os"
)

func main() {
	method := "Authorize"
	args := []string{
		"veefil-21159",
		"B4F62CEF",
	}

	simulator := charger.Simulator{}
	response := simulator.Call(method, args...) //os.Args...)

	fmt.Println(response)

}
