package main

import "fmt"

func shouldEscape(c rune) bool {
	switch c {
	case ' ', '?', '&', '=', '#', '+', '%':
		return true
	}
	return false
}

func main() {
	for _, s := range "日本語ですか？英語ですか?" {
		if shouldEscape(s) {
			fmt.Printf("character %#U should be escaped\n", s)
		} else {
			fmt.Printf("character %#U should not be escaped\n", s)
		}
	}
}
