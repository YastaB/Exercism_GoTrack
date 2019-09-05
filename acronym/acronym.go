package acronym

import "strings"

// Abbreviate get all concatanate first letter of each word in string
func Abbreviate(s string) string {
	var abbr = ""
	wordArr := strings.FieldsFunc(s, fields)
	for i := 0; i < len(wordArr); i++ {
		abbr += strings.ToUpper(wordArr[i][:1])
	}
	return abbr
}

func fields(r rune) bool {
	return (r == ' ') || (r == '-') || (r == '_')
}

func main() {
	println("Result: " + Abbreviate("Test Yoo Bra"))
}
