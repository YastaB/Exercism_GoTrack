package cryptosquare

import "testing"

var tests = []struct {
	pt string // plain text
	ct string // cipher text
}{
	{
		"s#$%^&plunk",
		"su pn lk",
	},
	{
		"1, 2, 3 GO!",
		"1g 2o 3 ",
	},
	{
		"1234",
		"13 24",
	},
	{
		"123456789",
		"147 258 369",
	},
	{
		"123456789abc",
		"159 26a 37b 48c",
	},
	{
		"Never vex thine heart with idle woes",
		"neewl exhie vtetw ehaho ririe vntds",
	},
	{
		"ZOMG! ZOMBIES!!!",
		"zzi ooe mms gb ",
	},
	{
		"Time is an illusion. Lunchtime doubly so.",
		"tasney inicds miohoo elntu  illib  suuml ",
	},
	{
		"We all know interspecies romance is weird.",
		"wneiaw eorene awssci liprer lneoid ktcms ",
	},
	{
		"Madness, and then illumination.",
		"msemo aanin dnin  ndla  etlt  shui ",
	},
	{
		"Vampires are people too!",
		"vrel aepe mset paoo irpo",
	},
	{
		"",
		"",
	},
	{
		"1",
		"1",
	},
	{
		"12",
		"1 2",
	},
	{
		"12 3",
		"13 2 ",
	},
	{
		"12345678",
		"147 258 36 ",
	},
	{
		"123456789a",
		"159 26a 37  48 ",
	},
	{
		"If man was meant to stay on the ground god would have given us roots",
		"imtgdvs fearwer mayoogo anouuio ntnnlvt wttddes aohghn  sseoau ",
	},
	{
		"Have a nice day. Feed the dog & chill out!",
		"hifei acedl veeol eddgo aatcu nyhht",
	},
}

func TestEncode(t *testing.T) {
	for _, test := range tests {
		if ct := Encode(test.pt); ct != test.ct {
			t.Fatalf(`Encode(%q):
got  %q
want %q`, test.pt, ct, test.ct)
		}
	}
}

var PrepareTestCases = []struct {
	given    string
	expected string
}{
	{
		given:    "ABC",
		expected: "abc",
	},
	{
		given:    ",.00**--",
		expected: "00",
	},
	{
		given:    ",.00**--",
		expected: "00",
	},
	{
		given:    "This is a test, check it out!!...",
		expected: "thisisatestcheckitout",
	},
}

var RectangleSizeTestCases = []struct {
	given    int
	expectedC int
	expectedR int
}{
	{
		given:   54,
		expectedC: 8,
		expectedR: 7,
	},
	{
		given:   64,
		expectedC: 8,
		expectedR: 9,
	},
}

func TestRectangleSize(t *testing.T){
	for _, tcase := range RectangleSizeTestCases{
		if c,r := findRectSize(tcase.given) ; c != tcase.expectedC || r != tcase.expectedR{
			t.Fatalf(`RectSize(%d):
got  C:%d,R:%d 
want C:%d,R:%d`, tcase.given, c, r , tcase.expectedC, tcase.expectedR)
		}
	}
}

func TestPrepare(t *testing.T) {

	for _, tcase := range PrepareTestCases{
		if result := prepareString(tcase.given) ; result != tcase.expected{
			t.Fatalf(`Prepare(%q):
got  %q
want %q`, tcase.given, result, tcase.expected)
		}
	}

}

func BenchmarkEncode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			Encode(test.pt)
		}
	}
}
