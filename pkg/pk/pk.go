package pk

func Reverse(input string) (result string) {
	for _, c := range input {
		result = string(c) + result
	}
	return result
}