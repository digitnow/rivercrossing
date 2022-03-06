package main

import (
	"fmt"
	"github.com/digitnow/rivercrossing/state"
)

func main() {
	fmt.Println(state.ViewState())

	state.EnterBoat()
	fmt.Println(state.ViewState())

	state.Rivercross()
	fmt.Println(state.ViewState())
}
