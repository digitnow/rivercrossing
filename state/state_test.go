package state

import "testing"

func TestViewState(t *testing.T) {
    wanted := "[kylling rev korn mann ---\\ \\__/ _________________/---]"
    state := ViewState();
    if state != wanted {
         t.Errorf("Feil, fikk %q, ønsket %q.", state, wanted)
    }
}
