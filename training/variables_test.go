package training_test

import (
	"testing"

	"github.com/FranciscoAguiar/gotraining/training"
)

func TestMascot(t *testing.T) {
	if training.BestMascot() != "sutao" {
		t.Error("Best mascot is not Sutao")
	}
}
