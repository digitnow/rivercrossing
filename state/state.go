package state

var state = "[hund katt hamster HS ---\\ \\--/---------/---]"

func EnterBoat() {
	state = "[hund katt hamster ---\\ \\_HS_/----------/---]"
}

func Rivercross() {
	state = "[hun katt hamster ---\\ \\----------\\-HS-/---]"
}

func ViewState() string {
	return state
}
