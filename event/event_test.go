// Package for testing events in the riverworld
// Possible events: put(item), getin(), getout(), crossriver(), takeout(item)
// start() or setup(), reset() osv.
package event

import "testing"

func TestPut(t *testing.T) {
    // Hva forventer jeg?
    wanted := "[kylling rev korn ---\\ \\_korn_/ _________________/---]"
    got := Put("korn") // Hva fikk jeg?
    if got != wanted {
             t.Errorf("Feil, fikk %q, ønsket %q.", got, wanted)
    }
}