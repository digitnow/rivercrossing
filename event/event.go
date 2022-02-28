package main

import (
	"fmt"
	"github.com/digitnow/rivercrossing/state"
)

func main() {
	fmt.Println(state.ViewState())

	state.Enterboat()
	fmt.Println(state.ViewState())

	state.Rivercross()
	fmt.Println(state.ViewState())
}
