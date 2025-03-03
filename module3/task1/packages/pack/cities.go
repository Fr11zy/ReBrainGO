package pack

import(
	"task1/packages/wordz"
)

var Cities = map[string] string {"one" : "Moscow",
								 "two" : "London",
								 "three" : "Tyumen",
								 "four"  : "New-York",
								 "five"  : "Kazan'"}

func City() string {
	return Cities[wordz.Random()]
}