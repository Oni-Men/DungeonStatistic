package dungeon_test

import (
	"encoding/json"
	"io/ioutil"
	"jp/thelow/static/dungeon"
	"jp/thelow/static/model"
	"log"
	"testing"
)

func TestFetchCompletions(t *testing.T) {
	read, err := ioutil.ReadFile("../data/test/dungeon.json")

	if err != nil {
		t.Fatal(err)
	}

	traces := new(model.Traces)
	if err := json.Unmarshal(read, traces); err != nil {
		log.Fatal(err)
	}

	result := dungeon.NewResult(2021, 7)

	for _, t := range traces.Data {
		dungeon.CountCompletesFromTrace(&t, result)
	}

	if result.GetTotal() != 40 {
		log.Fatalf("expect 40, but %d", result.GetTotal())
	}
}
