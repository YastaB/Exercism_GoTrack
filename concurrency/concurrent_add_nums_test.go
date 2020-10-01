package main

import "testing"

var cases = []struct {
	start    int
	end      int
	skip     int
	expected int
}{
	{0,15, 5, 120},
	{1,15,5,120},
	{2,15,5,119},
	{0,0,0,0},
	{0,2000,250,1000*2001},
}

func TestAddConcurrent(t *testing.T) {
	for _, tCase := range cases{
		sum := AddConcurrent(tCase.start, tCase.end, tCase.skip)
		if sum != tCase.expected{
			t.Fatalf("Given: (%d:%d:%d) Result: %d Expected: %d", tCase.start, tCase.end, tCase.skip, sum,tCase.expected)
		}
	}
}
