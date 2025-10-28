package main

import "fmt"

func is_valid_char(s string) bool {
	return (string(s) >= "a" && string(s) <= "z") || (string(s) >= "A" && string(s) <= "Z")
}

func word_freq_count(s string) map[string]int {
	var ans = make(map[string]int)
	for i := range(len(s)) {
		curr := string(s[i])
		if is_valid_char(curr) {
			ans[curr]++
		}
	}
	return ans
}

func main() {
	fmt.Println(word_freq_count("Banana_?Do you want it?"))
}
