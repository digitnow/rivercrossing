package state

import (
    "testing"
    
)

func TestCheckState(t *testing.T) {
    wanted :=["hund katt hamster HS ---\\ \\--/---------/---"]
    state := ViewState()
    if state != wanted {
t.Errorf( format: "Feil, fikk %q, ønsket %q.", state, wanted)
    }
}