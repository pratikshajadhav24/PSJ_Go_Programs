// string functions

package main

import (
	"fmt"
	"strings"
)

func main() {
	// 1. Concatenate strings
	str1 := "Hello"
	str2 := "World"
	concat := str1 + " " + str2
	fmt.Println("1. Concatenate:", concat)

	// 2. Compare strings
	fmt.Println("2. Compare:", strings.Compare(str1, str2)) // Returns 0 if equal, -1 if str1 < str2, 1 if str1 > str2

	// 3. Check string equality
	fmt.Println("3. Equal:", str1 == str2)

	// 4. String contains
	fmt.Println("4. Contains:", strings.Contains(concat, "World"))

	// 5. Prefix check
	fmt.Println("5. HasPrefix:", strings.HasPrefix(concat, "Hello"))

	// 6. Suffix check
	fmt.Println("6. HasSuffix:", strings.HasSuffix(concat, "World"))

	// 7. Index of a substring
	fmt.Println("7. Index:", strings.Index(concat, "World"))

	// 8. Last index of a substring
	fmt.Println("8. LastIndex:", strings.LastIndex(concat, "o"))

	// 9. Repeat a string
	fmt.Println("9. Repeat:", strings.Repeat("Go", 3))

	// 10. Replace substrings
	replaced := strings.Replace(concat, "World", "GoLang", 1)
	fmt.Println("10. Replace:", replaced)

	// 11. Split a string
	split := strings.Split(concat, " ")
	fmt.Println("11. Split:", split)

	// 12. Join a string array
	joined := strings.Join(split, ", ")
	fmt.Println("12. Join:", joined)

	// 13. Convert to uppercase
	fmt.Println("13. ToUpper:", strings.ToUpper(concat))

	// 14. Convert to lowercase
	fmt.Println("14. ToLower:", strings.ToLower(concat))

	// 15. Trim spaces
	trimmed := strings.Trim("  Hello World  ", " ")
	fmt.Println("15. Trim:", trimmed)

	// 16. Trim specific characters
	trimChars := strings.Trim("##GoLang##", "#")
	fmt.Println("16. TrimChars:", trimChars)

	// 17. Count occurrences of a substring
	fmt.Println("17. Count:", strings.Count(concat, "o"))

	// 18. Check if empty
	empty := ""
	fmt.Println("18. IsEmpty:", len(empty) == 0)

	// 19. Length of a string
	fmt.Println("19. Length:", len(concat))

	// 20. Replace all substrings
	replacedAll := strings.ReplaceAll(concat, "o", "0")
	fmt.Println("20. ReplaceAll:", replacedAll)
}
