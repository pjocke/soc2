package wahlin

import (
	"fmt"
	"strings"
)

// Encode a string to its Wåhlin encoded equivalent
func Encode(decoded string) string {
	//var consonants = [20]byte{'B', 'C', 'D', 'F', 'G', 'H', 'J', 'K', 'L', 'M', 'N', 'P', 'Q', 'R', 'S', 'T', 'V', 'W', 'X', 'Z'}

	var word = []rune(strings.ToUpper(decoded))
	var encoded []rune

	var i int
	var letter rune
	fmt.Printf("%d\n", len(word))
	for i = 0; i < len(word); i++ {
		letter = word[i]
		fmt.Printf("%q, %d, %d\n", letter, i, i+1)

		if letter == 'A' {
			// Rule set A
			if i+1 < len(word) && word[i+1] == 'E' {
				continue
			} else {
				encoded = append(encoded, letter)
			}
		} else if letter == 'C' && i+1 < len(word) && word[i+1] != 'H' {
			// Rule set C
			if word[i+1] == 'A' || word[i+1] == 'O' || word[i+1] == 'U' || word[i+1] == 'L' || word[i+1] == 'S' || word[i+1] == 'R' {
				encoded = append(encoded, 'K')
			} else if word[i+1] == 'E' || word[i+1] == 'I' || word[i+1] == 'Y' {
				encoded = append(encoded, 'S')
			} else {
				encoded = append(encoded, letter)
			}
		} else if letter == 'C' && i+1 < len(word) && word[i+1] == 'H' {
			// Rule set CH
			if i+2 < len(word) && (word[i+2] == 'R' || word[i+2] == 'S' || word[i+2] == 'T') {
				encoded = append(encoded, 'K')
			}
		} else if letter == 'D' {
			// Rule set D
			if i+1 < len(word) && (word[i+1] == 'J' || word[i+1] == 'T') {
				continue
			} else {
				encoded = append(encoded, letter)
			}
		} else if letter == 'F' {
			// Rule set F
			if (i > 0 && (word[i-1] == 'A' || word[i-1] == 'E' || word[i-1] == 'I' || word[i-1] == 'O' || word[i-1] == 'U' || word[i-1] == 'Y' || word[i-1] == 'Å' || word[i-1] == 'F' || word[i-1] == 'L' || word[i-1] == 'R')) ||
				(i+1 < len(word) && (word[i+1] == 'B' || word[i+1] == 'C' || word[i+1] == 'D' || word[i+1] == 'G' || word[i+1] == 'H' || word[i+1] == 'K' || word[i+1] == 'M' || word[i+1] == 'N' || word[i+1] == 'P' || word[i+1] == 'S' || word[i+1] == 'V')) {
				encoded = append(encoded, 'V')
			} else {
				encoded = append(encoded, letter)
			}
		} else if letter == 'G' {
			// Rule set G
			if (i-1 > 0 && (word[i-1] == 'R')) ||
				(i == 0 && i+1 < len(word) && word[i+1] == 'E') ||
				(i+1 < len(word) && (word[i+1] == 'I' || word[i+1] == 'Y' || word[i+1] == 'Ö')) {
				encoded = append(encoded, 'J')
			} else {
				encoded = append(encoded, letter)
			}
		}
	}

	return string(encoded)
}
