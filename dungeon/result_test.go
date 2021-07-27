package dungeon_test

import (
	"jp/thelow/static/dungeon"
	"log"
	"testing"
)

func TestResultIncrement(t *testing.T) {

	result := dungeon.NewResult()

	result.Increment("A")
	result.Increment("A")
	result.Increment("A")

	result.Increment("B")
	result.Increment("B")

	result.Increment("C")
	result.Increment("C")
	result.Increment("C")
	result.Increment("C")

	if result.GetCount("A") != 3 {
		log.Fatalf("Count of A is not 3")
	}

	if result.GetCount("B") != 2 {
		log.Fatalf("Count of B is not 2")
	}

	if result.GetCount("C") != 4 {
		log.Fatalf("Count of C is not 4")
	}

	if result.GetTotal() != 9 {
		log.Fatalf("Total is not 9")
	}
}
