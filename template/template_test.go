package template

import (
	"testing"
)

func TestMonthLiteral(t *testing.T) {
	cases := []struct {
		Case   int
		Expect string
	}{
		{Case: 1, Expect: "January"},
		{Case: 2, Expect: "February"},
		{Case: 3, Expect: "March"},
		{Case: 4, Expect: "April"},
		{Case: 5, Expect: "May"},
		{Case: 6, Expect: "June"},
		{Case: 7, Expect: "July"},
		{Case: 8, Expect: "August"},
		{Case: 9, Expect: "September"},
		{Case: 10, Expect: "October"},
		{Case: 11, Expect: "November"},
		{Case: 12, Expect: "December"},
	}

	for _, c := range cases {
		testMonthLiteral(t, c.Case, c.Expect, nil)
	}

	testMonthLiteral(t, 0, "", ErrMonthOutOfRange)
	testMonthLiteral(t, 13, "", ErrMonthOutOfRange)
}

func testMonthLiteral(t *testing.T, m int, expectLit string, expectErr error) {
	lit, err := GetMonthLiteral(m)

	if err != nil {
		if expectErr != err {
			t.Fatalf("expect no error but we got: %s", err.Error())
		}
	}

	if lit != expectLit {
		t.Fatalf("expect %s, but actual %s", expectLit, lit)
	}
}
