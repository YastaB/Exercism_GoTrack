package proverb

// Proverb generats provers according to given rhymes.
func Proverb(rhyme []string) []string {

	var slice []string
	lenRhyme := len(rhyme)
	if lenRhyme > 1 {
		for i := 0; i < lenRhyme-1; i++ {
			str := "For want of a " + rhyme[i] + " the " + rhyme[i+1] + " was lost."
			slice = append(slice, str)
		}
	}

	if lenRhyme >= 1 {
		slice = append(slice, "And all for the want of a "+rhyme[0]+".")
	}

	return slice
}
