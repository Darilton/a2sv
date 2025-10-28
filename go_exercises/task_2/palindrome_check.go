package main

import "fmt"
import "strings"

func reverse(s string) string {
	ans := ""
	for i := len(s) - 1; i > -1; i-- {
		ans += string(s[i])
	}
	return ans
}
func is_valid_char(s string) bool {
	return (s >= "a" && s <= "z")
}

func is_palindrome(s string) bool {
	ans := ""
	for i := range(len(s)) {
		curr := strings.ToLower(string(s[i]))
		if is_valid_char(curr) {
			ans += curr
		}
	}
	return reverse(ans) == ans
}

func main() {
	s := "banananab"
	fmt.Println(s, ":", is_palindrome(s))

	s = "banan?anab"
	fmt.Println(s, ":", is_palindrome(s))

	s = "Banan?anab"
	fmt.Println(s, ":", is_palindrome(s))

	s = "Aanan?anab"
	fmt.Println(s, ":", is_palindrome(s))

	s = "Ba               nan?anab"
	fmt.Println(s, ":", is_palindrome(s))
}
