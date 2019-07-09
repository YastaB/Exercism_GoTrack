package accumulate

import "fmt"

//Accumulate applies the converter function to all members of given array
func Accumulate(given []string, converter func(string) string) []string {
	result := make([]string, len(given))
	for index, str := range given {
		result[index] = converter(str)
	}
	return result
}

func add(str string) string {
	return str + str
}

func testAccumulate() {
	arr := []string{"a", "b"}
	res := Accumulate(arr, add)
	fmt.Println("Res: ", res, " Given arr: ", arr)

}
